---
title: Th2CoreBox CRD schema reference (group th2.exactpro.com)
linkTitle: Th2CoreBox
description: |
  Th2CoreBox defines Th2 core component instance
weight: 100
crd:
  name_camelcase: Th2CoreBox
  name_plural: th2coreboxes
  name_singular: th2corebox
  group: th2.exactpro.com
  technical_name: th2coreboxes.th2.exactpro.com
  scope: Namespaced
  source_repository: https://github.com/th2-net/th2-infra
  source_repository_ref: release-v1.8.0
  versions:
    - v1
  topics:
layout: crd
owner:
  - https://github.com/th2-net/th2-infra
aliases:
  - /reference/cp-k8s-api/th2coreboxes.th2.exactpro.com/
technical_name: th2coreboxes.th2.exactpro.com
source_repository: https://github.com/th2-net/th2-infra
source_repository_ref: release-v1.8.0
---

# Th2CoreBox


<p class="crd-description">Th2CoreBox defines Th2 core component instance</p>
<dl class="crd-meta">
<dt class="fullname">Full name:</dt>
<dd class="fullname">th2coreboxes.th2.exactpro.com</dd>
<dt class="groupname">Group:</dt>
<dd class="groupname">th2.exactpro.com</dd>
<dt class="singularname">Singular name:</dt>
<dd class="singularname">th2corebox</dd>
<dt class="pluralname">Plural name:</dt>
<dd class="pluralname">th2coreboxes</dd>
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
<p>Specification of desired box</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.cradleManager">.spec.cradleManager</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>this is custom configuration for cradle</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.custom-config">.spec.custom-config</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>this is custom configuration</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.disabled">.spec.disabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings">.spec.extended-settings</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.bookName">.spec.extended-settings.bookName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.chart-cfg">.spec.extended-settings.chart-cfg</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.chart-cfg.path">.spec.extended-settings.chart-cfg.path</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.chart-cfg.ref">.spec.extended-settings.chart-cfg.ref</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.envVariables">.spec.extended-settings.envVariables</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.externalBox">.spec.extended-settings.externalBox</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.externalBox.address">.spec.extended-settings.externalBox.address</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.externalBox.enabled">.spec.extended-settings.externalBox.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.externalBox.endpoints">.spec.extended-settings.externalBox.endpoints</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.externalBox.endpoints[*]">.spec.extended-settings.externalBox.endpoints[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.externalBox.endpoints[*].name">.spec.extended-settings.externalBox.endpoints[*].name</h3>
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
<h3 class="property-path" id="v1-.spec.extended-settings.externalBox.endpoints[*].targetPort">.spec.extended-settings.externalBox.endpoints[*].targetPort</h3>
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
<h3 class="property-path" id="v1-.spec.extended-settings.hostAliases">.spec.extended-settings.hostAliases</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>HostAliases is an optional list of hosts and IPs that will be injected into the pod&rsquo;s hosts file if specified. This is only valid for non-hostNetwork pods.</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.hostAliases[*]">.spec.extended-settings.hostAliases[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.hostAliases[*].hostnames">.spec.extended-settings.hostAliases[*].hostnames</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>Hostnames for the above IP address.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.hostAliases[*].hostnames[*]">.spec.extended-settings.hostAliases[*].hostnames[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.hostAliases[*].ip">.spec.extended-settings.hostAliases[*].ip</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>IP address of the host file entry.</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.hostNetwork">.spec.extended-settings.hostNetwork</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.k8sProbes">.spec.extended-settings.k8sProbes</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.mounting">.spec.extended-settings.mounting</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.mounting[*]">.spec.extended-settings.mounting[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.mounting[*].path">.spec.extended-settings.mounting[*].path</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.mounting[*].pvcName">.spec.extended-settings.mounting[*].pvcName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.nodeSelector">.spec.extended-settings.nodeSelector</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>map for selecting nodes</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.replicas">.spec.extended-settings.replicas</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.resources">.spec.extended-settings.resources</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.resources.limits">.spec.extended-settings.resources.limits</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.resources.limits.cpu">.spec.extended-settings.resources.limits.cpu</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.resources.limits.memory">.spec.extended-settings.resources.limits.memory</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.resources.requests">.spec.extended-settings.resources.requests</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.resources.requests.cpu">.spec.extended-settings.resources.requests.cpu</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.resources.requests.memory">.spec.extended-settings.resources.requests.memory</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.service">.spec.extended-settings.service</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.service.enabled">.spec.extended-settings.service.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.service.endpoints">.spec.extended-settings.service.endpoints</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.service.endpoints[*]">.spec.extended-settings.service.endpoints[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.service.endpoints[*].name">.spec.extended-settings.service.endpoints[*].name</h3>
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
<h3 class="property-path" id="v1-.spec.extended-settings.service.endpoints[*].nodePort">.spec.extended-settings.service.endpoints[*].nodePort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.service.endpoints[*].port">.spec.extended-settings.service.endpoints[*].port</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.service.endpoints[*].targetPort">.spec.extended-settings.service.endpoints[*].targetPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.service.ingress">.spec.extended-settings.service.ingress</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>Creates Ingress for the Box</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.service.ingress.urlPaths">.spec.extended-settings.service.ingress.urlPaths</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>Path is a suffix, it is appended to namespace url, e.g for th2-test namespace:</p>

<p>urlPaths:
- /my-box/</p>

<p>service will be available at /th2-test/my-box/</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.service.ingress.urlPaths[*]">.spec.extended-settings.service.ingress.urlPaths[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.service.nodePort">.spec.extended-settings.service.nodePort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.service.targetPort">.spec.extended-settings.service.targetPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.service.type">.spec.extended-settings.service.type</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.sharedMemory">.spec.extended-settings.sharedMemory</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.extended-settings.sharedMemory.enabled">.spec.extended-settings.sharedMemory.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.grpcRouter">.spec.grpcRouter</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>this is custom configuration for rouer grpc</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.image-name">.spec.image-name</h3>
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
<h3 class="property-path" id="v1-.spec.image-version">.spec.image-version</h3>
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
<h3 class="property-path" id="v1-.spec.loggingConfig">.spec.loggingConfig</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.mqRouter">.spec.mqRouter</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>this is custom configuration for rouer mq</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.params">.spec.params</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>List of generic parameters</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.params[*]">.spec.params[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.params[*].name">.spec.params[*].name</h3>
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
<h3 class="property-path" id="v1-.spec.params[*].value">.spec.params[*].value</h3>
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
<h3 class="property-path" id="v1-.spec.pins">.spec.pins</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>List of connection pins</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*]">.spec.pins[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].attributes">.spec.pins[*].attributes</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].attributes[*]">.spec.pins[*].attributes[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].connection-type">.spec.pins[*].connection-type</h3>
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
<h3 class="property-path" id="v1-.spec.pins[*].filters">.spec.pins[*].filters</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].filters[*]">.spec.pins[*].filters[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].filters[*].message">.spec.pins[*].filters[*].message</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

</div>
</div>

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].filters[*].message[*]">.spec.pins[*].filters[*].message[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].filters[*].message[*].expected-value">.spec.pins[*].filters[*].message[*].expected-value</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].filters[*].message[*].field-name">.spec.pins[*].filters[*].message[*].field-name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].filters[*].message[*].operation">.spec.pins[*].filters[*].message[*].operation</h3>
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
<h3 class="property-path" id="v1-.spec.pins[*].filters[*].metadata">.spec.pins[*].filters[*].metadata</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

</div>
</div>

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].filters[*].metadata[*]">.spec.pins[*].filters[*].metadata[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].filters[*].metadata[*].expected-value">.spec.pins[*].filters[*].metadata[*].expected-value</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].filters[*].metadata[*].field-name">.spec.pins[*].filters[*].metadata[*].field-name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].filters[*].metadata[*].operation">.spec.pins[*].filters[*].metadata[*].operation</h3>
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
<h3 class="property-path" id="v1-.spec.pins[*].filters[*].properties">.spec.pins[*].filters[*].properties</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

</div>
</div>

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].filters[*].properties[*]">.spec.pins[*].filters[*].properties[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].filters[*].properties[*].expected-value">.spec.pins[*].filters[*].properties[*].expected-value</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].filters[*].properties[*].field-name">.spec.pins[*].filters[*].properties[*].field-name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].filters[*].properties[*].operation">.spec.pins[*].filters[*].properties[*].operation</h3>
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
<h3 class="property-path" id="v1-.spec.pins[*].name">.spec.pins[*].name</h3>
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
<h3 class="property-path" id="v1-.spec.pins[*].service-class">.spec.pins[*].service-class</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].service-classes">.spec.pins[*].service-classes</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].service-classes[*]">.spec.pins[*].service-classes[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].settings">.spec.pins[*].settings</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].settings.overloadStrategy">.spec.pins[*].settings.overloadStrategy</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].settings.queueLength">.spec.pins[*].settings.queueLength</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].settings.storageOnDemand">.spec.pins[*].settings.storageOnDemand</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.pins[*].strategy">.spec.pins[*].strategy</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.prometheus">.spec.prometheus</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.prometheus.enabled">.spec.prometheus.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.prometheus.host">.spec.prometheus.host</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.prometheus.port">.spec.prometheus.port</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.spec.type">.spec.type</h3>
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
<h3 class="property-path" id="v1-.spec.version-range">.spec.version-range</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-0">
<div class="property-header">
<h3 class="property-path" id="v1-.status">.status</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>contains status information about this resource</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.status.conditions">.status.conditions</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>Conditions contains observations of the resource&rsquo;s state</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v1-.status.conditions[*]">.status.conditions[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.status.conditions[*].lastTransitionTime">.status.conditions[*].lastTransitionTime</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>LastTransitionTime is the timestamp corresponding to the last status change of this condition.</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.status.conditions[*].lastUpdateTime">.status.conditions[*].lastUpdateTime</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>LastUpdateTime is the timestamp corresponding to the last status update of this condition.</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.status.conditions[*].message">.status.conditions[*].message</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>Message is a human readable description of the details of the last transition, complementing reason.</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.status.conditions[*].reason">.status.conditions[*].reason</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>Reason is a brief machine readable explanation for the condition&rsquo;s last transition.</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.status.conditions[*].status">.status.conditions[*].status</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>Status of the condition, one of (&lsquo;True&rsquo;, &lsquo;False&rsquo;, &lsquo;Unknown&rsquo;).</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v1-.status.conditions[*].type">.status.conditions[*].type</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>Type of the condition</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.status.message">.status.message</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>Message describe current state of this resource</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.status.phase">.status.phase</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>ComponentStatus is the status as given by Operator for this resource</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v1-.status.subResourceName">.status.subResourceName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>SubResourceName is the name of produced helmrelease tied with this resource</p>

</div>

</div>
</div>





</div>



