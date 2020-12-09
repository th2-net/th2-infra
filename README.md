# th2 installation

## Prerequisites
Before you begin, please check the following prerequisites:
  * Fully functioning Kubernetes cluster suitable for your bussiness needs, please refer to [technical requirements](https://github.com/th2-net/th2-documentation/wiki/Technical-Requirements)
  * Operator-box that meets [hardware](https://github.com/th2-net/th2-documentation/wiki/Technical-Requirements) and [software](https://github.com/th2-net/th2-documentation/wiki/Technical-Requirements#software-requirements) requirements
* Installed [Apache Cassandra](https://cassandra.apache.org/) - [technical requirements](https://github.com/th2-net/th2-documentation/wiki/Technical-Requirements#apache-cassandra-cluster-hardware-requirements)
  
All th2 components are deployed via Helm charts by [Helm](https://helm.sh/) and [Helm Operator](https://docs.fluxcd.io/projects/helm-operator/en/stable/).

## Steps
The following steps should be performed on the operator-box for th2-infra deployment:
<!--ts-->
   * [Download th2 git repositories](#th2-git-repositories)
   * [Monitoring deployment](#monitoring-deployment)
   * [Cluster configuration](#cluster-configuration)
   * [th2 deployment](#th2-deployment)
<!--te-->

## th2 git repositories
Installation of th2 infra requires two git repositories. The information regarding this repository and its usage can be found in this guide below:
* https://github.com/th2-net/th2-infra - consists of charts and its values for deployment infrastructure components. The repository is common for everyone, but you can fork or clone it if you need to customize values.
* https://github.com/th2-net/th2-infra-schema-demo - schema repository. It's used by `th2-infra-mgr`.

The first step that should be done in the th2 deployment process is copying th2-infra repository into your operator-box:
```
git clone https://github.com/th2-net/th2-infra.git
```
change the current directory
```
cd ./th2-infra
```
Then https://github.com/th2-net/th2-infra-schema-demo should be created in your git as a fork or template:
  * [how to create template](https://docs.github.com/en/free-pro-team@latest/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template)
  * [how to fork](https://docs.github.com/en/free-pro-team@latest/github/getting-started-with-github/fork-a-repo#fork-an-example-repository)


## th2 infra namespaces
th2 infra components are split into two namespaces: _`monitoring`_ and _`service`_. These namespaces will be created below.

Next components of prometheus and grafana monitoring stack are deployed into _`monitoring`_ namespace:
  * [kubernetes-dashboard](https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/)
  * [grafana](https://grafana.com/oss/grafana/)
  * [loki](https://grafana.com/oss/loki/)
  * [prometheus](https://grafana.com/oss/prometheus/)

The _`service`_ namespace is used for core services of this project:
  * [RabbitMQ](https://www.rabbitmq.com/)
  * [NGINX Ingress Controller](https://kubernetes.github.io/ingress-nginx/)
  * [Helm Operator](https://github.com/fluxcd/helm-operator)
  
and for infrastructure components:
  * [th2-infra-editor]()
  * [th2-infra-operator](https://github.com/th2-net/th2-infra-operator)
  * [th2-infra-mgr](https://github.com/th2-net/th2-infra-mgr)
  
The following picture describes th2-infra cluster configuration:

![th2-infra](https://user-images.githubusercontent.com/690243/101453864-0c7b5980-3941-11eb-8616-11f00095ea99.png)

* Create namespaces
    * command:
    ```
    kubectl create namespace monitoring
    kubectl create namespace service
    ```
    * You can check result using this command:
    ```
    kubectl get namespaces 
    ```
    * In the output you should see the names of these namespaces:
    ```
        NAME              STATUS   AGE
        .......
        monitoring        Active   15s
        service           Active   7s
        .......
    ```


## Data persistence

Data persistence is required for the following components: Grafana, Prometheus, Loki, RabbitMQ components and should be set up on this step.

_Note: Examples below use HostPath type of [Persistent Volume(PV)](https://kubernetes.io/docs/concepts/storage/persistent-volumes/). Please read the documentation to choose an appropriate PV type for your environment_

Steps:

* the following command can require root permissions, create directory on th2 node:
```
mkdir /opt/grafana /opt/prometheus /opt/loki /opt/rabbitmq
```
* set node name in `./values/pvs.yaml`
* create PVs and PVCs:
    ```
    kubectl apply -f ./values/pvs.yaml
    kubectl apply -f ./values/pvcs.yaml
    ```

If you would like to include th2 read components into your configuration, you also have to set up a dedicated PersistentVolume for th2-read log directory.
You should add PersistentVolume mapped to /opt/components directory and then create PersistentVolumeClaim once a schema namespace installed. PV and PVC examples can be found here [./values/persistence/](./values/persistence/)

```
mkdir /opt/components
```
* set node name in `./values/persistence/pv.yaml`
* create PV:
    ```
    kubectl apply -f ./values/persistence/pv.yaml
    ```
* create 
    ```
    kubectl apply -f ./values/persistence/pvc.yaml
    ```

Details for th2-read-log [README.md](https://github.com/th2-net/th2-read-log#configuration)

## Monitoring deployment

_Note: It's an optional step, but it gets slightly simpler checking the result of installation. In all installation commands we explicitly define namespaces to avoid possible mistakes._
* Switch namespace to monitoring
  ```
  kubectl config set-context --current --namespace=monitoring
  ```
* Define Grafana and Dashboard host names:
  * in the [./values/dashboard.values.yaml](./values/dashboard.values.yaml) file
    ```
    ingress:
      hosts:
        - <dashboard_host_name>
    ```
  * in the [./values/prometheus-operator.values.yaml](./values/prometheus-operator.values.yaml) file
    ```
    grafana:
      ingress:
        hosts:
          - <grafana_host_name>
    ```

* Install [Kubernetes Dashboard](https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/)
    ```
    helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
    helm install dashboard -n monitoring kubernetes-dashboard/kubernetes-dashboard -f ./values/dashboard.values.yaml
    ```
* Deploy components
    ```
    helm repo add loki https://grafana.github.io/loki/charts
    helm repo add stable https://charts.helm.sh/stable
    helm upgrade --install loki --namespace=monitoring loki/loki-stack -f ./values/loki.values.yaml
    helm upgrade --install prometheus stable/prometheus-operator -n monitoring -f ./values/prometheus-operator.values.yaml
    ```
* Check result:
    * command:
        ```
        kubectl get pods
        ```
    * output:
        ```
        NAME                                                     READY   STATUS    RESTARTS   AGE
        ........
        pod/dashboard-kubernetes-dashboard-77d85586db-j9v8f   1/1     Running   0          56s
        alertmanager-prometheus-prometheus-oper-alertmanager-0   2/2     Running   0          75s
        loki-0                                                   1/1     Running   0          4m47s
        loki-promtail-wqfml                                      1/1     Running   0          4m47s
        prometheus-grafana-68f8dd6d57-2gtns                      2/2     Running   0          82s
        prometheus-kube-state-metrics-75d4cc9dbd-psb88           1/1     Running   0          82s
        prometheus-prometheus-node-exporter-gfzp6                1/1     Running   0          82s
        prometheus-prometheus-oper-operator-df668d457-snxks      1/1     Running   0          82s
        prometheus-prometheus-prometheus-oper-prometheus-0       3/3     Running   1          65s        
        ........
        ```

default password in Grafana: admin/prom-operator. Need to be changed
Add loki Datasource as http://loki:3100 and import Dashboard from ./values/components-logs.json

* Check access to Grafana _(default user/password: `admin/prom-operator`. Need to be changed)_: <br>
  http://your-host:30000/grafana/login

## Cluster configuration
Once all of the required software is installed on your test-box and operator-box and th2-infra repositories are ready you can start configuring the cluster.

* Switch namespace to service:
```
kubectl config set-context --current --namespace=service
```

### Set up access to Git repositories
Two types of access to repositories are used in th2 - via `https` and `ssh`. The `ssh` access is required by **th2-infra-mgr** component and `https` by **helm-operator**. So, we need to set up both of them.
#### 1. SSH access:

1. Generate keys without passphrase  
    ```
    ssh-keygen -t rsa -m pem -f ~/.ssh/id_gh_rsa
    ``` 
2. [Add a new SSH key to your GitHub account](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/adding-a-new-ssh-key-to-your-github-account)
3. Create infra-mgr secret from private key:
   ```
   kubectl -n service create secret generic infra-mgr --from-file=infra-mgr=$HOME/.ssh/id_gh_rsa
   ```
     
#### 2. HTTPS access for charts (part of th2-infra repository):
Create secret for git access (only for private repositories)
```
kubectl -n service create secret generic git-chart-creds --from-literal=username=git-username --from-literal=password=git-password
```
If you use a private repository for charts of project (for some security reasons i.e.) instead of public, you should provide valid credentials for `git-username` and `git-password` in the command above. Using a Personal Access Token(PAT) is the better choice instead of plain password.
Read more about this:
* [Creating a personal access token on Github](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token)
* [Creating a deployment token on Gitlab](https://docs.gitlab.com/ee/user/project/deploy_tokens/#creating-a-deploy-token)

If you use a public repository for the charts of project, you can keep the values for `git-username` and `git-password` as is or use empty values like this:
```
kubectl -n service create secret generic git-chart-creds --from-literal=username= --from-literal=password=
```

### Set the repository with configuration
* set `infraMgr.git.repository` value in the [./values/service.values.yaml](./values/service.values.yaml) file to **ssh** link of your schema repository, e.g:
```
infraMgr:
  git:
    repository: git@github.com:th2-net/th2-infra-demo-configuration.git
```

### Define cassandra host name
* set `cassandra.host` value for cassandra in the [./values/service.values.yaml](./values/service.values.yaml) file.
```
cassandra:
  internal: false
  host: <cassandra-host>
```

### Define th2 ingress hostname
Please add `ingress.hostname` value if required into [./values/service.values.yaml](./values/service.values.yaml) file otherwise th2 http services will be available on node IP address
```
ingress:
  host: example.com
```

### Create secret with th2 credentials

Create secrets.yaml in `./` folder (*do not commit into git*). Please provide valid credentials for Cassandra DB. Example:
```
# reguired only if images in private repository
# productRegistry:
#  username: user
#  password: password
#  name: registry.example.com # core components registry

# reguired only if images in private repository
# solutionRegistry:
#  username: user
#  password: password
#  name: private-registry.example.com # components registry

cassandra:
  dbUser:
    user: <user-name>
    password: <password>

rabbitmq:
  rabbitmqUsername: th2
  rabbitmqPassword: rab-pass
  managementUsername: th2-mng
  managementPassword: rab-mng-pass
  # must be random string
  rabbitmqErlangCookie: cookie
```

## th2 deployment
### Install helm-operator 
```
helm repo add fluxcd https://charts.fluxcd.io
helm install --version=1.2.0 helm-operator -n service fluxcd/helm-operator -f ./values/helm-operator.values.yaml
```

### Install NGINX Ingress Controller
* Install NGINX Ingress Controller:
    ```
    helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
    helm install -n service --version=3.12.0 ingress ingress-nginx/ingress-nginx -f ./values/ingress.values.yaml
    ```
* Check result:
    * command:
        ```
        kubectl get pods
        ```
    * output:
        ```
        NAME                                                READY   STATUS    RESTARTS   AGE
        ........
        ingress-ingress-nginx-controller-7979dcdd85-mw42w   1/1     Running   0          30s
        ........
        ```

### Install infrastructure components and ingress-rules via Helm and HelmOperator release in service namespace
* Define host name in the [./values/ingress-rules.helmrelease.yaml](./values/ingress-rules.helmrelease.yaml)
  ```
  ...
  spec:
    values:
      ingress:
        host: <hostname>
  ```
* Install components
  ```
  kubectl apply -n service -f ./values/ingress-rules.helmrelease.yaml
  helm install th2-infra-base -n service ./th2-service/ -f ./values/service.values.yaml -f ./secrets.yaml
  ```

Wait for all pods in service namespace are up and running, once completed proceed with [schema configuration](https://github.com/th2-net/th2-infra-schema-demo/blob/master/README.md) to deploy th2 namespaces.

## th2 infra links:
- Kubernetes dashboard http://your-host:30000/dashboard/
- Grafana http://your-host:30000/grafana/
- th2-infra-editor http://your-host:30000/editor/
- RabbitMQ http://your-host:30000/rabbitmq/
- th2-reports http://your-host:30000/your-namespace/

