# Migrations

## Migration to RELEASE v2.1.0
* jupyterhub UI is now accessable through same port as other components
  <details>
    <summary>Enabling jhub ingress</summary>

    ### Enabling host and ingressClassName
    * host and ingressClassName should be added in values
    ```
    jupyterhub:
      ingress:
        ingressClassName: nginx
        hosts:
        - <hostname>
    ```
  </details>

* cassandra dashboard has been corrected by th2 team
  <details>
    <summary>Replacing dashboard</summary>
    
    ### changing the name of the dashboard
    * cassandra dashboard name should be updated in grafana values
    ```
    grafana:
      dashboards:
        default:
          Cassandra-dashboard:
            url: http://infra-repo.service.svc.cluster.local:8080/dashboards/Cassandra-Dashboard(corrected-by-th2-team).json
    ```
  </details>

## Migration to RELEASE v2.0.0

* jupyterhub is now included as a dependency and should not be deployed separately. All its values are under jhub parent value. 
* Update secrets.yaml
  <details>
    <summary>Adding Jupyterhub users</summary>

    ### Adding Credentials
    * Values for admin and regular user credentials should be added into secrets.yaml
    ```
    jupyterhub:
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
  </details>

* Migrated to new Kubernetes API versions. Now th2-infra supports Kubernetes 1.22-1.26 releases
* NGINX Ingress Controller chart must be upgraded 4.1.2 > 4.3.0.
* Prometheus-stack should be upgraded 21.0.5 > 41.4.0.
* Loki-stack should be upgraded 2.6.5 > 2.8.3.

* Infra, Diagnostic and JVM dashboards should be added in grafana during deployment
  <details>
    <summary>Adding dashboards</summary>

    ### Adding Urls in Prometheus-stack
    * Infra, Diagnostic and JVM dashboards should be added in grafana, monitoring-old should be removed and monitoring-new should be renamed during deployment
    ```
      grafana:
        dashboards:
          default:
            Infra-dashboard:
              url: http://infra-repo.service.svc.cluster.local:8080/dashboards/infra-dashboard.json
            JVM-dashboard:
              url: http://infra-repo.service.svc.cluster.local:8080/dashboards/JVM-dashboard(for-Prometheus-Operator).json
            Diagnostic-dashboard:
              url: http://infra-repo.service.svc.cluster.local:8080/dashboards/diagnostic-dashboard-v1.0.2.json
    ```
  </details>
  <details>
    <summary>Editing monitoring dashboards</summary>

    ### Edit Urls in Prometheus-stack
    * The section with monitoring-old and monitoring-new dashboards should be replaced with the following.
    ```
      grafana:
        dashboards:
          default:
            Monitoring:
              url: http://infra-repo.service.svc.cluster.local:8080/dashboards/Monitoring.json
    ```
  </details>
* converter has been added to infra
  <details>
    <summary>Include converter in installation</summary>

    ### Adding converter values
    * schema repo should be passed to converter.git.repository
    ```
      converter:
        git:
          repository: <repository>
    ```

    ### Adding converter key in secrets
    * secrets.yaml should convain value from converter-ed25519.key
    ```
      converter:
        git:
          privateKey: <privateKey>
    ```
  </details>

## Migration to RELEASE v1.8.0
* Migrated to new Kubernetes API versions. Now th2-infra supports Kubernetes 1.19-1.23 releases
* Prometheus stack must be upgraded 15.0.0 > 21.0.5
* Loki-stack must be upgraded 2.4.1 > 2.6.5. Remove a release before upgrade. Set new values for `loki` with ones from `../example-values/loki.values.yaml`.
* NGINX Ingress Controller chart must be upgraded 3.31.0 > 4.1.2.
```
$ helm install -n service --version=4.1.2 ingress ingress-nginx/ingress-nginx -f ./ingress.values.yaml
```
* HelmOperator dependency upgraded 1.2.0 > 1.4.2. HelmRelease CRD must be removed before infra installation
```
$ kubectl delete customresourcedefinitions helmreleases.helm.fluxcd.io
```
* infra-mgr secrets are now created automatically.
  <details>
    <summary>Infra-mgr secrets should be removed from service namespace</summary>

    ### Delete secret infra-mgr
    * since infra-mgr secret is automatically created old secret should be deleted (if present)
    ```
    $ kubectl -n service delete secret infra-mgr
    ```

    ### Delete secret th2-git-access-schemas
    * since th2-git-access-schemas secret is automatically created old secret should be deleted (if present)
    ```
    $ kubectl -n service delete secret th2-git-access-schemas
    ```

    ### secrets.yaml file should be updated
    * secrets.yaml should contain value from infra-mgr-rsa.key
    ```
      infraMgr:
        git:
          privateKey: <privateKey>
    ```
  </details>
* Dashboards, Dashboard Provider and grafana plugins should be added in grafana during deployment.
  <details>
    <summary>Adding Dashboard Provider, Dashboards and plugins in Prometheus-stack</summary>

    ### Adding Dashboard Provider in Prometheus-stack
    * Dashboard Provider should be added in grafana by dashboardProviders.dashboardproviders.yaml.
    ```
      grafana:
        dashboardProviders:
          dashboardproviders.yaml:
            apiVersion: 1
            providers:
              - name: 'default'
                orgId: 1
                folder: ''
                type: file
                disableDeletion: false
                editable: true
                options:
                  path: /var/lib/grafana/dashboards/default
    ```
    ### Adding Dashboard Urls in Prometheus-stack
    * Dashboards should be added in grafana from infra-repo by Url.
    ```
      grafana:
        dashboards:
          default:
            Cassandra-dashboard:
              url: http://infra-repo.service.svc.cluster.local:8080/dashboards/cassandra-dashboard_rev2.json
            Rabbitmq-dashboard:
              url: http://infra-repo.service.svc.cluster.local:8080/dashboards/rabbitmq-overview_rev11.json
            Node-monitoring:
              url: http://infra-repo.service.svc.cluster.local:8080/dashboards/nodes-monitoring-v1.0.0.json
            Namespace-health:
              url: http://infra-repo.service.svc.cluster.local:8080/dashboards/namespace_health-v1.0.2.json
            Components-logs:
              url: http://infra-repo.service.svc.cluster.local:8080/dashboards/components-logs.json
            Monitoring-1:
              url: http://infra-repo.service.svc.cluster.local:8080/dashboards/Monitoring-old.json
            Monitoring-3:
              url: http://infra-repo.service.svc.cluster.local:8080/dashboards/Monitoring-new.json
    ``` 
    ### Adding Plugin Urls in Prometheus-stack
    * Plugins should be added in grafana from infra-repo by plugins.
    ```
      grafana:
        plugins:
          - http://infra-repo.service.svc.cluster.local:8080/plugins/flant-statusmap-panel-0.4.2.zip;flant-statusmap-panel
          - http://infra-repo.service.svc.cluster.local:8080/plugins/vonage-status-panel-1.0.11.zip;vonage-status-panel
          - http://infra-repo.service.svc.cluster.local:8080/plugins/blackmirror1-statusbygroup-panel-1.1.2.zip;blackmirror1-statusbygroup-panel
          - http://infra-repo.service.svc.cluster.local:8080/plugins/grafana-polystat-panel-1.2.8.any.zip;grafana-polystat-panel
          - http://infra-repo.service.svc.cluster.local:8080/plugins/briangann-gauge-panel-0.0.9.zip;briangann-gauge-panel
          - http://infra-repo.service.svc.cluster.local:8080/plugins/yesoreyeram-boomtable-panel-1.4.1.zip;yesoreyeram-boomtable-panel
    ```
  </details>
* Rabbitmq values should be updated.
  <details>
    <summary>Updating rabbitmq values</summary>

    ### Rabbitmq should be adjusted to bitnami chart
    * Rabbitmq values should be changed in service.values.
    ```
      rabbitmq:
        persistence:
          storageClass: local-storage
        ingress:
          hostname: <hostname>
    ```
    ### Authentication format should be updated
    * rabbitmqUsername and rabbitmqPassword should be replaced in secrets.yaml
    ```
      rabbitmq:
        auth:
          username: th2
          password: rab-pass
          # must be random string
          erlangCookie: cookie
    ```
    _Note: Persistence data from previous rabbitmq should be deleted before instalation_
  </details>
* Prometheus and Alert Manager should be added in Prometheus-stack.
  <details>
    <summary>Update Alert Manager and Prometheus Operator</summary>

    ### Add ingress rules to Alert Manager
    * ingress rules should be added in alertmanager.
    ```
      alertmanager:
        alertmanagerSpec:
          externalUrl: http://localhost:9093/alertmanager
        ingress:
          hosts: []
    ```
    ### Add ingress rules to Prometheus Operator
    * ingress rules should be added in prometheusOperator.
    ```
      prometheus:
        prometheusSpec:
          externalUrl: http://localhost:9090/prometheus
        ingress:
          hosts: []
    ```
  </details>
* Kube-state-metrics values should be updated
  <details>
    <summary>Update kube-state-metrics</summary>

    ### Add values to kube-state-metrics
    * metricLabelsAllowlist should be added in kube-state-metrics
    ```
      kube-state-metrics:
        metricLabelsAllowlist: ['pods=[*]','deployments=[*]']
    ```
  </details>
* InfraGit values have to be be updated.
  <details>
    <summary>new persistence configuration</summary>

    ### Changing persistence structure
    * persistence has to be updated to new format.
    ```
      infraGit:
        internal: true
        nodePort: 32600
        image:
          repository: ghcr.io/th2-net/git-ssh
          tag: v0.1.0
        persistence:
          enabled: true
          # -- "repos-volume" claim will be created and mounted if empty
          existingClaim: ""
    ```
  </details>
* secrets.yaml has to be updated.
  <details>
    <summary>Registries have new structure</summary>

    ### Changing Registries structure
    * registries have to updated to new format.
    ```
      registries:
        registry1.name.com:8080:
          username: <username>
          password: <password>
        registry2.name.com:8080:
          username: <username>
          password: <password>
    ```
  </details>

* ingress.annotations structure changed
  <details>
    <summary>Update custom annotations if required</summary>

    ```
      ingress:
        annotations:
          th2Namespace: {...}
          infraNamespace: {...}
          root: {...}
    ```
  </details>

## Migration to RELEASE v1.7.2
* Loki must be added in grafana during deployment. 
  <details>
    <summary>Adding Loki datasource in Prometheus-stack</summary>

    ### Adding Loki datasource in Prometheus-stack
    * Loki datasource can be added in Prometheus-stack by grafana.additionalDataSources.
    ```
      grafana:
        additionalDataSources:
        - name: Loki
          access: proxy
          type: loki
          url: http://loki:3100
          jsonData:
            maxLines: "5000"
    ``` 
  </details>
* Delete old infra-mgr secret and replace it with new one.
  <details>
    <summary>Add New Secret</summary>

    ### Add New Secret 
    * Delete old infra-mgr secret
    ```
    $ kubectl -n service delete secret infra-mgr
    ```
    * Create infra-mgr secret from the private key:
    ```
    $ kubectl -n service create secret generic infra-mgr --from-file=id_rsa=./infra-mgr-rsa.key
    ```
  </details>

## Migration to RELEASE v1.7.1
* Loki stack 2.4.1 chart version must be used during deployment.
  <details>
    <summary>Upgrade-loki</summary>

    ### Upgrade loki
    * Loki can be upgraded without additional configuration only if new version uses the same schema version. Current config can be retrieved from cluster:
    ```
    kubectl get secret -n monitoring loki -o jsonpath="{.data.loki\.yaml}"|base64 -d; echo
    ``` 
    schema version is defined in `schema_config.schema` parameter.
    Schema of new version loki can be found in chart default values for loki

    * If schema versions are different should be used transition config for loki. Example of this config:
    ```
        schema_config:
          configs:
          - from: "2018-04-15"
            index:
              period: 168h
              prefix: index_
            object_store: filesystem
            schema: v9
            store: boltdb
          - from: "2022-01-22"
            store: boltdb-shipper
            object_store: filesystem
            schema: v11
            index:
              prefix: index_
              period: 24h
        storage_config:
        # because boltdb is used in old schema we need to define this storage
          boltdb:
            directory: /data/loki/index
    ``` 
  </details>
* Helm operator chart now is included as dependency and should not be deployed separately. All its values are under *helmoperator* parent value.
* Kubernetes dashboard chart now is included as dependency and should not be deployed separately. Previous deployment must be uninstalled from "monitoring" namespace. All its values are under *dashboard* parent value.

More information about seamless migration between schemas:
https://grafana.com/docs/loki/v2.2.0/storage/#schema-configs
https://grafana.com/docs/loki/v2.2.0/configuration/#schema_config