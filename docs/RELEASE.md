# Process Diagram

Cycle of a particular release branch
```mermaid
flowchart TB
    init(master or release-v*.*.x branch, release-v1.2.x) --> branch
    init --> milestone(Create milestone, v1.2.3)
    branch(Create release branch, release-v1.2.3) --> pr(Feature / fix pull request / commit to PR)
    milestone --> pr
    pr --> tests
    tests{Automated tests} --> |Passed| review
    tests{Automated tests} --> |Not passed| pr
    review{Review} --> |Passed| merge(Merge into release branch)
    review{Review} --> |Not passed| pr
    merge --> buildDocs(Auto helm chart docs PR)
    merge --> buildInfraRepo(Auto infra-repo build PR)
    buildInfraRepo(Auto infra-repo build PR) --> mergeInfraRepo(Merge into release branch)
    mergeInfraRepo --> internal{Internal testing<br/>of a build}
    internal --> |Passed| buildDocs
    internal --> |Not passed| pr
    buildDocs --> mergeDocs(Merge into release branch)
    mergeDocs --> tag(Git tag as v*.*.*, v1.2.3)
    tag --> publish(Publish)
    publish --> mergeToMaster(Merge into release-v*.*.x or master branch)
    
    linkStyle 5 stroke:green
    linkStyle 7 stroke:green
    linkStyle 13 stroke:green
    
    linkStyle 6 stroke:red
    linkStyle 8 stroke:red
    linkStyle 14 stroke:red

```

