# Process Diagram

Cycle of a particular release branch
```mermaid
flowchart TB
    branch(release-v*.*.* branch) --> pr(Feature / fix pull request / commit to PR)
    pr --> tests
    tests{Automated tests} --> |Passed| review
    tests{Automated tests} --> |Not passed| pr
    review{Review} --> |Passed| merge(Merge)
    review{Review} --> |Not passed| pr
    merge --> buildDocs(Auto helm chart docs PR)
    merge --> buildInfraRepo(Auto infra-repo build PR)
    buildInfraRepo(Auto infra-repo build PR) --> mergeInfraRepo(Merge)
    mergeInfraRepo --> internal{Internal testing<br/>of a build}
    internal --> |Passed| buildDocs
    internal --> |Not passed| pr
    buildDocs --> mergeDocs(Merge)
    mergeDocs --> release(Release)
    release --> publish(Publish)
    
    linkStyle 2 stroke:green
    linkStyle 4 stroke:green
    linkStyle 10 stroke:green
    
    linkStyle 3 stroke:red
    linkStyle 5 stroke:red
    linkStyle 11 stroke:red

```

