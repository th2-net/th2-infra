th2
===
th2 service Helm chart

Current chart version is `1.0.0`



## Chart Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.bitnami.com/bitnami | cassandra | 5.4.1 |
| https://kubernetes-charts.storage.googleapis.com/ | rabbitmq-ha | 1.44.4 |

## Chart Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| cassandra.cluster.datacenter | string | `"datacenter1"` |  |
| cassandra.dbUser.password | string | `""` |  |
| cassandra.dbUser.user | string | `"cassandra"` |  |
| cassandra.fullnameOverride | string | `"cassandra"` |  |
| cassandra.internal | bool | `true` |  |
| cassandra.keyspace | string | `"cradle"` |  |
| cassandra.persistence.enabled | bool | `false` |  |
| cassandra.persistence.size | string | `"50Gi"` |  |
| cassandra.persistence.storageClass | string | `"local-storage"` |  |
| infraEditor.image.repository | string | `"ghcr.io/th2-net/th2-infra-editor"` |  |
| infraEditor.image.tag | string | `"1.0.29"` |  |
| infraMgr.cassandra.keyspacePrefix | string | `"schema_"` |  |
| infraMgr.git.privateKeyFileSecret | string | `"infra-mgr"` |  |
| infraMgr.git.repository | string | `"git@github.com:th2-net/th2-demo-configuration.git"` |  |
| infraMgr.git.repositoryLocalCache | string | `"/home/service/repository"` |  |
| infraMgr.git.secretMountPath | string | `"/home/service/keys"` |  |
| infraMgr.git.secretName | string | `"infra-mgr"` |  |
| infraMgr.image.repository | string | `"ghcr.io/th2-net/th2-infra-mgr"` |  |
| infraMgr.image.tag | string | `"0.8.5"` |  |
| infraMgr.kubernetes.configMaps.cassandra | string | `"cradle"` |  |
| infraMgr.kubernetes.configMaps.logging | string | `"java-logging-config"` |  |
| infraMgr.kubernetes.configMaps.rabbitmq | string | `"rabbit-mq-app-config"` |  |
| infraMgr.kubernetes.configMaps.rabbitmqManagement | string | `"rabbitmq-mng-params"` |  |
| infraMgr.kubernetes.ingress | string | `"ingress-rules"` |  |
| infraMgr.kubernetes.namespacePrefix | string | `"schema-"` |  |
| infraMgr.kubernetes.secrets[0] | string | `"chart-secrets"` |  |
| infraMgr.kubernetes.secrets[1] | string | `"git-chart-creds"` |  |
| infraMgr.kubernetes.secrets[2] | string | `"th2-core"` |  |
| infraMgr.kubernetes.secrets[3] | string | `"th2-solution"` |  |
| infraMgr.kubernetes.secrets[4] | string | `"cassandra"` |  |
| infraMgr.rabbitmq.passwordLength | int | `24` |  |
| infraMgr.rabbitmq.secret | string | `"rabbitmq"` |  |
| infraMgr.rabbitmq.usernamePrefix | string | `"schema-user-"` |  |
| infraMgr.rabbitmq.vHostPrefix | string | `"schema-"` |  |
| infraOperator.image.repository | string | `"ghcr.io/th2-net/th2-infra-operator"` |  |
| infraOperator.image.tag | string | `"2.0.14"` |  |
| ingress.host | string | `nil` |  |
| productRegistry.name | string | `nil` |  |
| productRegistry.password | string | `nil` |  |
| productRegistry.secret | string | `"th2-core"` |  |
| productRegistry.username | string | `nil` |  |
| rabbitmq.fullnameOverride | string | `"rabbitmq"` |  |
| rabbitmq.internal | bool | `true` |  |
| rabbitmq.managementPassword | string | `""` |  |
| rabbitmq.managementUsername | string | `"th2-mng"` |  |
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
| rabbitmq.replicaCount | int | `1` |  |
| rabbitmq.service.amqpNodePort | int | `32000` |  |
| rabbitmq.service.type | string | `"NodePort"` |  |
| solutionRegistry.password | string | `nil` |  |
| solutionRegistry.registry | string | `nil` |  |
| solutionRegistry.secret | string | `"th2-solution"` |  |
| solutionRegistry.username | string | `nil` |  |
