# th2

![Version: 2.0.0](https://img.shields.io/badge/Version-2.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

th2 service Helm chart

## Requirements

| Repository | Name | Version |
|------------|------|---------|
|  | cassandra(cassandra) | 5.6.7 |
| https://charts.bitnami.com/bitnami | rabbitmq(rabbitmq) | 8.24.13 |
| https://charts.fluxcd.io | helmoperator(helm-operator) | 1.4.2 |
| https://kubernetes.github.io/dashboard/ | dashboard(kubernetes-dashboard) | 5.0.4 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| cassandra.cluster.datacenter | string | `"dc1"` |  |
| cassandra.cradle.cradleMaxEventBatchSize | string | `"1048576"` |  |
| cassandra.cradle.cradleMaxMessageBatchSize | string | `"1048576"` |  |
| cassandra.cradle.instanceName | string | `"th2-infra"` |  |
| cassandra.cradle.pageSize | int | `5000` |  |
| cassandra.cradle.timeout | int | `5000` |  |
| cassandra.dbUser.password | string | `""` | will be generated if empty |
| cassandra.dbUser.user | string | `"th2"` |  |
| cassandra.fullnameOverride | string | `"cassandra"` |  |
| cassandra.internal | bool | `true` | If service not internal - ExternalName service will be created, credentials will be mapped to secrets / config maps otherwise service will be deployed as a chart dependency |
| cassandra.keyspace | string | `"cradle"` |  |
| cassandra.persistence.enabled | bool | `false` |  |
| cassandra.persistence.size | string | `"50Gi"` |  |
| cassandra.persistence.storageClass | string | `"local-storage"` |  |
| clusterDomain | string | `"cluster.local"` | Kubernetes cluster domain |
| commonAnnotations | object | `{}` | Annotations will be added to all deployments, configMaps and ingresses that are created by th2-infra. |
| dashboard.image.repository | string | `"kubernetesui/dashboard"` |  |
| dashboard.ingress.annotations."kubernetes.io/ingress.class" | string | `"nginx"` |  |
| dashboard.ingress.annotations."nginx.ingress.kubernetes.io/configuration-snippet" | string | `"rewrite ^/([a-z\\-0-9]+)$ $scheme://$http_host/$1/ redirect;"` |  |
| dashboard.ingress.annotations."nginx.ingress.kubernetes.io/rewrite-target" | string | `"$1"` |  |
| dashboard.ingress.enabled | bool | `true` |  |
| dashboard.ingress.hosts[0] | string | `""` |  |
| dashboard.ingress.paths[0] | string | `"/dashboard($\|/.*)"` | dashboard UI ingess path |
| dashboard.internal | bool | `true` | Kubernetes dashboard values. If true - will be deployed as dependency |
| dashboard.protocolHttp | bool | `true` |  |
| dashboard.rbac.clusterRoleMetrics | bool | `true` |  |
| dashboard.rbac.create | bool | `true` |  |
| dashboard.serviceAccount.create | bool | `false` |  |
| dashboard.serviceAccount.name | string | `"th2infra-kubernetes-dashboard"` |  |
| helmoperator.chartsSyncInterval | string | `"300m"` |  |
| helmoperator.fullnameOverride | string | `"helm-operator"` |  |
| helmoperator.helm.versions | string | `"v3"` |  |
| helmoperator.internal | bool | `true` |  |
| helmoperator.nameOverrde | string | `"helm-operator"` |  |
| helmoperator.prometheus.enabled | string | `"enabled"` |  |
| helmoperator.prometheus.serviceMonitor.create | bool | `true` |  |
| hostAliases | list | `[]` | Adds host aliases to infra-mgr |
| infraEditor.image.repository | string | `"ghcr.io/th2-net/th2-infra-editor"` |  |
| infraEditor.image.tag | string | `"1.1.20"` |  |
| infraEditorV2.enabled | bool | `false` |  |
| infraEditorV2.image.repository | string | `"ghcr.io/th2-net/th2-infra-editor-v2"` |  |
| infraEditorV2.image.tag | string | `"2.1.1"` |  |
| infraGit.image.repository | string | `"ghcr.io/th2-net/git-ssh"` |  |
| infraGit.image.tag | string | `"v0.1.0"` |  |
| infraGit.internal | bool | `false` | If there is no ability to get access for cluster to repo with schemas namespaces configs git repo can be deployed internal in the cluster. Change to "internal = true" in case you need internal git repo. |
| infraGit.nodePort | int | `32600` |  |
| infraGit.persistence.accessModes | string | `"ReadWriteMany"` |  |
| infraGit.persistence.enabled | bool | `true` |  |
| infraGit.persistence.existingClaim | string | `""` | "repos-volume" claim will be created and mounted if empty |
| infraGit.persistence.storage | string | `"5Gi"` |  |
| infraGit.persistence.storageClassName | string | `"local-storage"` |  |
| infraGit.resources.limits.cpu | string | `"100m"` |  |
| infraGit.resources.limits.memory | string | `"200Mi"` |  |
| infraGit.resources.requests.cpu | string | `"50m"` |  |
| infraGit.resources.requests.memory | string | `"100Mi"` |  |
| infraMgr.cassandra.keyspacePrefix | string | `"th2_"` |  |
| infraMgr.cassandra.secret | string | `"cassandra"` |  |
| infraMgr.git.gitFetchInterval | string | `"14000"` | Infra manager git fetch interval in milliseconds |
| infraMgr.git.httpAuthPassword | string | `""` | Required for Git repository via HTTPS |
| infraMgr.git.httpAuthUsername | string | `""` | Required for Git repository via HTTPS |
| infraMgr.git.privateKey | string | `""` | Required for Git repository via SSH. Rsa private key in base64 encoding |
| infraMgr.git.repository | string | `"https://github.com/th2-net/th2-demo-configuration.git"` |  |
| infraMgr.git.repositoryLocalCache | string | `"/home/service/repository"` |  |
| infraMgr.git.secretName | string | `"infra-mgr"` |  |
| infraMgr.git.sshDir | string | `"/home/service/keys"` |  |
| infraMgr.image.repository | string | `"ghcr.io/th2-net/th2-infra-mgr"` |  |
| infraMgr.image.tag | string | `"2.2.1-infra-2.0-2955083181"` |  |
| infraMgr.jvm.javaToolOptions | string | `"-XX:+ExitOnOutOfMemoryError -XX:+UseContainerSupport -XX:MaxRAMPercentage=85"` |  |
| infraMgr.kubernetes.configMaps.cassandra | string | `"cradle"` |  |
| infraMgr.kubernetes.configMaps.cassandra-ext | string | `"cradle-external"` |  |
| infraMgr.kubernetes.configMaps.logging | string | `"logging-config-template"` |  |
| infraMgr.kubernetes.configMaps.prometheus | string | `"prometheus-app-config"` |  |
| infraMgr.kubernetes.configMaps.rabbitmq | string | `"rabbit-mq-app-config"` |  |
| infraMgr.kubernetes.configMaps.rabbitmq-ext | string | `"rabbit-mq-external-app-config"` |  |
| infraMgr.kubernetes.namespacePrefix | string | `"th2-"` | must be not more than 5 symbols |
| infraMgr.kubernetes.secrets | list | `[]` |  |
| infraMgr.livenessProbe.initialDelaySeconds | int | `30` |  |
| infraMgr.livenessProbe.periodSeconds | int | `30` |  |
| infraMgr.livenessProbe.timeoutSeconds | int | `5` |  |
| infraMgr.prometheusConfiguration.enabled | bool | `true` |  |
| infraMgr.prometheusConfiguration.port | int | `9752` |  |
| infraMgr.rabbitmq.passwordLength | int | `24` |  |
| infraMgr.rabbitmq.secret | string | `"rabbitmq"` |  |
| infraMgr.rabbitmq.usernamePrefix | string | `"th2-user-"` |  |
| infraMgr.rabbitmq.vHostPrefix | string | `"th2-"` |  |
| infraMgr.resources.limits.cpu | string | `"500m"` |  |
| infraMgr.resources.limits.memory | string | `"2500Mi"` |  |
| infraMgr.resources.requests.cpu | string | `"200m"` |  |
| infraMgr.resources.requests.memory | string | `"500Mi"` |  |
| infraOperator.config.chart.name | string | `"box-chart"` |  |
| infraOperator.config.chart.repository | string | `"http://infra-repo:8080"` |  |
| infraOperator.config.chart.version | string | `"1.0.0"` | Do not change the version |
| infraOperator.config.k8sUrl | string | `"<kubernetes-external-entrypoint>"` |  |
| infraOperator.config.namespacePrefixes[0] | string | `"th2-"` |  |
| infraOperator.config.rabbitMQManagement.exchangeName | string | `"global-notification"` |  |
| infraOperator.config.rabbitMQManagement.password | string | `"${RABBITMQ_PASS}"` |  |
| infraOperator.config.rabbitMQManagement.persistence | bool | `true` |  |
| infraOperator.config.rabbitMQManagement.schemaPermissions.configure | string | `""` |  |
| infraOperator.config.rabbitMQManagement.schemaPermissions.read | string | `".*"` |  |
| infraOperator.config.rabbitMQManagement.schemaPermissions.write | string | `".*"` |  |
| infraOperator.image.repository | string | `"ghcr.io/th2-net/th2-infra-operator"` |  |
| infraOperator.image.tag | string | `"4.3.5-infra-2.0-2972153075"` |  |
| infraOperator.jvm.javaToolOptions | string | `"-XX:+ExitOnOutOfMemoryError -XX:+UseContainerSupport -XX:MaxRAMPercentage=85"` |  |
| infraOperator.livenessProbe.initialDelaySeconds | int | `30` |  |
| infraOperator.livenessProbe.periodSeconds | int | `30` |  |
| infraOperator.livenessProbe.timeoutSeconds | int | `5` |  |
| infraOperator.prometheusConfiguration.enabled | bool | `true` |  |
| infraOperator.prometheusConfiguration.port | int | `9752` |  |
| infraOperator.resources.limits.cpu | string | `"800m"` |  |
| infraOperator.resources.limits.memory | string | `"1200Mi"` |  |
| infraOperator.resources.requests.cpu | string | `"200m"` |  |
| infraOperator.resources.requests.memory | string | `"500Mi"` |  |
| infraRepo.image.repository | string | `"ghcr.io/th2-net/infra-repo"` |  |
| infraRepo.image.tag | string | `"2.0.0@sha256:c7754eb8928a9002091537828705ba86d1efad54bb443f5760c9004606bc562c"` |  |
| ingress.annotations.infraNamespace | object | `{"nginx.ingress.kubernetes.io/configuration-snippet":"rewrite ^/([a-z\\-0-9]+)$ $scheme://$http_host/$1/ redirect;","nginx.ingress.kubernetes.io/enable-cors":"true","nginx.ingress.kubernetes.io/rewrite-target":"/$1","nginx.ingress.kubernetes.io/use-regex":"true"}` | Annotations for infra services |
| ingress.annotations.root | object | `{}` | Annotations for th2 root URL |
| ingress.annotations.th2Namespace | object | `{"nginx.ingress.kubernetes.io/enable-cors":"true","nginx.ingress.kubernetes.io/rewrite-target":"/$1","nginx.ingress.kubernetes.io/use-regex":"true"}` | Annotations for th2 schema services |
| ingress.host | string | `""` | Hostname for th2 namespace services |
| ingress.infraHost | string | `""` | Hostname for infra services. If not set, than host will be used |
| ingress.ingressClass | string | `"nginx"` | Hostname to access ingress |
| openshift | object | `{"enabled":false}` | Enable th2 for Openshift, impacts on Ingress. |
| prometheus.operator.enabled | bool | `true` | Set true if kube-prometheus-stack is used |
| prometheus.operator.serviceMonitor.name | string | `"pods-monitoring"` |  |
| prometheus.operator.serviceMonitor.selector.app | string | `"kube-prometheus-stack"` |  |
| prometheus.operator.serviceMonitor.selector.release | string | `"prometheus"` |  |
| rabbitmq.auth.erlangCookie | string | `""` |  |
| rabbitmq.auth.password | string | `""` |  |
| rabbitmq.auth.username | string | `"th2"` |  |
| rabbitmq.extraConfiguration | string | `"default_vhost = th2\ndisk_free_limit.absolute = 10GB\nprometheus.return_per_object_metrics=true"` |  |
| rabbitmq.fullnameOverride | string | `"rabbitmq"` |  |
| rabbitmq.ingress.annotations."kubernetes.io/ingress.class" | string | `"nginx"` |  |
| rabbitmq.ingress.annotations."nginx.ingress.kubernetes.io/configuration-snippet" | string | `"rewrite ^/([a-z\\-0-9]+)$ $scheme://$http_host/$1/ redirect;\nif ($request_uri ~ \"^/rabbitmq(/.*)\") {\n  proxy_pass http://upstream_balancer$1;\n  break;\n}\n"` |  |
| rabbitmq.ingress.annotations."nginx.ingress.kubernetes.io/rewrite-target" | string | `"$1"` |  |
| rabbitmq.ingress.enabled | bool | `true` |  |
| rabbitmq.ingress.hostname | string | `" "` |  |
| rabbitmq.ingress.path | string | `"/rabbitmq($\|/.*)"` | RabbitmMQ management UI ingess path |
| rabbitmq.internal | bool | `true` | If service not internal - ExternalName service will be created, credentials will be mapped to secrets / config maps otherwise service will be deployed as a chart dependency |
| rabbitmq.memoryHighWatermark.enabled | bool | `true` |  |
| rabbitmq.memoryHighWatermark.type | string | `"absolute"` |  |
| rabbitmq.memoryHighWatermark.value | string | `"2500MB"` |  |
| rabbitmq.metrics.enabled | bool | `true` |  |
| rabbitmq.metrics.serviceMonitor.additionalLabels.app | string | `"kube-prometheus-stack"` |  |
| rabbitmq.metrics.serviceMonitor.additionalLabels.release | string | `"prometheus"` |  |
| rabbitmq.metrics.serviceMonitor.enabled | bool | `true` |  |
| rabbitmq.metrics.serviceMonitor.path | string | `"/metrics"` |  |
| rabbitmq.persistence.enabled | bool | `true` |  |
| rabbitmq.persistence.size | string | `"10Gi"` |  |
| rabbitmq.persistence.storageClass | string | `"local-storage"` |  |
| rabbitmq.podAntiAffinityPreset | string | `"hard"` |  |
| rabbitmq.prometheus.operator.enabled | bool | `true` |  |
| rabbitmq.rabbitmqExchange | string | `"th2-exchange"` |  |
| rabbitmq.replicaCount | int | `1` |  |
| rabbitmq.resources.limits.memory | string | `"3000Mi"` |  |
| rabbitmq.service.managerNodePort | int | `32025` |  |
| rabbitmq.service.nodePort | int | `32000` |  |
| rabbitmq.service.type | string | `"NodePort"` |  |
| registries | object | `{}` | Registries are attached as pull secrets to pods |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.8.1](https://github.com/norwoodj/helm-docs/releases/v1.8.1)