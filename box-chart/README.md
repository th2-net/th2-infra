# box-chart

![Version: 1.0.0](https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square)

th2 Box template chart for infra-operator

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| commonAnnotations | object | `{}` |  |
| component.cradleManager.checksum | string | `"cradleMangerchecksum"` |  |
| component.cradleManager.config | string | `nil` |  |
| component.custom | string | `nil` |  |
| component.dictionaries | string | `nil` |  |
| component.extendedSettings.service.enabled | bool | `false` |  |
| component.grpc | string | `nil` |  |
| component.grpcRouter.checksum | string | `"grpcchecksum"` |  |
| component.grpcRouter.config | string | `nil` |  |
| component.image | string | `nil` |  |
| component.ingress.annotations.default | object | `{}` |  |
| component.ingress.annotations.extra | object | `{}` |  |
| component.ingress.host | string | `nil` |  |
| component.logging.checksum | string | `"loggingchecksum"` |  |
| component.logging.config | string | `nil` |  |
| component.mq | string | `nil` |  |
| component.mqRouter.checksum | string | `"mqchecksum"` |  |
| component.mqRouter.config | string | `nil` |  |
| component.name | string | `"comp1"` |  |
| component.prometheus.enabled | bool | `false` |  |
| component.secrets.cassandra | string | `"cassandra"` |  |
| component.secrets.rabbitMQ | string | `"rabbitMQ"` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.7.0](https://github.com/norwoodj/helm-docs/releases/v1.7.0)