# Process Diagram

From top to bottom
```mermaid
flowchart TB
    branch(release-v*.*.*<br/>branch) --> change(Pull Request<br/>Commit to PR)
    change --> tests
    tests{Automated tests} --> |Passed| review
    tests{Automated tests} --> |Not passed| change
    review{Review} --> |Not passed| change
    review{Review} --> |Passed| merge(Merge)
    merge --> buildDocs(Auto helm chart docs PR)
    buildDocs --> mergeDocs(Merge)
    mergeDocs --> release
    merge --> buildInfraRepo(Auto infra-repo build PR)
    buildInfraRepo(Auto infra-repo build PR) --> mergeInfraRepo(Merge)
    mergeInfraRepo --> internal{Internal testing<br/>of a build}
    internal --> |Not passed| change
    internal --> |Passed| buildDocs
    release --> publish(Publish)
```

From left to right
```mermaid
flowchart LR
    branch(release-v*.*.*<br/>branch) --> change(Pull Request<br/>Commit to PR)
    change --> tests
    tests{Automated tests} --> |Passed| review
    tests{Automated tests} --> |Not passed| change
    review{Review} --> |Not passed| change
    review{Review} --> |Passed| merge(Merge)
    merge --> buildDocs(Auto helm chart docs PR)
    buildDocs --> mergeDocs(Merge)
    mergeDocs --> release
    merge --> buildInfraRepo(Auto infra-repo build PR)
    buildInfraRepo(Auto infra-repo build PR) --> mergeInfraRepo(Merge)
    mergeInfraRepo --> internal{Internal testing<br/>of a build}
    internal --> |Not passed| change
    internal --> |Passed| buildDocs
    release --> publish(Publish)
```
