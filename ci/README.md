# th2-infra end-to-end test
This project leverages Kind (Kubernetes in Docker), Ansible and Go lang for end-to-end tests execution.

Coverage:
* deployment th2-infra components
* deployment schema components from git via ssh
* http endpoints from service namespace
* http endpoints from schema namespace

Output:
* workflow logs
* artifact logs attached to the job

## Structure
```
├── .github
|   └── e2e-test.yml                         - e2e testing Github Action
└── ci
    ├── cassandra
    |   └── initdb.yaml                      - ConfigMap with Cassandra init cql script
    ├── e2e-via-ssh-deployment-playbook.yaml - ansible playbook to deploy th2 with infra-mgr gitops via ssh
    ├── th2_infra_test.go                    - tests written in Go with Terratest module
    └── kind-cluster.yaml                    - kind cluster config
```

## Settings and variables
Repository with schema contains version branch from which th2 namespace will be deployed, e.g. v130 https://github.com/th2-net/e2e-test-schema/tree/v130.

Namespaces for th2-infra can be set up in the following workflow variables:
```yaml
env:
  INFRA_NAMESPACE: "${{ github.event.inputs.infraNamespace || 'service' }}"
  MONITORING_NAMESPACE: monitoring
  SCHEMA_NAMESPACE: "${{ github.event.inputs.schemaNamespace || 'th2-v150' }}"
```
**Note**: _CRs with any API changes must be allocated in a new versioned branch!_

Repository with schema is set as a value in e2e-via-ssh-deployment-playbook.yaml > "Deploy th2-infra-base" step
```
    infraMgr:
      git:
        repository: git@github.com:th2-net/e2e-test-schema.git
```

If it is required to test infra operator template chart from git, then it can be set in e2e-via-ssh-deployment-playbook.yaml > "Deploy th2-infra-base":
```yaml
          # infraOperator:
          #   chart:
          #     git: https://github.com/th2-net/infra-operator-tpl.git
          #     ref: v0.2.0
          #     path: ./
```

Private key for infra-mgr is set as a repository secret and passed as E2E_PRIVATE_KEY_NEW env variable to playbook. Public key is set as e2e-test deploy key in the schema repository.
