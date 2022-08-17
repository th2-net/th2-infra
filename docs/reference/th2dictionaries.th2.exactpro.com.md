---
title: Th2Dictionary CRD schema reference (group th2.exactpro.com)
linkTitle: Th2Dictionary
description: |
  Th2Dictionary defines th2 dictionary instance
weight: 100
crd:
  name_camelcase: Th2Dictionary
  name_plural: th2dictionaries
  name_singular: th2dictionary
  group: th2.exactpro.com
  technical_name: th2dictionaries.th2.exactpro.com
  scope: Namespaced
  source_repository: https://github.com/th2-net/th2-infra
  source_repository_ref: release-v1.8.0
  versions:
    - v1
layout: crd
owner:
  - https://github.com/th2-net/th2-infra
aliases:
  - /reference/cp-k8s-api/th2dictionaries.th2.exactpro.com/
technical_name: th2dictionaries.th2.exactpro.com
source_repository: https://github.com/th2-net/th2-infra
source_repository_ref: release-v1.8.0
---

# Th2Dictionary


<p class="crd-description">Th2Dictionary defines th2 dictionary instance</p>
<dl class="crd-meta">
<dt class="fullname">Full name:</dt>
<dd class="fullname">th2dictionaries.th2.exactpro.com</dd>
<dt class="groupname">Group:</dt>
<dd class="groupname">th2.exactpro.com</dd>
<dt class="singularname">Singular name:</dt>
<dd class="singularname">th2dictionary</dd>
<dt class="pluralname">Plural name:</dt>
<dd class="pluralname">th2dictionaries</dd>
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
<p>Specification of desired dictionary</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.compressed">.spec.compressed</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.data">.spec.data</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>





</div>



