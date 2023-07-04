# Monitoring deployment

## Enable monitoring for user-defined projects - initial settings.
### Configuration in this article should be performed with ADMINISTRATIVE PRIVILEGES.
In order to be able to get TH2 custom metrics you need to [enable](https://docs.openshift.com/container-platform/4.11/monitoring/enabling-monitoring-for-user-defined-projects.html) monitoring for user-defined projects. 

Create `cluster-monitoring-config` ConfigMap object in `openshift-monitoring namespace` (or edit it if you already have a new one. Add `enableUserWorkload: true` under data/config.yaml.):
```console
oc apply -f cluster-monitoring-config.yaml
```
Save the file to apply the changes. Monitoring for user-defined projects is then enabled automatically.

You can optionally create and configure the `user-workload-monitoring-config` ConfigMap object (`user-workload-monitoring-config.yaml`) in the `openshift-user-workload-monitoring` project. You can add configuration options to this ConfigMap object for the components that monitor user-defined projects. EG - you can distribute workload on worker nodes, and configure user workload prometheus persistance

Check that the `prometheus-operator`, `prometheus-user-workload` and `thanos-ruler-user-workload` pods are running in the `openshift-user-workload-monitoring` project.
```console
oc get po -n openshift-user-workload-monitoring
```
It might take a short while for the pods to start.

## Loki and Grafana deployment
### TO;DO - replace helm repo with git repo for helm chart - for restricted environment; Get rid of Multiple Sources for an Application;

Create namespace for application and make sure that it is available for ArgoCD:
```console
oc create ns argo-monitoring
oc config set-context --current --namespace argo-monitoring
oc label namespace argo-monitoring argocd.argoproj.io/managed-by=openshift-gitops

```
In case of monitoring stack is been deployed before TH2-infra - grafana dashboards and plugins could be temporary switched off in [loki-values.yaml](./loki-values.yaml).

Grafana helm [repository](https://grafana.github.io/helm-charts) should be added to to /settings/repos in ArgoCD GUI

## ArgoCD Application deployment
Create 2 applications in ArgoCD GUI and fill the forms as in [argo-monitoring.yaml](./argo-monitoring.yaml) and [grafana-access.yaml](./grafana-access/grafana-access.yaml)

Or perform CLI command to deploy monitoring:
```console
oc -n openshift-gitops apply -f argo-monitoring.yaml
```

`privileged` security context constraint (SCC) should be added to serviceaccounts loki-promtail and `anyuid` SCC should be added to SAs loki and loki-grafana:
```console
oc adm policy add-scc-to-user privileged -n argo-monitoring -z monitoring-promtail
oc adm policy add-scc-to-user anyuid -z monitoring-grafana -n argo-monitoring
```

### Preparing access to internal OpenShift Prometheus
Originaly Grafana has no access to OpenShift Prometheus to provide it a proper Authorization HTTP Header should be passed to Prometheus Data Source  [configuration](https://medium.com/faun/openshift-leveraging-prometheus-cluster-metrics-in-your-own-grafana-7077fb0725ab).

The following creates ServiceAccount in argo-monitoring namespace with Secret of `kubernetes.io/service-account-token` type and bindes `cluster-monitoring-view ClusterRole` to mentioned ServiceAccount.

Required resources should already be created with ArgoCD we can check the secret following way:
```console
oc get secrets -n argo-monitoring | grep grafana-prometheus-access-token
```
and obtain Authorization HTTP Header with the following command:
```console
echo -n "Bearer "; oc -n argo-monitoring get secrets grafana-prometheus-access-token-hcr6q -o jsonpath="{..token}" | base64 -d ; printf "\n"
```
The output should be added to grafana.datasources.datasources.yaml.datasources.secureJsonData.httpHeaderValue1