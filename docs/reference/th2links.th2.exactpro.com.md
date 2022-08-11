---
title: Th2Link CRD schema reference (group th2.exactpro.com)
linkTitle: Th2Link
description: |
  Custom resource definition (CRD) schema reference page for the Th2Link resource (th2links.th2.exactpro.com), as part of the Giant Swarm Management API documentation.
weight: 100
crd:
  name_camelcase: Th2Link
  name_plural: th2links
  name_singular: th2link
  group: th2.exactpro.com
  technical_name: th2links.th2.exactpro.com
  scope: Namespaced
  source_repository: https://github.com/th2-net/th2-infra
  source_repository_ref: crd-gerenrator-flow
  versions:
    - v1
  topics:
    - managementcluster
layout: crd
owner:
  - https://github.com/th2-net/th2-infra
aliases:
  - /reference/cp-k8s-api/th2links.th2.exactpro.com/
technical_name: th2links.th2.exactpro.com
source_repository: https://github.com/th2-net/th2-infra
source_repository_ref: crd-gerenrator-flow
---

# Th2Link

<dl class="crd-meta">
<dt class="fullname">Full name:</dt>
<dd class="fullname">th2links.th2.exactpro.com</dd>
<dt class="groupname">Group:</dt>
<dd class="groupname">th2.exactpro.com</dd>
<dt class="singularname">Singular name:</dt>
<dd class="singularname">th2link</dd>
<dt class="pluralname">Plural name:</dt>
<dd class="pluralname">th2links</dd>
<dt class="scope">Scope:</dt>
<dd class="scope">Namespaced</dd>
<dt class="versions">Versions:</dt>
<dd class="versions"><a class="version" href="#v1" title="Show schema for version v1">v1</a></dd>
</dl>



<div class="crd-schema-version">
<h2 id="v1">Version v1</h2>


<h3 id="crd-example-v1">Example CR</h3>

```yaml
apiVersion: th2.exactpro.com/v1
kind: Th2Link
metadata:
  name: dictionary-links
spec:
  dictionaries-relation:
  - name: codec-fix-dictionary
    box: codec-fix
    dictionary:
      name: fix50-aix
      type: MAIN
  - name: codec-fix-neg-dictionary
    box: codec-fix-neg
    dictionary:
      name: fix50-aix-neg
      type: MAIN
  - name: demo-conn1-dictionary
    box: demo-conn1
    dictionary:
      name: fix50-aix
      type: MAIN
  - name: demo-conn2-dictionary
    box: demo-conn2
    dictionary:
      name: fix50-aix
      type: MAIN
  - name: fix-t99rt001-dictionary
    box: fix-t99rt001
    dictionary:
      name: fix50-aix
      type: MAIN
  - name: fix-t99rt002-dictionary
    box: fix-t99rt002
    dictionary:
      name: fix50-aix
      type: MAIN
  - name: fix-t99rt003-dictionary
    box: fix-t99rt003
    dictionary:
      name: fix50-aix
      type: MAIN
  - name: fix-t99rt004-dictionary
    box: fix-t99rt004
    dictionary:
      name: fix50-aix-neg
      type: MAIN    
  - name: fix-t99rt005-dictionary
    box: fix-t99rt005
    dictionary:
      name: fix50-aix
      type: MAIN
  - name: fix-t99rt006-dictionary
    box: fix-t99rt006
    dictionary:
      name: fix50-aix
      type: MAIN
  - name: fix-t99rt009-dictionary
    box: fix-t99rt009
    dictionary:
      name: fix50-aix
      type: MAIN
  - name: fix-t99rt010-dictionary
    box: fix-t99rt010
    dictionary:
      name: fix50-aix
      type: MAIN   
  - name: fix-t99rt011-dictionary
    box: fix-t99rt011
    dictionary:
      name: fix50-aix
      type: MAIN
```


<h3 id="property-details-v1">Properties</h3>


<div class="property depth-0">
<div class="property-header">
<h3 class="property-path" id="v1-.apiVersion">.apiVersion</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-0">
<div class="property-header">
<h3 class="property-path" id="v1-.kind">.kind</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-0">
<div class="property-header">
<h3 class="property-path" id="v1-.metadata">.metadata</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-0">
<div class="property-header">
<h3 class="property-path" id="v1-.spec">.spec</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>Specification of desired links. Includes three main sections - boxesRelation, dictionariesRelation and multiDictionariesRelation</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation">.spec.boxesRelation</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>this section includes two subsections routerGrpc and routerMq. Note that the link is invalid if the boxes “from” and  “to” have the same name.</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerGrpc">.spec.boxesRelation.routerGrpc</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>List of grpc connection pins. Used to enable messaging between boxes.</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerGrpc[*]">.spec.boxesRelation.routerGrpc[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerGrpc[*].from">.spec.boxesRelation.routerGrpc[*].from</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerGrpc[*].from.box">.spec.boxesRelation.routerGrpc[*].from.box</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name of the box</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerGrpc[*].from.pin">.spec.boxesRelation.routerGrpc[*].from.pin</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>represents the pin of the box that we want to connect</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerGrpc[*].name">.spec.boxesRelation.routerGrpc[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name for the pin. must follow the regex pattern ^<a href="[-a-z0-9]*[a-z0-9]*[_a-z0-9]">a-z0-9</a>+$  and max length of 256</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerGrpc[*].to">.spec.boxesRelation.routerGrpc[*].to</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerGrpc[*].to.box">.spec.boxesRelation.routerGrpc[*].to.box</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name of the box</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerGrpc[*].to.pin">.spec.boxesRelation.routerGrpc[*].to.pin</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>represents the pin of the box that we want to connect</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerMq">.spec.boxesRelation.routerMq</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>List of mq connection pins. Used to enable messaging between boxes.</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerMq[*]">.spec.boxesRelation.routerMq[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerMq[*].from">.spec.boxesRelation.routerMq[*].from</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerMq[*].from.box">.spec.boxesRelation.routerMq[*].from.box</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name of the box</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerMq[*].from.pin">.spec.boxesRelation.routerMq[*].from.pin</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>represents the pin of the box that we want to connect</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerMq[*].name">.spec.boxesRelation.routerMq[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name for the pin. must follow the regex pattern ^<a href="[-a-z0-9]*[a-z0-9]*[_a-z0-9]">a-z0-9</a>+$  and max length of 256</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerMq[*].to">.spec.boxesRelation.routerMq[*].to</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerMq[*].to.box">.spec.boxesRelation.routerMq[*].to.box</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name of the box</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxesRelation.routerMq[*].to.pin">.spec.boxesRelation.routerMq[*].to.pin</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>represents the pin of the box that we want to connect</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.dictionariesRelation">.spec.dictionariesRelation</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>this section is used to link boxes with their associated dictionaries</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.dictionariesRelation[*]">.spec.dictionariesRelation[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.dictionariesRelation[*].box">.spec.dictionariesRelation[*].box</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name of the box that we want to link to the dictionary</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.dictionariesRelation[*].dictionary">.spec.dictionariesRelation[*].dictionary</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>section of dictionary</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.dictionariesRelation[*].dictionary.name">.spec.dictionariesRelation[*].dictionary.name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name of a dictionary</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.dictionariesRelation[*].dictionary.type">.spec.dictionariesRelation[*].dictionary.type</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>type of a dictionary</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.dictionariesRelation[*].name">.spec.dictionariesRelation[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name for the pin. must follow the regex pattern ^<a href="[-a-z0-9]*[a-z0-9]*[_a-z0-9]">a-z0-9</a>+$  and max length of 256</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.multiDictionariesRelation">.spec.multiDictionariesRelation</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>this section is used to link boxes with their associated multidictionaries</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.multiDictionariesRelation[*]">.spec.multiDictionariesRelation[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.multiDictionariesRelation[*].box">.spec.multiDictionariesRelation[*].box</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name of the box that we want to link to the multidictionary</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.multiDictionariesRelation[*].dictionaries">.spec.multiDictionariesRelation[*].dictionaries</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.multiDictionariesRelation[*].dictionaries[*]">.spec.multiDictionariesRelation[*].dictionaries[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.multiDictionariesRelation[*].dictionaries[*].alias">.spec.multiDictionariesRelation[*].dictionaries[*].alias</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>alias for the dictionary. must follow the regex pattern ^<a href="[-a-z0-9]*[a-z0-9]*[_a-z0-9]">a-z0-9</a>+$</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.multiDictionariesRelation[*].dictionaries[*].name">.spec.multiDictionariesRelation[*].dictionaries[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name for the dictionary. must follow the regex pattern ^<a href="[-a-z0-9]*[a-z0-9]*[_a-z0-9]">a-z0-9</a>+$</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.multiDictionariesRelation[*].name">.spec.multiDictionariesRelation[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name for the pin. must follow the regex pattern ^<a href="[-a-z0-9]*[a-z0-9]*[_a-z0-9]">a-z0-9</a>+$  and max length of 256</p>

</div>

</div>
</div>





</div>



