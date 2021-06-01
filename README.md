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

## th2 Git repository
Installation of th2 infra requires a Git repository for maintaining th2 schema configuration. The information regarding this repository and its usage can be found in the guide further.
* https://github.com/th2-net/th2-infra-schema-demo - can be used as a starter kit for schema repository
* https://github.com/th2-net/th2-infra/example-values - can be used as a starter kit for th2 infra, we also recommend to store these values in a separate git repository

The first step that should be done in the th2 deployment process is copying th2-infra repository into your operator-box:
```
$ git clone https://github.com/th2-net/th2-infra.git
```
change the current directory
```
$ cd ./th2-infra/example-values
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

![k8s cluster](https://user-images.githubusercontent.com/690243/101762881-0925d080-3aef-11eb-9d15-70e9277b0fa5.jpg)

### Create namespaces:
```
$ kubectl create namespace monitoring
namespace/monitoring created
$ kubectl create namespace service
namespace/service created
```

## Data persistence

Data persistence is required for the following components: Grafana, Prometheus, Loki, RabbitMQ components and should be set up on this step.

_Note: Examples below use HostPath type of [Persistent Volume(PV)](https://kubernetes.io/docs/concepts/storage/persistent-volumes/). Please read the documentation to choose an appropriate PV type for your environment_

* the following command can require root permissions, create directory on th2 node:
```
$ mkdir /opt/grafana /opt/prometheus /opt/loki /opt/rabbitmq
```
* set node name in `pvs.yaml`
* create PVs and PVCs:
```
$ kubectl apply -f ./pvs.yaml
$ kubectl apply -f ./pvcs.yaml
```

If you would like to include th2 read components into your configuration, you also have to set up a dedicated PersistentVolume for th2-read log directory.
You should add PersistentVolume mapped to /opt/components directory and then create PersistentVolumeClaim once a schema namespace installed. PV and PVC examples can be found here [persistence/](./example-values/persistence/)

```
$ mkdir /opt/components
```
* set node name in `persistence/pv.yaml`
* create PV:
```
$ kubectl apply -f ./persistence/pv.yaml
```
* create PVC:
```
$ kubectl apply -f ./persistence/pvc.yaml
```

Details for th2-read-log [README.md](https://github.com/th2-net/th2-read-log#configuration)

## Monitoring deployment

_Note: It's an optional step, but it gets slightly simpler checking the result of installation. In all installation commands we explicitly define namespaces to avoid possible mistakes._
* Switch namespace to monitoring
```
$ kubectl config set-context --current --namespace=monitoring
```
* Define Grafana and Dashboard host names (the name must be resolved from QA boxes):
  * in the [dashboard.values.yaml](./example-values/dashboard.values.yaml) file
    ```
    ingress:
      hosts:
        - <th2_host_name>
    ```
  * in the [prometheus-operator.values.yaml](./example-values/prometheus-operator.values.yaml) file
    ```
    grafana:
      ingress:
        hosts:
          - <th2_host_name>
    ```

* Install [Kubernetes Dashboard](https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/)
```
$ helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
$ helm install dashboard -n monitoring kubernetes-dashboard/kubernetes-dashboard -f ./dashboard.values.yaml
```
* Deploy components
```
$ helm repo add loki https://grafana.github.io/loki/charts
$ helm repo add stable https://charts.helm.sh/stable
$ helm upgrade --install loki --namespace=monitoring loki/loki-stack -f ./loki.values.yaml
$ helm upgrade --install prometheus stable/prometheus-operator -n monitoring -f ./prometheus-operator.values.yaml
```
* Check result:
```
$ kubectl get pods
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

Add loki Datasource as http://loki:3100 and import Dashboard from components-logs.json and RabbitMQ Overview from here: https://grafana.com/grafana/dashboards/10991

* Check access to Grafana _(default user/password: `admin/prom-operator`. Must be changed)_: <br>
  http://your-host:30000/grafana/login

## Cluster configuration
Once all of the required software is installed on your test-box and operator-box and th2-infra repositories are ready you can start configuring the cluster.

* Switch namespace to service:
```
$ kubectl config set-context --current --namespace=service
```

### Access for infra-mgr th2 schema git repository:

`ssh` access with write permissions is required by **th2-infra-mgr** component

* Generate keys without passphrase  
```
$ ssh-keygen -t rsa -m pem -f ./infra-mgr-rsa.key
``` 
* [Add a new SSH key to your GitHub account](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/adding-a-new-ssh-key-to-your-github-account)
* Create infra-mgr secret from the private key:
```
$ kubectl -n service create secret generic infra-mgr --from-file=infra-mgr=./infra-mgr-rsa.key
```

### Set the repository with schema configuration
* set `infraMgr.git.repository` value in the [service.values.yaml](./example-values/service.values.yaml) file to **ssh** link of your schema repository, e.g:
```
infraMgr:
  git:
    repository: git@github.com:th2-net/th2-infra-demo-configuration.git
```

### Define cassandra host name
* set `cassandra.host` value for cassandra in the [service.values.yaml](./example-values/service.values.yaml) file.
```
cassandra:
  internal: false
  host: <cassandra-host>
```

### Define th2 ingress hostname
Add `ingress.hostname` value if required into [service.values.yaml](./example-values/service.values.yaml) file otherwise th2 http services will be available on node IP address
```
ingress:
  host: example.com
```

### Create secret with th2 credentials

Create secrets.yaml in `./` folder (*do not commit into git*). Example:
```
# reguired only for images from a private registry, will be attached as the first PullSecret to deployments
#productRegistry:
#  username: user
#  password: password
#  name: private-registry-1.example.com # core components registry

# reguired only for images from a private registry, will be attached as the second PullSecret to deployments
#solutionRegistry:
#  username: user
#  password: password
#  name: private-registry-2.example.com # components registry

# reguired only for images from a private registry, will be attached as the third PullSecret to deployments
#proprietaryRegistry:
#  username: user
#  password: password
#  name: private-registry-3.example.com # components registry

cassandra:
# set credentials for existing Cassandra cluster
  dbUser:
    user: <user-name>
    password: <password>

rabbitmq:
# set admin user credentials, it will be created during deployment
  rabbitmqUsername: th2
  rabbitmqPassword: rab-pass
  # must be random string
  rabbitmqErlangCookie: cookie
```

## th2 deployment
### Install NGINX Ingress Controller
```
$ helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
$ helm install -n service --version=3.12.0 ingress ingress-nginx/ingress-nginx -f ./ingress.values.yaml
```
Check:
```
$ kubectl get pods
NAME                                                READY   STATUS    RESTARTS   AGE
........
ingress-ingress-nginx-controller-7979dcdd85-mw42w   1/1     Running   0          30s
........
```

### Install th2-infra components in the service namespace
```
$ helm repo add th2 https://th2-net.github.io
$ helm install -n service --version=<version> th2-infra-base th2/th2 -f ./service.values.yaml -f ./secrets.yaml
```

Wait for all pods in service namespace are up and running, once completed proceed with [schema configuration](https://github.com/th2-net/th2-infra-schema-demo/blob/master/README.md) to deploy th2 namespaces.

### Upgrade/migration th2-infra

* Set "deny" in "infra-mgr-config.yml" file for all namespaces managed by th2 to delete it.
* Uninstall th2-infra-base release:
```
$ helm -n service uninstall th2-infra-base
```
* Revise "Custom resource" files for namespaces according to the release documentation (if required).
* Delete CRDs:
```
$ kubectl delete customresourcedefinitions th2boxes.th2.exactpro.com th2coreboxes.th2.exactpro.com th2dictionaries.th2.exactpro.com th2estores.th2.exactpro.com th2links.th2.exactpro.com th2mstores.th2.exactpro.com
```
 _Note_: the list can be various, see the full list in documentation or in k8s with the following command:
```
$ kubectl get customresourcedefinitions | grep "^th2"
```
* Change th2-service values file according to the th2-infra release notes (if required):
* Check the state of pv in k8s.
```
$ kubectl get pv
```
It has to be availalble, if Released:
```
$ kubectl patch pv <pv-name> -p '{"spec":{"claimRef": null}}'
```
  
* Install th2-infra:
```
$ helm repo add th2 https://th2-net.github.io
$ helm install -n service --version=<new_version> th2-infra-base th2/th2 -f ./service.values.yaml -f ./secrets.yaml
```

* Apply PVC in th2 namespaces (if required)

## th2 infra links:
- Kubernetes dashboard http://your-host:30000/dashboard/
- Grafana http://your-host:30000/grafana/
- th2-infra-editor http://your-host:30000/editor/
- RabbitMQ http://your-host:30000/rabbitmq/
- th2-reports http://your-host:30000/your-namespace/
