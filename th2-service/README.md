# th2

![Version: 1.3.1](https://img.shields.io/badge/Version-1.3.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

th2 service Helm chart

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.bitnami.com/bitnami | cassandra(cassandra) | 5.6.7 |
| https://charts.helm.sh/stable | rabbitmq(rabbitmq-ha) | 1.44.4 |

## Values

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
| infraEditor.image.tag | string | `"1.0.56"` |  |
| infraMgr.cassandra.keyspacePrefix | string | `"schema_"` |  |
| infraMgr.git.httpAuthPassword | string | `""` |  |
| infraMgr.git.httpAuthUsername | string | `""` |  |
| infraMgr.git.privateKeyFileSecret | string | `"infra-mgr"` |  |
| infraMgr.git.repository | string | `"git@github.com:th2-net/th2-demo-configuration.git"` |  |
| infraMgr.git.repositoryLocalCache | string | `"/home/service/repository"` |  |
| infraMgr.git.secretMountPath | string | `"/home/service/keys"` |  |
| infraMgr.git.secretName | string | `"infra-mgr"` |  |
| infraMgr.image.repository | string | `"ghcr.io/th2-net/th2-infra-mgr"` |  |
| infraMgr.image.tag | string | `"0.12.2"` |  |
| infraMgr.kubernetes.configMaps.cassandra | string | `"cradle"` |  |
| infraMgr.kubernetes.configMaps.cassandra-ext | string | `"cradle-external"` |  |
| infraMgr.kubernetes.configMaps.logging | string | `"java-logging-config"` |  |
| infraMgr.kubernetes.configMaps.prometheus | string | `"prometheus-app-config"` |  |
| infraMgr.kubernetes.configMaps.rabbitmq | string | `"rabbit-mq-app-config"` |  |
| infraMgr.kubernetes.configMaps.rabbitmq-ext | string | `"rabbit-mq-external-app-config"` |  |
| infraMgr.kubernetes.ingress | string | `"ingress-rules"` |  |
| infraMgr.kubernetes.namespacePrefix | string | `"th2-"` | not more than 5 symbols |
| infraMgr.kubernetes.secrets[0] | string | `"th2-core"` |  |
| infraMgr.kubernetes.secrets[1] | string | `"th2-solution"` |  |
| infraMgr.kubernetes.secrets[2] | string | `"th2-proprietary"` |  |
| infraMgr.kubernetes.secrets[3] | string | `"cassandra"` |  |
| infraMgr.rabbitmq.passwordLength | int | `24` |  |
| infraMgr.rabbitmq.secret | string | `"rabbitmq"` |  |
| infraMgr.rabbitmq.usernamePrefix | string | `"schema-user-"` |  |
| infraMgr.rabbitmq.vHostPrefix | string | `"schema-"` |  |
| infraOperator.config.chart.name | string | `"infra-operator-tpl"` |  |
| infraOperator.config.chart.repository | string | `"https://th2-net.github.io"` |  |
| infraOperator.config.chart.version | string | `"0.2.0"` |  |
| infraOperator.config.k8sUrl | string | `"<kubernetes-external-entrypoint>"` |  |
| infraOperator.config.namespacePrefixes[0] | string | `"schema-"` |  |
| infraOperator.config.rabbitMQManagement.host | string | `"rabbitmq-discovery.service.svc.cluster.local"` |  |
| infraOperator.config.rabbitMQManagement.password | string | `"${RABBITMQ_PASS}"` |  |
| infraOperator.config.rabbitMQManagement.persistence | bool | `true` |  |
| infraOperator.config.rabbitMQManagement.port | string | `"15672"` |  |
| infraOperator.config.rabbitMQManagement.schemaPermissions.configure | string | `""` |  |
| infraOperator.config.rabbitMQManagement.schemaPermissions.read | string | `".*"` |  |
| infraOperator.config.rabbitMQManagement.schemaPermissions.write | string | `".*"` |  |
| infraOperator.config.rabbitMQManagement.username | string | `"th2"` |  |
| infraOperator.config.schemaSecrets.cassandra | string | `"cassandra"` |  |
| infraOperator.config.schemaSecrets.rabbitMQ | string | `"rabbitmq"` |  |
| infraOperator.image.repository | string | `"ghcr.io/th2-net/th2-infra-operator"` |  |
| infraOperator.image.tag | string | `"2.6.7"` |  |
| infraOperator.persistence | bool | `true` |  |
| ingress.host | string | `nil` |  |
| productRegistry.name | string | `nil` |  |
| productRegistry.password | string | `nil` |  |
| productRegistry.secret | string | `"th2-core"` |  |
| productRegistry.username | string | `nil` |  |
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
| solutionRegistry.password | string | `nil` |  |
| solutionRegistry.registry | string | `nil` |  |
| solutionRegistry.secret | string | `"th2-solution"` |  |
| solutionRegistry.username | string | `nil` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.5.0](https://github.com/norwoodj/helm-docs/releases/v1.5.0)
