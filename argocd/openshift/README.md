# Installation steps for th2 via ArgoCD 
NOTE: files could be found in an appropriate folders
1. ### [Preparation for cassandra installation](https://github.com/th2-net/th2-infra/tree/master/argocd/openshift/cassandra-instance)
    <details>
      <summary>cassandra preparation steps </summary>  

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
      Create application in ArgoCD GUI and fill the form as in [argo-cassandra.yaml](./cassandra-instance/argo-cassandra.yaml)
      
      Or perform CLI command to deploy cassandra:
      ```console
      oc -n <namespace-with-argocd-instance> apply -f argo-cassandra.yaml
      ```
      
      ## SCC for cassandra SA should be added manualy
      
      Add anyuid SCC to cassandra SA:
      ```console
      oc adm policy add-scc-to-user anyuid -n argo-cassandra -z cassandra
      ```

    </details>

2. ### [Monitoring configuration](https://github.com/th2-net/th2-infra/tree/master/argocd/openshift/monitoring)
    <details>
      <summary>monitoring preparation steps </summary>  

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
      In case of monitoring stack is been deployed before TH2-infra - grafana dashboards and plugins could be temporary switched off in [loki-values.yaml](./monitoring/loki-values.yaml).
      
      Grafana helm [repository](https://grafana.github.io/helm-charts) should be added to /settings/repos in ArgoCD GUI
      
      ## ArgoCD Application deployment
      Create 2 applications in ArgoCD GUI and fill the forms as in [argo-monitoring.yaml](./monitoring/argo-monitoring.yaml) and [grafana-access.yaml](./monitoring/grafana-access/grafana-access.yaml)
      
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

    </details>

3. ### [Sealed secret configuration](https://github.com/th2-net/th2-infra/tree/master/argocd/openshift/sealed-secrets)
    <details>
      <summary>Sealed secret preparation steps </summary> 

      # Sealed Secrets deployment
      ## Create namespace
      
      Detail documentation for Sealed Secrets can be found [here](https://github.com/bitnami-labs/sealed-secrets/tree/v0.22.0#sealed-secrets-for-kubernetes)
      
      Create service namespace and make sure that it is available for ArgoCD:
      ```console
      oc create namespace argo-service
      oc label namespace argo-service argocd.argoproj.io/managed-by=<namespace-with-argocd-instance>
      ```
      
      ## Deploy Sealed Secrets
      
      Create application in ArgoCD GUI and fill the form as in [argo-secrets.yaml](./sealed-secrets/argo-ssecrets.yaml)
      
      Or perform CLI command to deploy cassandra:
      ```console
      oc -n <namespace-with-argocd-instance> create -f argo-secrets.yaml
      ```
      Sealing key renewal is turned off in [values.yaml](./sealed-secrets/values.yaml)
      
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

    </details>

4. ### [th2 installation](https://github.com/th2-net/th2-infra/tree/master/argocd/openshift/sealed-secrets)
    <details>
      <summary>th2 installation </summary> 

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
      Create application in ArgoCD GUI and fill the form as in [argo-th2.yaml](./values/argo-th2.yaml)
      
      Or perform CLI command to deploy:
      ```console
      oc -n <namespace-with-argocd-instance> apply -f argo-th2.yaml
      ```

    </details>
