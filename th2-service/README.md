# th2

![Version: 1.7.2](https://img.shields.io/badge/Version-1.7.2-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

th2 service Helm chart

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.bitnami.com/bitnami | cassandra(cassandra) | 5.6.7 |
| https://charts.fluxcd.io | helmoperator(helm-operator) | 1.2.0 |
| https://charts.helm.sh/stable | rabbitmq(rabbitmq-ha) | 1.44.4 |
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
| clusterDomain | string | `"cluster.local"` |  |
| commonAnnotations | object | `{}` | Annotations will be added to all deployments, configMaps and ingresses that are created by th2-infra. |
| dashboard.image.repository | string | `"kubernetesui/dashboard"` |  |
| dashboard.ingress.annotations."kubernetes.io/ingress.class" | string | `"nginx"` |  |
| dashboard.ingress.annotations."nginx.ingress.kubernetes.io/configuration-snippet" | string | `"rewrite ^/([a-z\\-0-9]*)$ $scheme://$http_host/$1/ redirect;"` |  |
| dashboard.ingress.annotations."nginx.ingress.kubernetes.io/rewrite-target" | string | `"$1"` |  |
| dashboard.ingress.enabled | bool | `true` |  |
| dashboard.ingress.hosts[0] | string | `""` |  |
| dashboard.ingress.paths[0] | string | `"/dashboard($|/.*)"` |  |
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
| helmoperator.prometheus.serviceMonitor.namespace | string | `"monitoring"` |  |
| infraEditor.image.repository | string | `"ghcr.io/th2-net/th2-infra-editor"` |  |
| infraEditor.image.tag | string | `"1.0.65"` |  |
| infraGit.image.repository | string | `"ghcr.io/th2-net/git-ssh"` |  |
| infraGit.image.tag | string | `"v0.1.0"` |  |
| infraGit.internal | bool | `false` | If there is no ability to get access for cluster to repo with schemas namespaces configs git repo can be deployed internal in the cluster. Change to "internal = true" in case you need internal git repo. |
| infraGit.nodePort | int | `32600` |  |
| infraGit.pvc.accessModes | string | `"ReadWriteMany"` |  |
| infraGit.pvc.storage | string | `"5Gi"` |  |
| infraGit.pvc.storageClassName | string | `"local-storage"` |  |
| infraGit.resources.limits.cpu | string | `"100m"` |  |
| infraGit.resources.limits.memory | string | `"200Mi"` |  |
| infraGit.resources.requests.cpu | string | `"50m"` |  |
| infraGit.resources.requests.memory | string | `"100Mi"` |  |
| infraMgr.cassandra.keyspacePrefix | string | `"schema_"` |  |
| infraMgr.cassandra.secret | string | `"cassandra"` |  |
| infraMgr.git.gitFetchInterval | string | `"14000"` | Infra manager git fetch interval in milliseconds |
| infraMgr.git.httpAuthPassword | string | `""` |  |
| infraMgr.git.httpAuthUsername | string | `""` |  |
| infraMgr.git.privateKeyFileSecret | string | `"infra-mgr"` |  |
| infraMgr.git.repository | string | `"git@github.com:th2-net/th2-demo-configuration.git"` |  |
| infraMgr.git.repositoryLocalCache | string | `"/home/service/repository"` |  |
| infraMgr.git.secretMountPath | string | `"/home/service/keys"` |  |
| infraMgr.git.secretName | string | `"infra-mgr"` |  |
| infraMgr.image.repository | string | `"ghcr.io/th2-net/th2-infra-mgr"` |  |
| infraMgr.image.tag | string | `"1.6.1"` |  |
| infraMgr.jvm.javaToolOptions | string | `"-XX:+ExitOnOutOfMemoryError -XX:+UseContainerSupport -XX:MaxRAMPercentage=85"` |  |
| infraMgr.kubernetes.configMaps.cassandra | string | `"cradle"` |  |
| infraMgr.kubernetes.configMaps.cassandra-ext | string | `"cradle-external"` |  |
| infraMgr.kubernetes.configMaps.logging | string | `"logging-config-template"` |  |
| infraMgr.kubernetes.configMaps.prometheus | string | `"prometheus-app-config"` |  |
| infraMgr.kubernetes.configMaps.rabbitmq | string | `"rabbit-mq-app-config"` |  |
| infraMgr.kubernetes.configMaps.rabbitmq-ext | string | `"rabbit-mq-external-app-config"` |  |
| infraMgr.kubernetes.ingress | string | `"ingress-rules"` |  |
| infraMgr.kubernetes.namespacePrefix | string | `"th2-"` | must be not more than 5 symbols |
| infraMgr.kubernetes.secrets[0] | string | `"th2-core"` |  |
| infraMgr.kubernetes.secrets[1] | string | `"th2-solution"` |  |
| infraMgr.kubernetes.secrets[2] | string | `"th2-proprietary"` |  |
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
| infraOperator.config.chart.name | string | `"infra-operator-tpl"` |  |
| infraOperator.config.chart.repository | string | `"http://infra-repo:8080"` |  |
| infraOperator.config.chart.version | string | `"0.7.2"` |  |
| infraOperator.config.k8sUrl | string | `"<kubernetes-external-entrypoint>"` |  |
| infraOperator.config.namespacePrefixes[0] | string | `"th2-"` |  |
| infraOperator.config.rabbitMQManagement.password | string | `"${RABBITMQ_PASS}"` |  |
| infraOperator.config.rabbitMQManagement.persistence | bool | `true` |  |
| infraOperator.config.rabbitMQManagement.port | string | `"15672"` |  |
| infraOperator.config.rabbitMQManagement.schemaPermissions.configure | string | `""` |  |
| infraOperator.config.rabbitMQManagement.schemaPermissions.read | string | `".*"` |  |
| infraOperator.config.rabbitMQManagement.schemaPermissions.write | string | `".*"` |  |
| infraOperator.config.rabbitMQManagement.username | string | `"th2"` |  |
| infraOperator.image.repository | string | `"ghcr.io/th2-net/th2-infra-operator"` |  |
| infraOperator.image.tag | string | `"3.5.2"` |  |
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
| infraRepo.image.tag | string | `"0.7.2"` |  |
| ingress.annotations.default."kubernetes.io/ingress.class" | string | `"nginx"` |  |
| ingress.annotations.default."nginx.ingress.kubernetes.io/configuration-snippet" | string | `"rewrite ^/([a-z\\-0-9]*)$ $scheme://$http_host/$1/ redirect;"` |  |
| ingress.annotations.default."nginx.ingress.kubernetes.io/enable-cors" | string | `"true"` |  |
| ingress.annotations.default."nginx.ingress.kubernetes.io/rewrite-target" | string | `"/$1"` |  |
| ingress.annotations.default."nginx.ingress.kubernetes.io/use-regex" | string | `"true"` |  |
| ingress.annotations.extra | object | `{}` | Annotations will be added to all ingress-rules generated by th2-infra |
| ingress.host | string | `""` |  |
| productRegistry.name | string | `""` | Image repository and credentials to create pull secrets |
| productRegistry.password | string | `""` |  |
| productRegistry.secret | string | `"th2-core"` |  |
| productRegistry.username | string | `""` |  |
| prometheus.operator.alerts.selector | object | `{}` |  |
| prometheus.operator.enabled | bool | `true` |  |
| prometheus.operator.serviceMonitor.namespace | string | `"monitoring"` | Namespace to install ServiceMonitor |
| prometheus.operator.serviceMonitor.selector.app | string | `"kube-prometheus-stack"` |  |
| prometheus.operator.serviceMonitor.selector.prometheus | string | `"kube-prometheus"` |  |
| prometheus.operator.serviceMonitor.selector.release | string | `"prometheus"` |  |
| proprietaryRegistry.password | string | `""` |  |
| proprietaryRegistry.registry | string | `""` |  |
| proprietaryRegistry.secret | string | `"th2-proprietary"` |  |
| proprietaryRegistry.username | string | `""` |  |
| rabbitmq.extraConfig | string | `"disk_free_limit.absolute = 10GB"` |  |
| rabbitmq.fullnameOverride | string | `"rabbitmq"` |  |
| rabbitmq.ingress.annotations."kubernetes.io/ingress.class" | string | `"nginx"` |  |
| rabbitmq.ingress.annotations."nginx.ingress.kubernetes.io/configuration-snippet" | string | `"rewrite ^/([a-z\\-0-9]*)$ $scheme://$http_host/$1/ redirect;\nif ($request_uri ~ \"^/rabbitmq(/.*)\") {\n  proxy_pass http://upstream_balancer$1;\n  break;\n}\n"` |  |
| rabbitmq.ingress.annotations."nginx.ingress.kubernetes.io/rewrite-target" | string | `"$1"` |  |
| rabbitmq.ingress.enabled | bool | `true` |  |
| rabbitmq.ingress.hostName | string | `""` |  |
| rabbitmq.ingress.path | string | `"/rabbitmq($|/.*)"` |  |
| rabbitmq.internal | bool | `true` | If service not internal - ExternalName service will be created, credentials will be mapped to secrets / config maps otherwise service will be deployed as a chart dependency |
| rabbitmq.livenessProbe.exec.command[0] | string | `"/bin/bash"` |  |
| rabbitmq.livenessProbe.exec.command[1] | string | `"-ec"` |  |
| rabbitmq.livenessProbe.exec.command[2] | string | `"rabbitmq-diagnostics -q check_running"` |  |
| rabbitmq.livenessProbe.failureThreshold | int | `6` |  |
| rabbitmq.livenessProbe.initialDelaySeconds | int | `120` |  |
| rabbitmq.livenessProbe.periodSeconds | int | `30` |  |
| rabbitmq.livenessProbe.successThreshold | int | `1` |  |
| rabbitmq.livenessProbe.timeoutSeconds | int | `20` |  |
| rabbitmq.persistentVolume.enabled | bool | `true` |  |
| rabbitmq.persistentVolume.size | string | `"10Gi"` |  |
| rabbitmq.persistentVolume.storageClass | string | `"local-storage"` |  |
| rabbitmq.podAntiAffinity | string | `"hard"` |  |
| rabbitmq.prometheus.exporter.enabled | bool | `false` |  |
| rabbitmq.prometheus.operator.alerts.enabled | bool | `true` |  |
| rabbitmq.prometheus.operator.alerts.selector | bool | `false` |  |
| rabbitmq.prometheus.operator.enabled | bool | `true` |  |
| rabbitmq.prometheus.operator.serviceMonitor.selector.app | string | `"kube-prometheus-stack"` |  |
| rabbitmq.prometheus.operator.serviceMonitor.selector.release | string | `"prometheus"` |  |
| rabbitmq.rabbitmqErlangCookie | string | `""` |  |
| rabbitmq.rabbitmqExchange | string | `"th2-exchange"` |  |
| rabbitmq.rabbitmqMemoryHighWatermark | string | `"1024MB"` |  |
| rabbitmq.rabbitmqPassword | string | `""` | Will be generated if empty |
| rabbitmq.rabbitmqPrometheusPlugin.enabled | bool | `true` |  |
| rabbitmq.rabbitmqUsername | string | `"th2"` |  |
| rabbitmq.rabbitmqVhost | string | `"th2"` |  |
| rabbitmq.readinessProbe.exec.command[0] | string | `"/bin/bash"` |  |
| rabbitmq.readinessProbe.exec.command[1] | string | `"-ec"` |  |
| rabbitmq.readinessProbe.exec.command[2] | string | `"rabbitmq-diagnostics -q check_running"` |  |
| rabbitmq.readinessProbe.failureThreshold | int | `3` |  |
| rabbitmq.readinessProbe.initialDelaySeconds | int | `10` |  |
| rabbitmq.readinessProbe.periodSeconds | int | `30` |  |
| rabbitmq.readinessProbe.successThreshold | int | `1` |  |
| rabbitmq.readinessProbe.timeoutSeconds | int | `20` |  |
| rabbitmq.replicaCount | int | `1` |  |
| rabbitmq.service.amqpNodePort | int | `32000` |  |
| rabbitmq.service.type | string | `"NodePort"` |  |
| solutionRegistry.password | string | `""` |  |
| solutionRegistry.registry | string | `""` |  |
| solutionRegistry.secret | string | `"th2-solution"` |  |
| solutionRegistry.username | string | `""` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.7.0](https://github.com/norwoodj/helm-docs/releases/v1.7.0)
