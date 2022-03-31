# Migrations

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