# Process Diagram

From top to bottom
```mermaid
flowchart TB
    branch(release-v*.*.*<br/>branch) --> pr(Pull Request)
    pr --> tests
    tests{Automated tests} --> |Passed| review
    tests{Automated tests} --> |Not passed| commit

    review{Review} --> |Passed| merge(Merge)
    review{Review} --> |Not passed| commit

    merge --> buildDocs(Auto helm chart docs PR)
    buildDocs --> mergeDocs(Merge)
    mergeDocs --> release
    merge --> buildInfraRepo(Auto infra-repo build PR)
    buildInfraRepo(Auto infra-repo build PR) --> mergeInfraRepo(Merge)
    mergeInfraRepo --> internal{Internal testing<br/>of a build}
    commit --> pr
    internal --> |Passed| buildDocs
    internal --> |Not passed| commit(Commit to PR)
    release --> publish(Publish)
    linkStyle 2 stroke:green
    linkStyle 4 stroke:green
    linkStyle 13 stroke:green
    linkStyle 3 stroke:red
    linkStyle 5 stroke:red
    linkStyle 14 stroke:red

```
