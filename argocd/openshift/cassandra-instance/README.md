# Cassandra deployment

### [Cass-operator](https://github.com/k8ssandra/cass-operator) should be installed beforehand e.g. from OperatorHub

## Create namespace

Create namespace for cassandra and make sure that it is available for ArgoCD:
```console
oc create namespace argo-cassandra
oc config set-context --current --namespace argo-cassandra
oc label namespace argo-cassandra argocd.argoproj.io/managed-by=<namespace-with-argocd-instance>
```
Make sure that you have storage class with reclame policy RETAIN chosen for cassandra persistence.

## Upload secrets to the cluster

In order to store [secrets.yaml](https://github.com/th2-net/th2-infra#create-secret-with-th2-credentials) with sensetive data safely in git repository it should be [sealed](https://github.com/bitnami-labs/sealed-secrets/tree/v0.22.0#usage) with kubeseal and uploaded to the cluster as SealedSecret resource:

Create [cassandra-secret.yaml](tmp/cassandra-secret.yaml) with the following content:
```console
password=<cassandra superuser password>
username=th2
```
Encrypt the secret, after that it can be stored in git repo:
```console
oc create secret generic cassandra-secret --dry-run=client --from-env-file=tmp/cassandra-secret.yaml -o json > tmp/cassandra-secret.json
kubeseal --cert ~/.kube/profiles/okd.sealed.pubkey.pem <tmp/cassandra-secret.json >cassandra-secret.json
```
## ArgoCD Application deployment
Create application in ArgoCD GUI and fill the form as in [argo-cassandra.yaml](./argo-cassandra.yaml)

Or perform CLI command to deploy cassandra:
```console
oc -n <namespace-with-argocd-instance> apply -f argo-cassandra.yaml
```

## SCC for cassandra SA should be added manualy

Add anyuid SCC to cassandra SA:
```console
oc adm policy add-scc-to-user anyuid -n argo-cassandra -z cassandra
```
