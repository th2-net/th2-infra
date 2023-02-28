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
  source_repository_ref: Release-v2.0.0-conflict-solve
  versions:
    - v2
layout: crd
owner:
  - https://github.com/th2-net/th2-infra
aliases:
  - /reference/cp-k8s-api/th2dictionaries.th2.exactpro.com/
technical_name: th2dictionaries.th2.exactpro.com
source_repository: https://github.com/th2-net/th2-infra
source_repository_ref: Release-v2.0.0-conflict-solve
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
<dd class="versions"><a class="version" href="#v2" title="Show schema for version v2">v2</a></dd>
</dl>



<div class="crd-schema-version">
<h2 id="v2">Version v2</h2>



<h3 id="property-details-v2">Properties</h3>


<div class="property depth-0">
<div class="property-header">
<h3 class="property-path" id="v2-.apiVersion">.apiVersion</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-0">
<div class="property-header">
<h3 class="property-path" id="v2-.kind">.kind</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-0">
<div class="property-header">
<h3 class="property-path" id="v2-.metadata">.metadata</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-0">
<div class="property-header">
<h3 class="property-path" id="v2-.spec">.spec</h3>
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
<h3 class="property-path" id="v2-.spec.compressed">.spec.compressed</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>Indicates whether dictionary&rsquo;s data is already compressed. If set to true, no further compression will take place in infra. default is set to false</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.data">.spec.data</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>data for the dictionary</p>

</div>

</div>
</div>





</div>



