# Process Diagram

From top to bottom
```mermaid
flowchart TB
    branch(release-v*.*.*<br/>branch) --> change(Pull Request<br/>Commit to PR)
    change --> tests
    tests{Automated tests} --> |Passed| review
    linkStyle 2 stroke:green
    tests{Automated tests} --> |Not passed| change
    linkStyle 3 stroke:red
    review{Review} --> |Not passed| change
    linkStyle 4 stroke:red
    review{Review} --> |Passed| merge(Merge)
    linkStyle 5 stroke:green
    merge --> buildDocs(Auto helm chart docs PR)
    buildDocs --> mergeDocs(Merge)
    mergeDocs --> release
    merge --> buildInfraRepo(Auto infra-repo build PR)
    buildInfraRepo(Auto infra-repo build PR) --> mergeInfraRepo(Merge)
    mergeInfraRepo --> internal{Internal testing<br/>of a build}
    internal --> |Not passed| change
    linkStyle 12 stroke:red
    internal --> |Passed| buildDocs
    linkStyle 13 stroke:green
    release --> publish(Publish)
```

From left to right
```mermaid
flowchart LR
    branch(release-v*.*.*<br/>branch) --> change(Pull Request<br/>Commit to PR)
    change --> tests
    tests{Automated tests} --> |Passed| review
    linkStyle 2 stroke:green
    tests{Automated tests} --> |Not passed| change
    linkStyle 3 stroke:red
    review{Review} --> |Not passed| change
    linkStyle 4 stroke:red
    review{Review} --> |Passed| merge(Merge)
    linkStyle 5 stroke:green
    merge --> buildDocs(Auto helm chart docs PR)
    buildDocs --> mergeDocs(Merge)
    mergeDocs --> release
    merge --> buildInfraRepo(Auto infra-repo build PR)
    buildInfraRepo(Auto infra-repo build PR) --> mergeInfraRepo(Merge)
    mergeInfraRepo --> internal{Internal testing<br/>of a build}
    internal --> |Not passed| change
    linkStyle 12 stroke:red
    internal --> |Passed| buildDocs
    linkStyle 13 stroke:green
    release --> publish(Publish)
```
