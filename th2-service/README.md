# th2

![Version: 1.7.1](https://img.shields.io/badge/Version-1.7.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

th2 service Helm chart

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.bitnami.com/bitnami | cassandra(cassandra) | 5.6.7 |
| https://charts.helm.sh/stable | rabbitmq(rabbitmq-ha) | 1.44.4 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| cassandra.cluster.datacenter | string | `"dc1"` |  |
| cassandra.cradle.cradleMaxEventBatchSize | string | `"1048576"` |  |
| cassandra.cradle.cradleMaxMessageBatchSize | string | `"1048576"` |  |
| cassandra.cradle.instanceName | string | `"th2-infra"` |  |
| cassandra.cradle.pageSize | int | `5000` |  |
| cassandra.cradle.timeout | int | `5000` |  |
| cassandra.dbUser.password | string | `""` |  |
| cassandra.dbUser.user | string | `"th2"` |  |
| cassandra.fullnameOverride | string | `"cassandra"` |  |
| cassandra.internal | bool | `true` |  |
| cassandra.keyspace | string | `"cradle"` |  |
| cassandra.persistence.enabled | bool | `false` |  |
| cassandra.persistence.size | string | `"50Gi"` |  |
| cassandra.persistence.storageClass | string | `"local-storage"` |  |
| infraEditor.image.repository | string | `"ghcr.io/th2-net/th2-infra-editor"` |  |
| infraEditor.image.tag | string | `"1.0.65"` |  |
| infraGit.image.repository | string | `"ghcr.io/th2-net/git-ssh"` |  |
| infraGit.image.tag | string | `"v0.1.0"` |  |
| infraGit.internal | bool | `false` |  |
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
| infraMgr.git.httpAuthPassword | string | `""` |  |
| infraMgr.git.httpAuthUsername | string | `""` |  |
| infraMgr.git.privateKeyFileSecret | string | `"infra-mgr"` |  |
| infraMgr.git.repository | string | `"git@github.com:th2-net/th2-demo-configuration.git"` |  |
| infraMgr.git.repositoryLocalCache | string | `"/home/service/repository"` |  |
| infraMgr.git.secretMountPath | string | `"/home/service/keys"` |  |
| infraMgr.git.secretName | string | `"infra-mgr"` |  |
| infraMgr.image.repository | string | `"ghcr.io/th2-net/th2-infra-mgr"` |  |
| infraMgr.image.tag | string | `"1.5.1-infra-1.7.1-1708594635"` |  |
| infraMgr.jvm.javaToolOptions | string | `"-XX:+ExitOnOutOfMemoryError -XX:+UseContainerSupport -XX:MaxRAMPercentage=85"` |  |
| infraMgr.kubernetes.configMaps.cassandra | string | `"cradle"` |  |
| infraMgr.kubernetes.configMaps.cassandra-ext | string | `"cradle-external"` |  |
| infraMgr.kubernetes.configMaps.logging | string | `"logging-config-template"` |  |
| infraMgr.kubernetes.configMaps.prometheus | string | `"prometheus-app-config"` |  |
| infraMgr.kubernetes.configMaps.rabbitmq | string | `"rabbit-mq-app-config"` |  |
| infraMgr.kubernetes.configMaps.rabbitmq-ext | string | `"rabbit-mq-external-app-config"` |  |
| infraMgr.kubernetes.ingress | string | `"ingress-rules"` |  |
| infraMgr.kubernetes.namespacePrefix | string | `"th2-"` |  |
| infraMgr.kubernetes.secrets[0] | string | `"th2-core"` |  |
| infraMgr.kubernetes.secrets[1] | string | `"th2-solution"` |  |
| infraMgr.kubernetes.secrets[2] | string | `"th2-proprietary"` |  |
| infraMgr.kubernetes.secrets[3] | string | `"git-chart-creds"` |  |
| infraMgr.prometheusConfiguration.enabled | bool | `true` |  |
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
| infraOperator.config.chart.version | string | `"0.7.1"` |  |
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
| infraOperator.image.tag | string | `"3.4.1-infra-1.7.1-1712191378"` |  |
| infraOperator.jvm.javaToolOptions | string | `"-XX:+ExitOnOutOfMemoryError -XX:+UseContainerSupport -XX:MaxRAMPercentage=85"` |  |
| infraOperator.prometheusConfiguration.enabled | bool | `true` |  |
| infraOperator.resources.limits.cpu | string | `"800m"` |  |
| infraOperator.resources.limits.memory | string | `"1200Mi"` |  |
| infraOperator.resources.requests.cpu | string | `"200m"` |  |
| infraOperator.resources.requests.memory | string | `"500Mi"` |  |
| infraRepo.image.repository | string | `"ghcr.io/th2-net/infra-repo"` |  |
| infraRepo.image.tag | string | `"0.7.0"` |  |
| ingress.host | string | `""` |  |
| productRegistry.name | string | `""` |  |
| productRegistry.password | string | `""` |  |
| productRegistry.secret | string | `"th2-core"` |  |
| productRegistry.username | string | `""` |  |
| proprietaryRegistry.password | string | `""` |  |
| proprietaryRegistry.registry | string | `""` |  |
| proprietaryRegistry.secret | string | `"th2-proprietary"` |  |
| proprietaryRegistry.username | string | `""` |  |
| rabbitmq.fullnameOverride | string | `"rabbitmq"` |  |
| rabbitmq.internal | bool | `true` |  |
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
| rabbitmq.prometheus.operator.enabled | bool | `true` |  |
| rabbitmq.prometheus.operator.serviceMonitor.selector.release | string | `"prometheus"` |  |
| rabbitmq.rabbitmqErlangCookie | string | `""` |  |
| rabbitmq.rabbitmqExchange | string | `"th2-exchange"` |  |
| rabbitmq.rabbitmqMemoryHighWatermark | string | `"1024MB"` |  |
| rabbitmq.rabbitmqPassword | string | `""` |  |
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
Autogenerated from chart metadata using [helm-docs v1.5.0](https://github.com/norwoodj/helm-docs/releases/v1.5.0)
