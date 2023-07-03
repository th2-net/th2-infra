# th2 deployment
## Create namespace

argo-service namespace has already been created during Sealed Secrets operator deployment else it should be created the following way:

```console
oc create namespace argo-service
oc config set-context --current --namespace argo-service
oc label namespace argo-service argocd.argoproj.io/managed-by=<namespace-with-argocd-instance>
```

To be able to run th2 workloads in multiple namespaces with the same domain name the route admission policy need to be [configured](https://docs.openshift.com/container-platform/4.11/networking/ingress-operator.html#nw-route-admission-policy_configuring-ingress) the following way:

```console
oc -n openshift-ingress-operator patch ingresscontroller/default --patch '{"spec":{"routeAdmission":{"namespaceOwnership":"InterNamespaceAllowed"}}}' --type=merge
```

## Upload secrets to the cluster

In order to store [secrets.yaml](https://github.com/th2-net/th2-infra#create-secret-with-th2-credentials) with sensetive data safely in git repository it should be [sealed](https://github.com/bitnami-labs/sealed-secrets/tree/v0.22.0#usage) with kubeseal and uploaded to the cluster as SealedSecret resource:


Create [rabbitmq-secret.yaml](tmp/rabbitmq-secret.yaml) with the following content:
```console
rabbitmq-password=<my_password>
rabbitmq-erlang-cookie=<random_string>
```

```console
oc create secret generic rabbitmq --dry-run=client --from-env-file=tmp/rabbitmq-secret.yaml -o json > tmp/rabbitmq-secret.json
oc create secret generic cassandra --dry-run=client --from-env-file=tmp/cassandra-secret.yaml -o json > tmp/cassandra-secret.json
oc create secret docker-registry nexus-proxy --docker-server=<your-registry-server> --docker-username=<your-name> --docker-password=<your-password> --dry-run=client -o json > tmp/nexus-proxy.secret.json
oc create secret docker-registry nexus-private --docker-server=<your-one-more-registry-server> --docker-username=<your-name> --docker-password=<your-password> --dry-run=client -o json > tmp/nexus-private.secret.json
```
base64 encoded private key should be inserted to [inframgr-secret.json](tmp/inframgr-secret.json) secret manualy:
```console
{
  "kind": "SealedSecret",
  "apiVersion": "bitnami.com/v1alpha1",
  "metadata": {
    "name": "inframgr-secret",
    "namespace": "argo-service",
    "creationTimestamp": null
  },
  "spec": {
    "template": {
      "metadata": {
        "name": "inframgr-secret",
        "namespace": "argo-service",
        "creationTimestamp": null
      }
    },
    "encryptedData": {
      "id_rsa": "<HERE>    <=== !!! " 
    }
  }
}
```

Encrypt the secrets, after that they can be stored in git repo:
```console
for SECRET in $(ls tmp/ | grep .json$); do kubeseal --cert ~/.kube/profiles/okd.sealed.pubkey.pem <tmp/$SECRET >secrets/$SECRET; done
```

We do not install `converter` here, because initially assumed that OpenShift cluster doesn't have write access to git repository.

## jupyterhub installation (comes within TH2 infra installation)

To deploy jupyterhub in OpenShift (OKD) `anyuid` SCC should be added to SA hub:

```console
oc adm policy add-scc-to-user anyuid -n argo-service -z hub
```

## ArgoCD Application deployment
Create application in ArgoCD GUI and fill the form as in [argo-th2.yaml](./argo-th2.yaml)

Or perform CLI command to deploy:
```console
oc -n <namespace-with-argocd-instance> apply -f argo-th2.yaml
```
