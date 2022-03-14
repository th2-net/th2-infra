# th2-infra end-to-end test
This project leverages Kind (Kubernetes in Docker), Ansible and Go lang for end-to-end tests execution.

Coverage:
* deployment th2-infra components
* deployment schema components from infra-git via ssh
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
    ├── deploy                               - Ansible templates and files
    ├── e2e-test-schema                      - th2 schema files (boxes, core, links, dictionaries and etc.)
    ├── e2e-via-ssh-deployment-playbook.yaml - Ansible playbook to deploy th2 with infra-mgr gitops via ssh
    ├── th2_infra_test.go                    - tests written in Go with Terratest module
    └── kind-cluster.yaml                    - kind cluster config
```

## Settings and variables
There are no special settings and variables except default. Test schema is deployed from e2e-test-schema directory into th2-schema namespace

Namespaces for th2-infra can be set up in the following workflow variables:
```yaml
env:
  INFRA_NAMESPACE: "${{ github.event.inputs.infraNamespace || 'service' }}"
  MONITORING_NAMESPACE: monitoring
  SCHEMA_NAMESPACE: "${{ github.event.inputs.schemaNamespace || 'th2-v150' }}"
```

If it is required to test infra operator template chart from git, then it can be set in e2e-via-ssh-deployment-playbook.yaml > "Deploy th2-infra":
```yaml
          # infraOperator:
          #  config:
          #    chart:
          #      git: https://github.com/th2-net/infra-operator-tpl.git
          #      ref: v0.2.0
          #      path: ./
```