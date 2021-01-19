# th2-infra end-to-end test

This project leverages Kind (Kubernetes in Docker), Ansible and Go lang for end-to-end tests execution.

Coverage:
* deployment th2-infra components
* deployment schema components from git via ssh
* http endpoints from service namespace
* http endpoint from schema namespace

## Structure

```
├── .github
|   └── e2e-test.yml                         - e2e testing Github Action
└── ci
    ├── e2e-via-ssh-deployment-playbook.yaml - ansible playbook to deploy th2 with infra-mgr gitops via ssh
    ├── th2_infra_test.go                    - tests written in Go with Terratest module
    └── kind-cluster.yaml                    - kind cluster config
```

## Settings and variables
Repository with schema contains version branch from which th2 namespace will be deployed, e.g. e2e-v101 https://github.com/th2-net/th2-infra-schema-demo/tree/e2e-v101.

Namespace with schema is set as a constant in go tests accordingly.

```
schemaNamespace  = "schema-e2e-v101"
```
**Note**: _CRs with any API changes must be allocated in a new versioned branch!_

Repository with schema is set in values in e2e-via-ssh-deployment-playbook.yaml > "Deploy th2-infra-base" step
```
    infraMgr:
      git:
        repository: git@github.com:th2-net/th2-infra-schema-demo.git
```
Private key is set as Secret in the current repository and passed as E2E_PRIVATE_KEY env variable to playbook. Public key is set as e2e-test deploy key in the schema repository.