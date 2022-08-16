---
title: Th2Link CRD schema reference (group th2.exactpro.com)
linkTitle: Th2Link
description: |
  Th2link defines th2 link instance
weight: 100
crd:
  name_camelcase: Th2Link
  name_plural: th2links
  name_singular: th2link
  group: th2.exactpro.com
  technical_name: th2links.th2.exactpro.com
  scope: Namespaced
  source_repository: https://github.com/th2-net/th2-infra
  source_repository_ref: release-v1.8.0
  versions:
    - v1
layout: crd
owner:
  - https://github.com/th2-net/th2-infra
aliases:
  - /reference/cp-k8s-api/th2links.th2.exactpro.com/
technical_name: th2links.th2.exactpro.com
source_repository: https://github.com/th2-net/th2-infra
source_repository_ref: release-v1.8.0
---

# Th2Link


<p class="crd-description">Th2link defines th2 link instance</p>
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
<p>Specification of desired links</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxes-relation">.spec.boxes-relation</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxes-relation.router-grpc">.spec.boxes-relation.router-grpc</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>List of grpc connection pins</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxes-relation.router-grpc[*]">.spec.boxes-relation.router-grpc[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxes-relation.router-grpc[*].from">.spec.boxes-relation.router-grpc[*].from</h3>
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
<h3 class="property-path" id="v1-.spec.boxes-relation.router-grpc[*].from.box">.spec.boxes-relation.router-grpc[*].from.box</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxes-relation.router-grpc[*].from.pin">.spec.boxes-relation.router-grpc[*].from.pin</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxes-relation.router-grpc[*].name">.spec.boxes-relation.router-grpc[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxes-relation.router-grpc[*].to">.spec.boxes-relation.router-grpc[*].to</h3>
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
<h3 class="property-path" id="v1-.spec.boxes-relation.router-grpc[*].to.box">.spec.boxes-relation.router-grpc[*].to.box</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxes-relation.router-grpc[*].to.pin">.spec.boxes-relation.router-grpc[*].to.pin</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxes-relation.router-mq">.spec.boxes-relation.router-mq</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>List of mq connection pins</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxes-relation.router-mq[*]">.spec.boxes-relation.router-mq[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxes-relation.router-mq[*].from">.spec.boxes-relation.router-mq[*].from</h3>
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
<h3 class="property-path" id="v1-.spec.boxes-relation.router-mq[*].from.box">.spec.boxes-relation.router-mq[*].from.box</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxes-relation.router-mq[*].from.pin">.spec.boxes-relation.router-mq[*].from.pin</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxes-relation.router-mq[*].name">.spec.boxes-relation.router-mq[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxes-relation.router-mq[*].to">.spec.boxes-relation.router-mq[*].to</h3>
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
<h3 class="property-path" id="v1-.spec.boxes-relation.router-mq[*].to.box">.spec.boxes-relation.router-mq[*].to.box</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.boxes-relation.router-mq[*].to.pin">.spec.boxes-relation.router-mq[*].to.pin</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.dictionaries-relation">.spec.dictionaries-relation</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>List of dictionary relations</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.dictionaries-relation[*]">.spec.dictionaries-relation[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.dictionaries-relation[*].box">.spec.dictionaries-relation[*].box</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.dictionaries-relation[*].dictionary">.spec.dictionaries-relation[*].dictionary</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.dictionaries-relation[*].dictionary.name">.spec.dictionaries-relation[*].dictionary.name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.dictionaries-relation[*].dictionary.type">.spec.dictionaries-relation[*].dictionary.type</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.dictionaries-relation[*].name">.spec.dictionaries-relation[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.multi-dictionaries-relation">.spec.multi-dictionaries-relation</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>List of multi dictionary relations</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.multi-dictionaries-relation[*]">.spec.multi-dictionaries-relation[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.multi-dictionaries-relation[*].box">.spec.multi-dictionaries-relation[*].box</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.multi-dictionaries-relation[*].dictionaries">.spec.multi-dictionaries-relation[*].dictionaries</h3>
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
<h3 class="property-path" id="v1-.spec.multi-dictionaries-relation[*].dictionaries[*]">.spec.multi-dictionaries-relation[*].dictionaries[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.multi-dictionaries-relation[*].dictionaries[*].alias">.spec.multi-dictionaries-relation[*].dictionaries[*].alias</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.multi-dictionaries-relation[*].dictionaries[*].name">.spec.multi-dictionaries-relation[*].dictionaries[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.multi-dictionaries-relation[*].name">.spec.multi-dictionaries-relation[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>





</div>



