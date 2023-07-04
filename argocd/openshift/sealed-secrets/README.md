# Sealed Secrets deployment
## Create namespace

Detail documentation for Sealed Secrets can be found [here](https://github.com/bitnami-labs/sealed-secrets/tree/v0.22.0#sealed-secrets-for-kubernetes)

Create service namespace and make sure that it is available for ArgoCD:
```console
oc create namespace argo-service
oc label namespace argo-service argocd.argoproj.io/managed-by=<namespace-with-argocd-instance>
```

## Deploy Sealed Secrets

Create application in ArgoCD GUI and fill the form as in [argo-secrets.yaml](./argo-ssecrets.yaml)

Or perform CLI command to deploy cassandra:
```console
oc -n <namespace-with-argocd-instance> create -f argo-secrets.yaml
```
Sealing key renewal is turned off in [values.yaml](./values.yaml)

## Install kubeseal for operator's console

kubeseal tool provides a way to cipher secrets - it should be installed on operator's host

```console
wget https://github.com/bitnami-labs/sealed-secrets/releases/download/<release-tag>/kubeseal-<version>-linux-amd64.tar.gz
tar -xvzf kubeseal-<version>-linux-amd64.tar.gz kubeseal
sudo install -m 755 kubeseal /usr/local/bin/kubeseal
```
Certificate can be fetched from controller log and stored on operator's host:
```console
oc -n argo-service logs sealed-secrets-57cc6d7d5-qltk9
```

After that kubeseal may be checked e.g. the following way:

```console
kubeseal --cert ~/.kube/profiles/okd.sealed.pubkey.pem
```
