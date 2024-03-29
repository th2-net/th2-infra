# th2 installation

## Prerequisites
Before you begin, please check the following prerequisites:
* Fully functioning Kubernetes cluster suitable for your bussiness needs, please refer to [technical requirements](https://github.com/th2-net/th2-documentation/wiki/Technical-Requirements) and [version compatibility](https://github.com/th2-net/th2-documentation/wiki/Software-compatibility-matrix)
* Operator-box that meets [hardware](https://github.com/th2-net/th2-documentation/wiki/Technical-Requirements) and [software](https://github.com/th2-net/th2-documentation/wiki/Technical-Requirements#software-requirements) requirements
* Installed [Apache Cassandra](https://cassandra.apache.org/) - [technical requirements](https://github.com/th2-net/th2-documentation/wiki/Technical-Requirements#apache-cassandra-cluster-hardware-requirements)

### All th2 components could be deployed by one of the following options:
1. Helm charts by [Helm](https://helm.sh/) and [Helm Operator](https://docs.fluxcd.io/projects/helm-operator/en/stable/). Please follow [steps](./README.md#installation-steps-for-th2-via-helm-chart).
2. If you have an openshift cluster installed you can use [ArgoCD](https://docs.openshift.com/container-platform/4.10/cicd/gitops/setting-up-argocd-instance.html). Steps are described into a separate [folder](./argocd/openshift).

***  
***  

## Installation steps for th2 via Helm chart
The following steps should be performed on the operator-box for th2-infra deployment:
<!--ts-->
   * [Download th2 git repositories](#th2-git-repository)
   * [Monitoring deployment](#monitoring-deployment)
   * [Cluster configuration](#cluster-configuration)
   * [th2 deployment](#th2-deployment)
<!--te-->

## th2 git repository
Installation of th2 infra requires a Git repository for maintaining th2 schema configuration. The information regarding this repository and its usage can be found in the guide further.

https://github.com/th2-net/th2-infra-schema-demo - can be used as a starter kit for schema repository.
[Template](https://docs.github.com/en/free-pro-team@latest/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template) or [fork](https://docs.github.com/en/free-pro-team@latest/github/getting-started-with-github/fork-a-repo#fork-an-example-repository) it.

[https://github.com/th2-net/th2-infra/example-values](https://github.com/th2-net/th2-infra/tree/master/example-values) - contains example values for th2 infra charts, we also recommend to store these values in a separate git repository

Clone th2-infra values repository into your operator-box:
```
$ git clone https://github.com/th2-net/th2-infra.git
```
change the current directory
```
$ cd ./th2-infra/example-values
```

## Infrastructure namespaces
Infrastructure components are split into two namespaces: _`monitoring`_ and _`service`_. These namespaces will be created below.

Next components of monitoring stack are deployed into _`monitoring`_ namespace:
* [grafana](https://grafana.com/oss/grafana/)
* [loki](https://grafana.com/oss/loki/)
* [prometheus](https://grafana.com/oss/prometheus/)

The _`service`_ namespace is used for infrastructure services:
* [kubernetes-dashboard](https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/)
* [RabbitMQ](https://www.rabbitmq.com/)
* [NGINX Ingress Controller](https://kubernetes.github.io/ingress-nginx/)
* [Helm Operator](https://github.com/fluxcd/helm-operator)
* [JupyterHub](https://github.com/jupyterhub/jupyterhub)

and for th2-infra components:
* [th2-infra-editor](https://github.com/th2-net/th2-infra-editor-v2)
* [th2-infra-operator](https://github.com/th2-net/th2-infra-operator)
* [th2-infra-mgr](https://github.com/th2-net/th2-infra-mgr)
* [th2-infra-repo](https://github.com/th2-net/infra-operator-tpl)

The following picture describes a cluser with monitoring stack, th2-infra and th2 namespace:

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
$ mkdir /opt/grafana /opt/prometheus /opt/loki /opt/rabbitmq /opt/jupiter_users /opt/jupiter_db /opt/jupiter_shared
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

* Define Grafana host names (the name must be resolved from QA boxes) in the [prometheus-operator.values.yaml](./example-values/prometheus-operator.values.yaml) file
```
grafana:
  ingress:
    hosts:
      - <th2_host_name>
```

* Deploy components
```
$ helm repo add grafana https://grafana.github.io/helm-charts
$ helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
$ helm install --version=2.8.3 loki -n monitoring grafana/loki-stack -f ./loki.values.yaml
$ helm install --version=41.4.0 prometheus -n monitoring prometheus-community/kube-prometheus-stack -f ./prometheus-operator.values.yaml
```
* Check result:
```
$ kubectl get pods
NAME                                                     READY   STATUS    RESTARTS   AGE
........
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
* Check access to Grafana _(default user/password: `admin/prom-operator`. Must be changed)_: <br>
  http://your-host:30000/grafana/login

## Cluster configuration
Once all of the required software is installed on your test-box and operator-box and th2-infra repositories are ready you can start configuring the cluster.

### Access for infra-mgr th2 schema git repository:

`ssh` or `https` access with write permissions is required by **th2-infra-mgr** component

#### Set up __ssh__ access

* Generate keys without passphrase  
```
$ ssh-keygen -t rsa -m pem -f ./infra-mgr-rsa.key
```

* Get key in base64 format and put in infraMgr.git.privateKey
```
$ base64 -w 0 ./infra-mgr-rsa.key
```
* [Add a new deploy key to your schema repository on GitHub with read and write permissions](https://docs.github.com/en/developers/overview/managing-deploy-keys#deploy-keys)

#### Set up __https__ access

* [Generate access token for schema repository with read and write permissions](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)
* Set up values in secrets.yaml file (described below)

### Access for converter th2 schema git repository:

`ssh` or `https` access with write permissions is required by **th2-converter** component

#### Set up __ssh__ access

* Generate keys without passphrase  
```
$ ssh-keygen -t ed25519 -m pem -f ./converter-ed25519.key
```

* Get key in base64 format and put in converter.git.privateKey
```
$ base64 -w 0 ./converter-ed25519.key
```
* [Add a new deploy key to your schema repository on GitHub with read and write permissions](https://docs.github.com/en/developers/overview/managing-deploy-keys#deploy-keys)
* Write permission are required by infra-manager to complete Dynamic Resource Monitoring and by converter to be able to push converted schemas on git

### Set the repository with schema configuration
set `infraMgr.git.repository` value in the [service.values.yaml](./example-values/service.values.yaml) file to link of your schema repository, `ssh` or `https`:
* **ssh**
```
infraMgr:
  git:
    repository: git@github.com:th2-net/th2-infra-demo-configuration.git
converter:
  git:
    repository: git@github.com:th2-net/th2-infra-demo-configuration.git
```
* **https**
```
infraMgr:
  git:
    repository: https://github.com/th2-net/th2-infra-schema-demo.git
converter:
  git:
    repository: https://github.com/th2-net/th2-infra-schema-demo.git
```

### Define cassandra host name
* set `cassandra.host` value for cassandra in the [service.values.yaml](./example-values/service.values.yaml) file.
```
cassandra:
  internal: false
  host: <cassandra-host>
```

### Define th2 ingress parameters
* Add `ingress.hostname` value if required into [service.values.yaml](./example-values/service.values.yaml) file otherwise th2 http services will be available on node IP address
```
ingress:
  host: example.com
...
kubernetes-dashboard:
  ingress:
    hosts:
    - example.com
...
rabbitmq:
  ingress:
    hostname: example.com
...
jupyterhub:
  ingress:
    ingressClassName: nginx
    hosts:
    - example.com
```

### Create secret with th2 credentials

Create secrets.yaml in `./` folder (*do not commit into git*). Example:
```
# reguired only for images from a private registry, will be attached as the first PullSecret to deployments
#registries:
#  registry1.name.com:8080:
#    username: <username>
#    password: <password>
#  registry2.name.com:8080:
#    username: <username>
#    password: <password>

cassandra:
# set credentials for existing Cassandra cluster
  dbUser:
    user: <user-name>
    password: <password>

rabbitmq:
# set admin user credentials, it will be created during deployment
  auth:
    username: th2
    password: rab-pass
    # must be random string
    erlangCookie: cookie

# required if http(s) access to gitlab/github repositories is used
#infraMgr:
#  git:
#    privateKey: <private key in base64>
#    httpAuthUsername: username
#    # authentication username
#    # when using token auth for GitLab it should be equal to "oauth2"
#    # when using token auth for GitHub it should be equal to token itself
#    httpAuthPassword:
#    # authentication password
#    # when using token auth for GitLab it should be equal to token itself
#    # when using token auth for GitHub it should be equal to empty string

# required if http(s) access to gitlab/github repositories is used
#converter:
#  git:
#    privateKey: <private key in base64>
#    httpAuthUsername: username
#    # authentication username
#    # when using token auth for GitLab it should be equal to "oauth2"
#    # when using token auth for GitHub it should be equal to token itself
#    httpAuthPassword: 
#    # authentication password
#    # when using token auth for GitLab it should be equal to token itself
#    # when using token auth for GitHub it should be equal to empty string

jupyterhub:
# set credentials for admin and other users
  hub:
    config:
      Authenticator:
        admin_users:
          - <admin-username>
        allowed_users:
          - <username>
      DummyAuthenticator:
        password: <password>
```
### infra-git deployment

If you have any restrictions to get access to any external repositories from the k8s cluster git service can be deployed according to the following instruction:

*  Create PersistentVolume "repos-volume", example is presented in the ./example-values/persistence/pv.yaml or your own PVC
*  Create configmap "keys-repo" from public part of key from point "Access for infra-mgr th2 schema git repository":
```
$ kubectl -n service create configmap keys-repo --from-file=authorized_keys=./infra-mgr-rsa.pub
```
*  Create configmap "keys-repo-converter" from public part of key from point "Access for converter th2 schema git repository":
```
$ kubectl -n service create configmap keys-repo-converter --from-file=authorized_keys=./converter-ed25519.pub
```
*  Define configs for infra-git in services.values.yaml
*  set `infraMgr.git.repository` and `converter.git.repository` value in the service.values.yaml file to **ssh** link of your repository, e.g:
```
infraMgr:
  git:
    repository: ssh://git@infra-git/home/git/repo/schema.git
converter:
  git:
    repository: ssh://git@infra-git/home/git/repo/schema.git
```
* after installation you should create folder with the same path and the name ```schema.git``` that you define in previous step inside infra-git pod and initialise it as a git repo.
```
$ su git
$ mkdir /home/git/repo/schema.git
$ cd /home/git/repo/schema.git
$ git init --bare
```

* to connect to your created repository add this host to your ~/.ssh/config
```
Host <node-address>
    User git
    Port 32600
    IdentityFile ~/path_to_private_key/infra-mgr-rsa.key
```
* clone your infra-git repo using
```
$ git clone git@<node-address>:/home/git/repo/schema.git
```
## th2 deployment
### Install NGINX Ingress Controller
```
$ helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
$ helm install -n service --version=4.3.0 ingress ingress-nginx/ingress-nginx -f ./ingress.values.yaml
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
$ helm install -n service --version=<version> th2-infra th2/th2 -f ./service.values.yaml -f ./secrets.yaml
```
_Note_: replace <version> with th2-infra release version you need, please follow to https://github.com/th2-net/th2-infra/releases

Wait for all pods in service namespace are up and running, once completed proceed with [schema configuration](https://github.com/th2-net/th2-infra-schema-demo/blob/master/README.md) to deploy th2 namespaces.

### Upgrade th2-infra

* Purge th2 namespaces and uninstall th2-infra Helm release
```
$ helm -n service uninstall th2-infra
```
remove th2 namespaces
```
$ kubectel delete <namespace-1> <namespace-2> <namespace-..>
```
or set "deny" in "infra-mgr-config.yml" file for all namespaces managed by th2-infra. Wait until it is removed, once completed uninstall th2-infra release
```
$ helm -n service uninstall th2-infra
```
* Delete CRDs:
```
$ kubectl delete customresourcedefinitions th2boxes.th2.exactpro.com th2coreboxes.th2.exactpro.com th2dictionaries.th2.exactpro.com th2estores.th2.exactpro.com th2jobs.th2.exactpro.com th2mstores.th2.exactpro.com
```
 _Note_: the list can be various, see the full list in documentation or in k8s with the following command:
```
$ kubectl get customresourcedefinitions | grep "^th2"
```
* Change service.values.yaml if it is required by th2-infra release notes
* Revise "Custom Resource" files for namespaces if it is required by th2-infra release notes
* Install th2-infra:
```
$ helm repo update
$ helm install -n service --version=<new_version> th2-infra th2/th2 -f ./service.values.yaml -f ./secrets.yaml
```
_Note_: replace <new_version> with th2-infra release version you need, please follow to https://github.com/th2-net/th2-infra/release

### Re-adding persistence for components in th2 namespaces
PersistentVolumeClaim is namespace scoped resource, so after namespace re-creation PVCs should be added for components require persistence.
* Check the state of PV in a cluster:
```
$ kubectl get pv
```
* Reset PVs that are in Released status:
```
$ kubectl patch pv <pv-name> -p '{"spec":{"claimRef": null}}'
```
* Apply PVCs
```
$ kubectl -n <th2-namespace> apply -f ./pvc.yaml
```
_Note_: replace <th2-namespace> with th2 namespace you use

## th2 infra links:
- Main page http://your-host:30000/
- Kubernetes dashboard http://your-host:30000/dashboard/
- Grafana http://your-host:30000/grafana/
- RabbitMQ http://your-host:30000/rabbitmq/
- th2-reports http://your-host:30000/your-namespace/
- Jupiter hub http://your-host:30000/jupyterhub/

## Migration to v2.1.0 th2-infra chart 
Follow to migration guide with link above [MIGRATION](docs/MIGRATION.md)
