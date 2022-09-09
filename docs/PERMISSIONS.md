
 ### infra-mgr     
| Resources                   | Non-Resource URLs     | Resource Names        |  Verbs                |
| ---------------------       | --------------------- | --------------------- | --------------------|
| configmaps                  | []                    | []                    | [*]                   |
| namespaces                  | []                    | []                    | [*]                   |
| pods                        | []                    | []                    | [*]                   |
| secrets                     | []                    | []                    | [*]                   |
| services                    | []                    | []                    | [*]                   |
| deployments.apps            | []                    | []                    | [*]                   |
| servicemonitors.monitoring.coreos.com | []                    | []                    | [*]                   |
| ingresses.networking.k8s.io | []                    | []                    | [*]                   |
| *.th2.exactpro.com          | []                    | []                    | [*]                   |



### infra-operator
| Resources                                      | Non-Resource URLs | Resource Names  | Verbs |
| ---------------------------------------------- | ----------------- | --------------  | ----- |
| configmaps                                     | []                | []              | [*]   |
| namespaces                                     | []                | []              | [*]   |
| secrets                                        | []                | []              | [*]   |
| customresourcedefinitions.apiextensions.k8s.io | []                | []              | [*]   |
| helmreleases.helm.fluxcd.io/finalizers         | []                | []              | [*]   |
| helmreleases.helm.fluxcd.io/status             | []                | []              | [*]   |
| helmreleases.helm.fluxcd.io                    | []                | []              | [*]   |
| *.th2.exactpro.com                             | []                | []              | [*]   |

### th2-kubernetes-dashboard 
| Resources                                      | Non-Resource URLs | Resource Names | Verbs                |
| --------------------------------------         | ----------------- | -------------- | -------------------- |
| bindings                                       | []                | []             | [get list watch]     |
| configmaps                                     | []                | []             | [get list watch]     |
| endpoints                                      | []                | []             | [get list watch]     |
| events                                         | []                | []             | [get list watch]     |
| limitranges                                    | []                | []             | [get list watch]     |
| namespaces/status                              | []                | []             | [get list watch]     |
| namespaces                                     | []                | []             | [get list watch]     |
| nodes                                          | []                | []             | [get list watch]     |
| persistentvolumeclaims                         | []                | []             | [get list watch]     |
| persistentvolumes                              | []                | []             | [get list watch]     |
| pods/log                                       | []                | []             | [get list watch]     |
| pods/status                                    | []                | []             | [get list watch]     |
| pods                                           | []                | []             | [get list watch]     |
| replicationcontrollers/scale                   | []                | []             | [get list watch]     |
| replicationcontrollers/status                  | []                | []             | [get list watch]     |
| replicationcontrollers                         | []                | []             | [get list watch]     |
| resourcequotas/status                          | []                | []             | [get list watch]     |
| resourcequotas                                 | []                | []             | [get list watch]     |
| serviceaccounts                                | []                | []             | [get list watch]     |
| services                                       | []                | []             | [get list watch]     |
| daemonsets.apps                                | []                | []             | [get list watch]     |
| deployments.apps/scale                         | []                | []             | [get list watch]     |
| deployments.apps                               | []                | []             | [get list watch]     |
| replicasets.apps/scale                         | []                | []             | [get list watch]     |
| replicasets.apps                               | []                | []             | [get list watch]     |
| statefulsets.apps                              | []                | []             | [get list watch]     |
| horizontalpodautoscalers.autoscaling           | []                | []             | [get list watch]     |
| cronjobs.batch                                 | []                | []             | [get list watch]     |
| jobs.batch                                     | []                | []             | [get list watch]     |
| daemonsets.extensions                          | []                | []             | [get list watch]     |
| deployments.extensions/scale                   | []                | []             | [get list watch]     |
| deployments.extensions                         | []                | []             | [get list watch]     |
| ingresses.extensions                           | []                | []             | [get list watch]     |
| networkpolicies.extensions                     | []                | []             | [get list watch]     |
| replicasets.extensions/scale                   | []                | []             | [get list watch]     |
| replicasets.extensions                         | []                | []             | [get list watch]     |
| replicationcontrollers.extensions/scale        | []                | []             | [get list watch]     |
| networkpolicies.networking.k8s.io              | []                | []             | [get list watch]     |
| poddisruptionbudgets.policy                    | []                | []             | [get list watch]     |
| clusterrolebindings.rbac.authorization.k8s.io  | []                | []             | [get list watch]     |
| clusterroles.rbac.authorization.k8s.io         | []                | []             | [get list watch]     |
| rolebindings.rbac.authorization.k8s.io         | []                | []             | [get list watch]     |
| roles.rbac.authorization.k8s.io                | []                | []             | [get list watch]     |
| storageclasses.storage.k8s.io                  | []                | []             | [get list watch]     |
| volumeattachments.storage.k8s.io               | []                | []             | [get list watch]     |
| customresourcedefinitions.apiextensions.k8s.io | []                | []             | [get watch list]     |
| *.helm.fluxcd.io                               | []                | []             | [get watch list]     |
| *.th2.exactpro.com                             | []                | []             | [get watch list]     |

### th2-dashboard-metrics
| Resources                      | Non-Resource URLs | Resource Names  | Verbs            |
| ------------------------------ | ----------------- | --------------  | -----            |
| nodes.metrics.k8s.io           | []                | []              | [get list watch] |
| pods.metrics.k8s.io            | []                | []              | [get list watch] |


### helm-operator
| Resources     | Non-Resource URLs | Resource Names  | Verbs |
| --------------|-------------------|-----------------|------ |
| [*]           | []                | []              | [*]   |
| [*]           | [*]               | []              | [*]   |

### ingress nginx
| Resources                          | Non-Resource URLs | Resource Names |Verbs                    |
| -----------------------------------| ----------------- | -------------- |-------------------------|
| events                             | []                | []             | [create patch]          |
| services                           | []                | []             | [get list update watch] |
| ingresses.extensions               | []                | []             | [get list watch]        |
| ingressclasses.networking.k8s.io   | []                | []             | [get list watch]        |
| ingresses.networking.k8s.io        | []                | []             | [get list watch]        |
| nodes                              | []                | []             | [list watch get]        |
| configmaps                         | []                | []             | [list watch]            |
| endpoints                          | []                | []             | [list watch]            |
| pods                               | []                | []             | [list watch]            |
| secrets                            | []                | []             | [list watch]            |
| ingresses.extensions/status        | []                | []             | [update]                |
| ingresses.networking.k8s.io/status | []                | []             | [update]                |

### loki-promtail
| Resources     | Non-Resource URLs | Resource Names  | Verbs            |
| ---------     | ----------------- | --------------  | -----            |
| endpoints     | []                | []              | [get watch list] |
| nodes/proxy   | []                | []              | [get watch list] |
| nodes         | []                | []              | [get watch list] |
| pods          | []                | []              | [get watch list] |
| services      | []                | []              | [get watch list] |

### prometheus-operator
| Resources                                            | Non-Resource URls    | Resource Names  | Verbs                       |
| ---------------------------------------------------- | -------------------- | --------------  | --------------------------- |
| configmaps                                           | []                   | []              | [*]                         |
| secrets                                              | []                   | []              | [*]                         |
| statefulsets.apps                                    | []                   | []              | [*]                         |
| alertmanagerconfigs.monitoring.coreos.com            | []                   | []              | [*]                         |
| alertmanagers.monitoring.coreos.com/finalizers       | []                   | []              | [*]                         |
| alertmanagers.monitoring.coreos.com                  | []                   | []              | [*]                         |
| podmonitors.monitoring.coreos.com                    | []                   | []              | [*]                         |
| probes.monitoring.coreos.com                         | []                   | []              | [*]                         |
| prometheuses.monitoring.coreos.com/finalizers        | []                   | []              | [*]                         |
| prometheuses.monitoring.coreos.com                   | []                   | []              | [*]                         |
| prometheusrules.monitoring.coreos.com                | []                   | []              | [*]                         |
| servicemonitors.monitoring.coreos.com                | []                   | []              | [*]                         |
| thanosrulers.monitoring.coreos.com/finalizers        | []                   | []              | [*]                         |
| thanosrulers.monitoring.coreos.com                   | []                   | []              | [*]                         |
| endpoints                                            | []                   | []              | [get create update delete]  |
| services/finalizers                                  | []                   | []              | [get create update delete]  |
| services                                             | []                   | []              | [get create update delete]  |
| namespaces                                           | []                   | []              | [get list watch]            |
| ingresses.networking.k8s.io                          | []                   | []              | [get list watch]            |
| pods                                                 | []                   | []              | [list delete]               |
| nodes                                                | []                   | []              | [list watch]                |

