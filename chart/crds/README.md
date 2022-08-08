---
title: Th2Box CRD schema reference (group th2.exactpro.com)
linkTitle: Th2Box
description: |
  Th2Box defines generic Th2 component instance
weight: 100
crd:
  name_camelcase: Th2Box
  name_plural: th2boxes
  name_singular: th2box
  group: th2.exactpro.com
  technical_name: th2boxes.th2.exactpro.com
  scope: Namespaced
  source_repository: https://github.com/th2-net/th2-infra
  source_repository_ref: release-v2.0.0
  versions:
    - v2
  topics:
    - managementcluster
layout: crd
owner:
  - https://github.com/th2-net/th2-infra
aliases:
  - /reference/cp-k8s-api/th2boxes.th2.exactpro.com/
technical_name: th2boxes.th2.exactpro.com
source_repository: https://github.com/th2-net/th2-infra
source_repository_ref: release-v2.0.0
---

# Th2Box


<p class="crd-description">Th2Box defines generic Th2 component instance</p>
<dl class="crd-meta">
<dt class="fullname">Full name:</dt>
<dd class="fullname">th2boxes.th2.exactpro.com</dd>
<dt class="groupname">Group:</dt>
<dd class="groupname">th2.exactpro.com</dd>
<dt class="singularname">Singular name:</dt>
<dd class="singularname">th2box</dd>
<dt class="pluralname">Plural name:</dt>
<dd class="pluralname">th2boxes</dd>
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
<p>Specification of desired box</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.bookName">.spec.bookName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>Can be used for passing custom book for a specific CR</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.cradleManager">.spec.cradleManager</h3>
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
<h3 class="property-path" id="v2-.spec.customConfig">.spec.customConfig</h3>
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
<h3 class="property-path" id="v2-.spec.disabled">.spec.disabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if set to true resource will act as if it is removed from schema without actually deleting the file. The default value is false.</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings">.spec.extendedSettings</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>section for extended settings</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.chartCfg">.spec.extendedSettings.chartCfg</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>specifies which charts to use for rolling out infra</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.chartCfg.path">.spec.extendedSettings.chartCfg.path</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.chartCfg.ref">.spec.extendedSettings.chartCfg.ref</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.envVariables">.spec.extendedSettings.envVariables</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>allows us to pass specific environment variables that are going to be set into the pods.</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox">.spec.extendedSettings.externalBox</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>using this section we can configure boxes that are going to be run outside of kubernetes cluster</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.address">.spec.extendedSettings.externalBox.address</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>address to the machine on which external box is running</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.enabled">.spec.extendedSettings.externalBox.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>represents the state of the external box</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints">.spec.extendedSettings.externalBox.endpoints</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>internal mapping for ports</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints[*]">.spec.extendedSettings.externalBox.endpoints[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints[*].name">.spec.extendedSettings.externalBox.endpoints[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name for the endpoint</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints[*].targetPort">.spec.extendedSettings.externalBox.endpoints[*].targetPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>value for exposed port</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases">.spec.extendedSettings.hostAliases</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*]">.spec.extendedSettings.hostAliases[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*].hostnames">.spec.extendedSettings.hostAliases[*].hostnames</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*].hostnames[*]">.spec.extendedSettings.hostAliases[*].hostnames[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*].ip">.spec.extendedSettings.hostAliases[*].ip</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.hostNetwork">.spec.extendedSettings.hostNetwork</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if the flag is set to true, pod will run on node network and kubernetes will decide which node will be used for running the box.</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.k8sProbes">.spec.extendedSettings.k8sProbes</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if enabled liveness probes will be collected from pod</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting">.spec.extendedSettings.mounting</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>allows to configure persistent volume mounting for pods</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting[*]">.spec.extendedSettings.mounting[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting[*].path">.spec.extendedSettings.mounting[*].path</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>a directory in pod where you mount external folder</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting[*].pvcName">.spec.extendedSettings.mounting[*].pvcName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name for the pvc</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.nodeSelector">.spec.extendedSettings.nodeSelector</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.replicas">.spec.extendedSettings.replicas</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>number of replicas</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources">.spec.extendedSettings.resources</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>here we can specify resource limitations and allowances for this specific component</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.limits">.spec.extendedSettings.resources.limits</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.limits.cpu">.spec.extendedSettings.resources.limits.cpu</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.limits.memory">.spec.extendedSettings.resources.limits.memory</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.requests">.spec.extendedSettings.resources.requests</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.requests.cpu">.spec.extendedSettings.resources.requests.cpu</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.requests.memory">.spec.extendedSettings.resources.requests.memory</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service">.spec.extendedSettings.service</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP">.spec.extendedSettings.service.clusterIP</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>endpoints for cluster IP service type</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP[*]">.spec.extendedSettings.service.clusterIP[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP[*].containerPort">.spec.extendedSettings.service.clusterIP[*].containerPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>allows to specify Kubernetes port for the pod.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP[*].name">.spec.extendedSettings.service.clusterIP[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name of the endpoint</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.enabled">.spec.extendedSettings.service.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if enabled allows pod to be exposed using ClusterIP</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.ingress">.spec.extendedSettings.service.ingress</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>section for ingress</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.ingress.urlPaths">.spec.extendedSettings.service.ingress.urlPaths</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>url paths</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.ingress.urlPaths[*]">.spec.extendedSettings.service.ingress.urlPaths[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort">.spec.extendedSettings.service.nodePort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>endpoints for cluster IP service type</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*]">.spec.extendedSettings.service.nodePort[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*].containerPort">.spec.extendedSettings.service.nodePort[*].containerPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>allows to specify Kubernetes port for the pod.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*].exposedPort">.spec.extendedSettings.service.nodePort[*].exposedPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>a port on which the service will be accessible.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*].name">.spec.extendedSettings.service.nodePort[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name of the endpoint</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.sharedMemory">.spec.extendedSettings.sharedMemory</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.sharedMemory.enabled">.spec.extendedSettings.sharedMemory.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.grpcRouter">.spec.grpcRouter</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>this is custom configuration for router grpc</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.imageName">.spec.imageName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>docker image repository URL</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.imageVersion">.spec.imageVersion</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>docker image tag</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.loggingConfig">.spec.loggingConfig</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>this is logging configuration</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.mqRouter">.spec.mqRouter</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>this is custom configuration for router mq</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins">.spec.pins</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>consists of grpc(server, client) and mq pin sections</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc">.spec.pins.grpc</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>pin section for grpc</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client">.spec.pins.grpc.client</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>grpc client subsection</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*]">.spec.pins.grpc.client[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].attributes">.spec.pins.grpc.client[*].attributes</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>attributes of grpc client</p>

</div>

</div>
</div>

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].attributes[*]">.spec.pins.grpc.client[*].attributes[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].filters">.spec.pins.grpc.client[*].filters</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>filters for grpc client</p>

</div>

</div>
</div>

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].filters[*]">.spec.pins.grpc.client[*].filters[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].filters[*].properties">.spec.pins.grpc.client[*].filters[*].properties</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].filters[*].properties[*]">.spec.pins.grpc.client[*].filters[*].properties[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-9">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].filters[*].properties[*].expectedValue">.spec.pins.grpc.client[*].filters[*].properties[*].expectedValue</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-9">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].filters[*].properties[*].fieldName">.spec.pins.grpc.client[*].filters[*].properties[*].fieldName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-9">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].filters[*].properties[*].operation">.spec.pins.grpc.client[*].filters[*].properties[*].operation</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].name">.spec.pins.grpc.client[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name for the pin. must follow the regex pattern ^<a href="[-a-z0-9]*[a-z0-9]*[_a-z0-9]">a-z0-9</a>+$ and max length of 71 characters</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].serviceClass">.spec.pins.grpc.client[*].serviceClass</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name of the service class used supported by the grpc client</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].strategy">.spec.pins.grpc.client[*].strategy</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>strategy to be used for grpc communication. default value is set to robin</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.server">.spec.pins.grpc.server</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>grpc server subsection</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.server[*]">.spec.pins.grpc.server[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.server[*].name">.spec.pins.grpc.server[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name for the pin. must follow the regex pattern ^<a href="[-a-z0-9]*[a-z0-9]*[_a-z0-9]">a-z0-9</a>+$ and max length of 71 characters</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.server[*].serviceClasses">.spec.pins.grpc.server[*].serviceClasses</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>list of the service classes supported by the grpc server</p>

</div>

</div>
</div>

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.server[*].serviceClasses[*]">.spec.pins.grpc.server[*].serviceClasses[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq">.spec.pins.mq</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>pin section for mq</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*]">.spec.pins.mq[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].attributes">.spec.pins.mq[*].attributes</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>attributes for mq</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].attributes[*]">.spec.pins.mq[*].attributes[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters">.spec.pins.mq[*].filters</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>filters for mq</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*]">.spec.pins.mq[*].filters[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].message">.spec.pins.mq[*].filters[*].message</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].message[*]">.spec.pins.mq[*].filters[*].message[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].message[*].expectedValue">.spec.pins.mq[*].filters[*].message[*].expectedValue</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].message[*].fieldName">.spec.pins.mq[*].filters[*].message[*].fieldName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].message[*].operation">.spec.pins.mq[*].filters[*].message[*].operation</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].metadata">.spec.pins.mq[*].filters[*].metadata</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>metadata for mq pin</p>

</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].metadata[*]">.spec.pins.mq[*].filters[*].metadata[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].metadata[*].expectedValue">.spec.pins.mq[*].filters[*].metadata[*].expectedValue</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].metadata[*].fieldName">.spec.pins.mq[*].filters[*].metadata[*].fieldName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].metadata[*].operation">.spec.pins.mq[*].filters[*].metadata[*].operation</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].properties">.spec.pins.mq[*].filters[*].properties</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].properties[*]">.spec.pins.mq[*].filters[*].properties[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].properties[*].expectedValue">.spec.pins.mq[*].filters[*].properties[*].expectedValue</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].properties[*].fieldName">.spec.pins.mq[*].filters[*].properties[*].fieldName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].properties[*].operation">.spec.pins.mq[*].filters[*].properties[*].operation</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].name">.spec.pins.mq[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name for the pin. must follow the regex pattern ^<a href="[-a-z0-9]*[a-z0-9]*[_a-z0-9]">a-z0-9</a>+$ and max length of 71 characters</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].settings">.spec.pins.mq[*].settings</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>settings for rabbitMq queue configuration</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].settings.overloadStrategy">.spec.pins.mq[*].settings.overloadStrategy</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>configuration for rabbit mq queue. default is set to “drop-head”.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].settings.queueLength">.spec.pins.mq[*].settings.queueLength</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>configuration for rabbit mq queue. default is set to 1000 msg. queueLength isn&rsquo;t used if storageOnDemand is set to true.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].settings.storageOnDemand">.spec.pins.mq[*].settings.storageOnDemand</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>configuration for rabbit mq queue. default value is set to false</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus">.spec.prometheus</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>custom configuration of prometheus for microservices</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus.enabled">.spec.prometheus.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>By default this is set to true.</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus.host">.spec.prometheus.host</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>host for prometheus</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus.port">.spec.prometheus.port</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>port for prometheus</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.type">.spec.type</h3>
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
<h3 class="property-path" id="v2-.spec.versionRange">.spec.versionRange</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>when the appropriate pattern is provided, image-version will be updated with the latest tag from the image repository that satisfies that pattern.</p>

</div>

</div>
</div>

<div class="property depth-0">
<div class="property-header">
<h3 class="property-path" id="v2-.status">.status</h3>
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
<h3 class="property-path" id="v2-.status.conditions">.status.conditions</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*]">.status.conditions[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.status.conditions[*].lastTransitionTime">.status.conditions[*].lastTransitionTime</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].lastUpdateTime">.status.conditions[*].lastUpdateTime</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].message">.status.conditions[*].message</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].reason">.status.conditions[*].reason</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].status">.status.conditions[*].status</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].type">.status.conditions[*].type</h3>
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
<h3 class="property-path" id="v2-.status.message">.status.message</h3>
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
<h3 class="property-path" id="v2-.status.phase">.status.phase</h3>
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
<h3 class="property-path" id="v2-.status.subResourceName">.status.subResourceName</h3>
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



---
title: Th2CoreBox CRD schema reference (group th2.exactpro.com)
linkTitle: Th2CoreBox
description: |
  Custom resource definition (CRD) schema reference page for the Th2CoreBox resource (th2coreboxes.th2.exactpro.com), as part of the Giant Swarm Management API documentation.
weight: 100
crd:
  name_camelcase: Th2CoreBox
  name_plural: th2coreboxes
  name_singular: th2corebox
  group: th2.exactpro.com
  technical_name: th2coreboxes.th2.exactpro.com
  scope: Namespaced
  source_repository: https://github.com/th2-net/th2-infra
  source_repository_ref: release-v2.0.0
  versions:
    - v2
  topics:
    - managementcluster
layout: crd
owner:
  - https://github.com/th2-net/th2-infra
aliases:
  - /reference/cp-k8s-api/th2coreboxes.th2.exactpro.com/
technical_name: th2coreboxes.th2.exactpro.com
source_repository: https://github.com/th2-net/th2-infra
source_repository_ref: release-v2.0.0
---

# Th2CoreBox

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
<p>Specification of desired box</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.bookName">.spec.bookName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>Can be used for passing custom book for a specific CR</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.cradleManager">.spec.cradleManager</h3>
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
<h3 class="property-path" id="v2-.spec.customConfig">.spec.customConfig</h3>
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
<h3 class="property-path" id="v2-.spec.disabled">.spec.disabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if set to true resource will act as if it is removed from schema without actually deleting the file. The default value is false.</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings">.spec.extendedSettings</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>section for extended settings</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.chartCfg">.spec.extendedSettings.chartCfg</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>specifies which charts to use for rolling out infra</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.chartCfg.path">.spec.extendedSettings.chartCfg.path</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.chartCfg.ref">.spec.extendedSettings.chartCfg.ref</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.envVariables">.spec.extendedSettings.envVariables</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>allows us to pass specific environment variables that are going to be set into the pods.</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox">.spec.extendedSettings.externalBox</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>using this section we can configure boxes that are going to be run outside of kubernetes cluster</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.address">.spec.extendedSettings.externalBox.address</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>address to the machine on which external box is running</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.enabled">.spec.extendedSettings.externalBox.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>represents the state of the external box</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints">.spec.extendedSettings.externalBox.endpoints</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>internal mapping for ports</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints[*]">.spec.extendedSettings.externalBox.endpoints[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints[*].name">.spec.extendedSettings.externalBox.endpoints[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name for the endpoint</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints[*].targetPort">.spec.extendedSettings.externalBox.endpoints[*].targetPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>value for exposed port</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases">.spec.extendedSettings.hostAliases</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*]">.spec.extendedSettings.hostAliases[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*].hostnames">.spec.extendedSettings.hostAliases[*].hostnames</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*].hostnames[*]">.spec.extendedSettings.hostAliases[*].hostnames[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*].ip">.spec.extendedSettings.hostAliases[*].ip</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.hostNetwork">.spec.extendedSettings.hostNetwork</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if the flag is set to true, pod will run on node network and kubernetes will decide which node will be used for running the box.</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.k8sProbes">.spec.extendedSettings.k8sProbes</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if enabled liveness probes will be collected from pod</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting">.spec.extendedSettings.mounting</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>allows to configure persistent volume mounting for pods</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting[*]">.spec.extendedSettings.mounting[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting[*].path">.spec.extendedSettings.mounting[*].path</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>a directory in pod where you mount external folder</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting[*].pvcName">.spec.extendedSettings.mounting[*].pvcName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name for the pvc</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.nodeSelector">.spec.extendedSettings.nodeSelector</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.replicas">.spec.extendedSettings.replicas</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>number of replicas</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources">.spec.extendedSettings.resources</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>here we can specify resource limitations and allowances for this specific component</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.limits">.spec.extendedSettings.resources.limits</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.limits.cpu">.spec.extendedSettings.resources.limits.cpu</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.limits.memory">.spec.extendedSettings.resources.limits.memory</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.requests">.spec.extendedSettings.resources.requests</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.requests.cpu">.spec.extendedSettings.resources.requests.cpu</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.requests.memory">.spec.extendedSettings.resources.requests.memory</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service">.spec.extendedSettings.service</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP">.spec.extendedSettings.service.clusterIP</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>endpoints for cluster IP service type</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP[*]">.spec.extendedSettings.service.clusterIP[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP[*].containerPort">.spec.extendedSettings.service.clusterIP[*].containerPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>allows to specify Kubernetes port for the pod.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP[*].name">.spec.extendedSettings.service.clusterIP[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name of the endpoint</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.enabled">.spec.extendedSettings.service.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if enabled allows pod to be exposed using ClusterIP</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.ingress">.spec.extendedSettings.service.ingress</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>section for ingress</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.ingress.urlPaths">.spec.extendedSettings.service.ingress.urlPaths</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>url paths</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.ingress.urlPaths[*]">.spec.extendedSettings.service.ingress.urlPaths[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort">.spec.extendedSettings.service.nodePort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>endpoints for cluster IP service type</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*]">.spec.extendedSettings.service.nodePort[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*].containerPort">.spec.extendedSettings.service.nodePort[*].containerPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>allows to specify Kubernetes port for the pod.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*].exposedPort">.spec.extendedSettings.service.nodePort[*].exposedPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>a port on which the service will be accessible.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*].name">.spec.extendedSettings.service.nodePort[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name of the endpoint</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.sharedMemory">.spec.extendedSettings.sharedMemory</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.sharedMemory.enabled">.spec.extendedSettings.sharedMemory.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.grpcRouter">.spec.grpcRouter</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>this is custom configuration for router grpc</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.imageName">.spec.imageName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>docker image repository URL</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.imageVersion">.spec.imageVersion</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>docker image tag</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.loggingConfig">.spec.loggingConfig</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>this is logging configuration</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.mqRouter">.spec.mqRouter</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>this is custom configuration for router mq</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.params">.spec.params</h3>
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
<h3 class="property-path" id="v2-.spec.params[*]">.spec.params[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.params[*].name">.spec.params[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name of parameter</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.params[*].value">.spec.params[*].value</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>value of parameter</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins">.spec.pins</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>consists of grpc(server, client) and mq pin sections</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc">.spec.pins.grpc</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>pin section for grpc</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client">.spec.pins.grpc.client</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>grpc client subsection</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*]">.spec.pins.grpc.client[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].attributes">.spec.pins.grpc.client[*].attributes</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>attributes of grpc client</p>

</div>

</div>
</div>

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].attributes[*]">.spec.pins.grpc.client[*].attributes[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].filters">.spec.pins.grpc.client[*].filters</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>filters for grpc client</p>

</div>

</div>
</div>

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].filters[*]">.spec.pins.grpc.client[*].filters[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].filters[*].properties">.spec.pins.grpc.client[*].filters[*].properties</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].filters[*].properties[*]">.spec.pins.grpc.client[*].filters[*].properties[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-9">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].filters[*].properties[*].expectedValue">.spec.pins.grpc.client[*].filters[*].properties[*].expectedValue</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-9">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].filters[*].properties[*].fieldName">.spec.pins.grpc.client[*].filters[*].properties[*].fieldName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-9">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].filters[*].properties[*].operation">.spec.pins.grpc.client[*].filters[*].properties[*].operation</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].name">.spec.pins.grpc.client[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name for the pin. must follow the regex pattern ^<a href="[-a-z0-9]*[a-z0-9]*[_a-z0-9]">a-z0-9</a>+$ and max length of 71 characters</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].serviceClass">.spec.pins.grpc.client[*].serviceClass</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name of the service class used supported by the grpc client</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].strategy">.spec.pins.grpc.client[*].strategy</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>strategy to be used for grpc communication. default value is set to robin</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.server">.spec.pins.grpc.server</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>grpc server subsection</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.server[*]">.spec.pins.grpc.server[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.server[*].name">.spec.pins.grpc.server[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name for the pin. must follow the regex pattern ^<a href="[-a-z0-9]*[a-z0-9]*[_a-z0-9]">a-z0-9</a>+$ and max length of 71 characters</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.server[*].serviceClasses">.spec.pins.grpc.server[*].serviceClasses</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>list of the service classes supported by the grpc server</p>

</div>

</div>
</div>

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.server[*].serviceClasses[*]">.spec.pins.grpc.server[*].serviceClasses[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq">.spec.pins.mq</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>pin section for mq</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*]">.spec.pins.mq[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].attributes">.spec.pins.mq[*].attributes</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>attributes for mq</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].attributes[*]">.spec.pins.mq[*].attributes[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters">.spec.pins.mq[*].filters</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>filters for mq</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*]">.spec.pins.mq[*].filters[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].message">.spec.pins.mq[*].filters[*].message</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].message[*]">.spec.pins.mq[*].filters[*].message[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].message[*].expectedValue">.spec.pins.mq[*].filters[*].message[*].expectedValue</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].message[*].fieldName">.spec.pins.mq[*].filters[*].message[*].fieldName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].message[*].operation">.spec.pins.mq[*].filters[*].message[*].operation</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].metadata">.spec.pins.mq[*].filters[*].metadata</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>metadata for mq pin</p>

</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].metadata[*]">.spec.pins.mq[*].filters[*].metadata[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].metadata[*].expectedValue">.spec.pins.mq[*].filters[*].metadata[*].expectedValue</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].metadata[*].fieldName">.spec.pins.mq[*].filters[*].metadata[*].fieldName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].metadata[*].operation">.spec.pins.mq[*].filters[*].metadata[*].operation</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].properties">.spec.pins.mq[*].filters[*].properties</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].properties[*]">.spec.pins.mq[*].filters[*].properties[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].properties[*].expectedValue">.spec.pins.mq[*].filters[*].properties[*].expectedValue</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].properties[*].fieldName">.spec.pins.mq[*].filters[*].properties[*].fieldName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].filters[*].properties[*].operation">.spec.pins.mq[*].filters[*].properties[*].operation</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].name">.spec.pins.mq[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name for the pin. must follow the regex pattern ^<a href="[-a-z0-9]*[a-z0-9]*[_a-z0-9]">a-z0-9</a>+$ and max length of 71 characters</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].settings">.spec.pins.mq[*].settings</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>settings for rabbitMq queue configuration</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].settings.overloadStrategy">.spec.pins.mq[*].settings.overloadStrategy</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>configuration for rabbit mq queue. default is set to “drop-head”.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].settings.queueLength">.spec.pins.mq[*].settings.queueLength</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>configuration for rabbit mq queue. default is set to 1000 msg. queueLength isn&rsquo;t used if storageOnDemand is set to true.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq[*].settings.storageOnDemand">.spec.pins.mq[*].settings.storageOnDemand</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>configuration for rabbit mq queue. default value is set to false</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus">.spec.prometheus</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>custom configuration of prometheus for microservices</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus.enabled">.spec.prometheus.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>By default this is set to true.</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus.host">.spec.prometheus.host</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>host for prometheus</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus.port">.spec.prometheus.port</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>port for prometheus</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.type">.spec.type</h3>
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
<h3 class="property-path" id="v2-.spec.versionRange">.spec.versionRange</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>when the appropriate pattern is provided, image-version will be updated with the latest tag from the image repository that satisfies that pattern.</p>

</div>

</div>
</div>

<div class="property depth-0">
<div class="property-header">
<h3 class="property-path" id="v2-.status">.status</h3>
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
<h3 class="property-path" id="v2-.status.conditions">.status.conditions</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*]">.status.conditions[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.status.conditions[*].lastTransitionTime">.status.conditions[*].lastTransitionTime</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].lastUpdateTime">.status.conditions[*].lastUpdateTime</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].message">.status.conditions[*].message</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].reason">.status.conditions[*].reason</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].status">.status.conditions[*].status</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].type">.status.conditions[*].type</h3>
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
<h3 class="property-path" id="v2-.status.message">.status.message</h3>
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
<h3 class="property-path" id="v2-.status.phase">.status.phase</h3>
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
<h3 class="property-path" id="v2-.status.subResourceName">.status.subResourceName</h3>
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



---
title: Th2Dictionary CRD schema reference (group th2.exactpro.com)
linkTitle: Th2Dictionary
description: |
  Custom resource definition (CRD) schema reference page for the Th2Dictionary resource (th2dictionaries.th2.exactpro.com), as part of the Giant Swarm Management API documentation.
weight: 100
crd:
  name_camelcase: Th2Dictionary
  name_plural: th2dictionaries
  name_singular: th2dictionary
  group: th2.exactpro.com
  technical_name: th2dictionaries.th2.exactpro.com
  scope: Namespaced
  source_repository: https://github.com/th2-net/th2-infra
  source_repository_ref: release-v2.0.0
  versions:
    - v2
  topics:
    - managementcluster
layout: crd
owner:
  - https://github.com/th2-net/th2-infra
aliases:
  - /reference/cp-k8s-api/th2dictionaries.th2.exactpro.com/
technical_name: th2dictionaries.th2.exactpro.com
source_repository: https://github.com/th2-net/th2-infra
source_repository_ref: release-v2.0.0
---

# Th2Dictionary

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



---
title: Th2Estore CRD schema reference (group th2.exactpro.com)
linkTitle: Th2Estore
description: |
  Custom resource definition (CRD) schema reference page for the Th2Estore resource (th2estores.th2.exactpro.com), as part of the Giant Swarm Management API documentation.
weight: 100
crd:
  name_camelcase: Th2Estore
  name_plural: th2estores
  name_singular: th2estore
  group: th2.exactpro.com
  technical_name: th2estores.th2.exactpro.com
  scope: Namespaced
  source_repository: https://github.com/th2-net/th2-infra
  source_repository_ref: release-v2.0.0
  versions:
    - v2
  topics:
    - managementcluster
layout: crd
owner:
  - https://github.com/th2-net/th2-infra
aliases:
  - /reference/cp-k8s-api/th2estores.th2.exactpro.com/
technical_name: th2estores.th2.exactpro.com
source_repository: https://github.com/th2-net/th2-infra
source_repository_ref: release-v2.0.0
---

# Th2Estore

<dl class="crd-meta">
<dt class="fullname">Full name:</dt>
<dd class="fullname">th2estores.th2.exactpro.com</dd>
<dt class="groupname">Group:</dt>
<dd class="groupname">th2.exactpro.com</dd>
<dt class="singularname">Singular name:</dt>
<dd class="singularname">th2estore</dd>
<dt class="pluralname">Plural name:</dt>
<dd class="pluralname">th2estores</dd>
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
<p>Specification of desired event store</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.bookName">.spec.bookName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>Can be used for passing custom book for a specific CR</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.cradleManager">.spec.cradleManager</h3>
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
<h3 class="property-path" id="v2-.spec.customConfig">.spec.customConfig</h3>
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
<h3 class="property-path" id="v2-.spec.disabled">.spec.disabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if set to true resource will act as if it is removed from schema without actually deleting the file. The default value is false.</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings">.spec.extendedSettings</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>section for extended settings</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.chartCfg">.spec.extendedSettings.chartCfg</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>specifies which charts to use for rolling out infra</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.chartCfg.path">.spec.extendedSettings.chartCfg.path</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.chartCfg.ref">.spec.extendedSettings.chartCfg.ref</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.envVariables">.spec.extendedSettings.envVariables</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>allows us to pass specific environment variables that are going to be set into the pods.</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox">.spec.extendedSettings.externalBox</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>using this section we can configure boxes that are going to be run outside of kubernetes cluster</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.address">.spec.extendedSettings.externalBox.address</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>address to the machine on which external box is running</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.enabled">.spec.extendedSettings.externalBox.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>represents the state of the external box</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints">.spec.extendedSettings.externalBox.endpoints</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>internal mapping for ports</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints[*]">.spec.extendedSettings.externalBox.endpoints[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints[*].name">.spec.extendedSettings.externalBox.endpoints[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name for the endpoint</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints[*].targetPort">.spec.extendedSettings.externalBox.endpoints[*].targetPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>value for exposed port</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases">.spec.extendedSettings.hostAliases</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*]">.spec.extendedSettings.hostAliases[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*].hostnames">.spec.extendedSettings.hostAliases[*].hostnames</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*].hostnames[*]">.spec.extendedSettings.hostAliases[*].hostnames[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*].ip">.spec.extendedSettings.hostAliases[*].ip</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.hostNetwork">.spec.extendedSettings.hostNetwork</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if the flag is set to true, pod will run on node network and kubernetes will decide which node will be used for running the box.</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.k8sProbes">.spec.extendedSettings.k8sProbes</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if enabled liveness probes will be collected from pod</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting">.spec.extendedSettings.mounting</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>allows to configure persistent volume mounting for pods</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting[*]">.spec.extendedSettings.mounting[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting[*].path">.spec.extendedSettings.mounting[*].path</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>a directory in pod where you mount external folder</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting[*].pvcName">.spec.extendedSettings.mounting[*].pvcName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name for the pvc</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.nodeSelector">.spec.extendedSettings.nodeSelector</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.replicas">.spec.extendedSettings.replicas</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>number of replicas</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources">.spec.extendedSettings.resources</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>here we can specify resource limitations and allowances for this specific component</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.limits">.spec.extendedSettings.resources.limits</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.limits.cpu">.spec.extendedSettings.resources.limits.cpu</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.limits.memory">.spec.extendedSettings.resources.limits.memory</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.requests">.spec.extendedSettings.resources.requests</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.requests.cpu">.spec.extendedSettings.resources.requests.cpu</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.requests.memory">.spec.extendedSettings.resources.requests.memory</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service">.spec.extendedSettings.service</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP">.spec.extendedSettings.service.clusterIP</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>endpoints for cluster IP service type</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP[*]">.spec.extendedSettings.service.clusterIP[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP[*].containerPort">.spec.extendedSettings.service.clusterIP[*].containerPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>allows to specify Kubernetes port for the pod.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP[*].name">.spec.extendedSettings.service.clusterIP[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name of the endpoint</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.enabled">.spec.extendedSettings.service.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if enabled allows pod to be exposed using ClusterIP</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.ingress">.spec.extendedSettings.service.ingress</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>section for ingress</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.ingress.urlPaths">.spec.extendedSettings.service.ingress.urlPaths</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>url paths</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.ingress.urlPaths[*]">.spec.extendedSettings.service.ingress.urlPaths[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort">.spec.extendedSettings.service.nodePort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>endpoints for cluster IP service type</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*]">.spec.extendedSettings.service.nodePort[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*].containerPort">.spec.extendedSettings.service.nodePort[*].containerPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>allows to specify Kubernetes port for the pod.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*].exposedPort">.spec.extendedSettings.service.nodePort[*].exposedPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>a port on which the service will be accessible.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*].name">.spec.extendedSettings.service.nodePort[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name of the endpoint</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.sharedMemory">.spec.extendedSettings.sharedMemory</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.sharedMemory.enabled">.spec.extendedSettings.sharedMemory.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.grpcRouter">.spec.grpcRouter</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>this is custom configuration for router grpc</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.imageName">.spec.imageName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>docker image repository URL</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.imageVersion">.spec.imageVersion</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>docker image tag</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.loggingConfig">.spec.loggingConfig</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>this is logging configuration</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.mqRouter">.spec.mqRouter</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>this is custom configuration for router mq</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.params">.spec.params</h3>
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
<h3 class="property-path" id="v2-.spec.params[*]">.spec.params[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.params[*].name">.spec.params[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name of parameter</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.params[*].value">.spec.params[*].value</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>value of parameter</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus">.spec.prometheus</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>custom configuration of prometheus for microservices</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus.enabled">.spec.prometheus.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>By default this is set to true.</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus.host">.spec.prometheus.host</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>host for prometheus</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus.port">.spec.prometheus.port</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>port for prometheus</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.type">.spec.type</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.versionRange">.spec.versionRange</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>when the appropriate pattern is provided, image-version will be updated with the latest tag from the image repository that satisfies that pattern.</p>

</div>

</div>
</div>

<div class="property depth-0">
<div class="property-header">
<h3 class="property-path" id="v2-.status">.status</h3>
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
<h3 class="property-path" id="v2-.status.conditions">.status.conditions</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*]">.status.conditions[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.status.conditions[*].lastTransitionTime">.status.conditions[*].lastTransitionTime</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].lastUpdateTime">.status.conditions[*].lastUpdateTime</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].message">.status.conditions[*].message</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].reason">.status.conditions[*].reason</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].status">.status.conditions[*].status</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].type">.status.conditions[*].type</h3>
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
<h3 class="property-path" id="v2-.status.message">.status.message</h3>
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
<h3 class="property-path" id="v2-.status.phase">.status.phase</h3>
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
<h3 class="property-path" id="v2-.status.subResourceName">.status.subResourceName</h3>
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
  source_repository_ref: release-v2.0.0
  versions:
    - v2
  topics:
    - managementcluster
layout: crd
owner:
  - https://github.com/th2-net/th2-infra
aliases:
  - /reference/cp-k8s-api/th2links.th2.exactpro.com/
technical_name: th2links.th2.exactpro.com
source_repository: https://github.com/th2-net/th2-infra
source_repository_ref: release-v2.0.0
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
<p>Specification of desired links. Includes three main sections - boxesRelation, dictionariesRelation and multiDictionariesRelation</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.boxesRelation">.spec.boxesRelation</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerGrpc">.spec.boxesRelation.routerGrpc</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerGrpc[*]">.spec.boxesRelation.routerGrpc[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.boxesRelation.routerGrpc[*].from">.spec.boxesRelation.routerGrpc[*].from</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerGrpc[*].from.box">.spec.boxesRelation.routerGrpc[*].from.box</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerGrpc[*].from.pin">.spec.boxesRelation.routerGrpc[*].from.pin</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerGrpc[*].name">.spec.boxesRelation.routerGrpc[*].name</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerGrpc[*].to">.spec.boxesRelation.routerGrpc[*].to</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerGrpc[*].to.box">.spec.boxesRelation.routerGrpc[*].to.box</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerGrpc[*].to.pin">.spec.boxesRelation.routerGrpc[*].to.pin</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerMq">.spec.boxesRelation.routerMq</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerMq[*]">.spec.boxesRelation.routerMq[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.boxesRelation.routerMq[*].from">.spec.boxesRelation.routerMq[*].from</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerMq[*].from.box">.spec.boxesRelation.routerMq[*].from.box</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerMq[*].from.pin">.spec.boxesRelation.routerMq[*].from.pin</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerMq[*].name">.spec.boxesRelation.routerMq[*].name</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerMq[*].to">.spec.boxesRelation.routerMq[*].to</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerMq[*].to.box">.spec.boxesRelation.routerMq[*].to.box</h3>
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
<h3 class="property-path" id="v2-.spec.boxesRelation.routerMq[*].to.pin">.spec.boxesRelation.routerMq[*].to.pin</h3>
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
<h3 class="property-path" id="v2-.spec.dictionariesRelation">.spec.dictionariesRelation</h3>
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
<h3 class="property-path" id="v2-.spec.dictionariesRelation[*]">.spec.dictionariesRelation[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.dictionariesRelation[*].box">.spec.dictionariesRelation[*].box</h3>
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
<h3 class="property-path" id="v2-.spec.dictionariesRelation[*].dictionary">.spec.dictionariesRelation[*].dictionary</h3>
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
<h3 class="property-path" id="v2-.spec.dictionariesRelation[*].dictionary.name">.spec.dictionariesRelation[*].dictionary.name</h3>
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
<h3 class="property-path" id="v2-.spec.dictionariesRelation[*].dictionary.type">.spec.dictionariesRelation[*].dictionary.type</h3>
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
<h3 class="property-path" id="v2-.spec.dictionariesRelation[*].name">.spec.dictionariesRelation[*].name</h3>
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
<h3 class="property-path" id="v2-.spec.multiDictionariesRelation">.spec.multiDictionariesRelation</h3>
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
<h3 class="property-path" id="v2-.spec.multiDictionariesRelation[*]">.spec.multiDictionariesRelation[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.multiDictionariesRelation[*].box">.spec.multiDictionariesRelation[*].box</h3>
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
<h3 class="property-path" id="v2-.spec.multiDictionariesRelation[*].dictionaries">.spec.multiDictionariesRelation[*].dictionaries</h3>
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
<h3 class="property-path" id="v2-.spec.multiDictionariesRelation[*].dictionaries[*]">.spec.multiDictionariesRelation[*].dictionaries[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.multiDictionariesRelation[*].dictionaries[*].alias">.spec.multiDictionariesRelation[*].dictionaries[*].alias</h3>
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
<h3 class="property-path" id="v2-.spec.multiDictionariesRelation[*].dictionaries[*].name">.spec.multiDictionariesRelation[*].dictionaries[*].name</h3>
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
<h3 class="property-path" id="v2-.spec.multiDictionariesRelation[*].name">.spec.multiDictionariesRelation[*].name</h3>
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



---
title: Th2Mstore CRD schema reference (group th2.exactpro.com)
linkTitle: Th2Mstore
description: |
  Custom resource definition (CRD) schema reference page for the Th2Mstore resource (th2mstores.th2.exactpro.com), as part of the Giant Swarm Management API documentation.
weight: 100
crd:
  name_camelcase: Th2Mstore
  name_plural: th2mstores
  name_singular: th2mstore
  group: th2.exactpro.com
  technical_name: th2mstores.th2.exactpro.com
  scope: Namespaced
  source_repository: https://github.com/th2-net/th2-infra
  source_repository_ref: release-v2.0.0
  versions:
    - v2
  topics:
    - managementcluster
layout: crd
owner:
  - https://github.com/th2-net/th2-infra
aliases:
  - /reference/cp-k8s-api/th2mstores.th2.exactpro.com/
technical_name: th2mstores.th2.exactpro.com
source_repository: https://github.com/th2-net/th2-infra
source_repository_ref: release-v2.0.0
---

# Th2Mstore

<dl class="crd-meta">
<dt class="fullname">Full name:</dt>
<dd class="fullname">th2mstores.th2.exactpro.com</dd>
<dt class="groupname">Group:</dt>
<dd class="groupname">th2.exactpro.com</dd>
<dt class="singularname">Singular name:</dt>
<dd class="singularname">th2mstore</dd>
<dt class="pluralname">Plural name:</dt>
<dd class="pluralname">th2mstores</dd>
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
<p>Specification of desired message store</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.bookName">.spec.bookName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>Can be used for passing custom book for a specific CR</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.cradleManager">.spec.cradleManager</h3>
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
<h3 class="property-path" id="v2-.spec.customConfig">.spec.customConfig</h3>
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
<h3 class="property-path" id="v2-.spec.disabled">.spec.disabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if set to true resource will act as if it is removed from schema without actually deleting the file. The default value is false.</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings">.spec.extendedSettings</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>section for extended settings</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.chartCfg">.spec.extendedSettings.chartCfg</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>specifies which charts to use for rolling out infra</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.chartCfg.path">.spec.extendedSettings.chartCfg.path</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.chartCfg.ref">.spec.extendedSettings.chartCfg.ref</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.envVariables">.spec.extendedSettings.envVariables</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>allows us to pass specific environment variables that are going to be set into the pods.</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox">.spec.extendedSettings.externalBox</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>using this section we can configure boxes that are going to be run outside of kubernetes cluster</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.address">.spec.extendedSettings.externalBox.address</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>address to the machine on which external box is running</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.enabled">.spec.extendedSettings.externalBox.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>represents the state of the external box</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints">.spec.extendedSettings.externalBox.endpoints</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>internal mapping for ports</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints[*]">.spec.extendedSettings.externalBox.endpoints[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints[*].name">.spec.extendedSettings.externalBox.endpoints[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name for the endpoint</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.externalBox.endpoints[*].targetPort">.spec.extendedSettings.externalBox.endpoints[*].targetPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>value for exposed port</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases">.spec.extendedSettings.hostAliases</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*]">.spec.extendedSettings.hostAliases[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*].hostnames">.spec.extendedSettings.hostAliases[*].hostnames</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*].hostnames[*]">.spec.extendedSettings.hostAliases[*].hostnames[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.hostAliases[*].ip">.spec.extendedSettings.hostAliases[*].ip</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.hostNetwork">.spec.extendedSettings.hostNetwork</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if the flag is set to true, pod will run on node network and kubernetes will decide which node will be used for running the box.</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.k8sProbes">.spec.extendedSettings.k8sProbes</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if enabled liveness probes will be collected from pod</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting">.spec.extendedSettings.mounting</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>allows to configure persistent volume mounting for pods</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting[*]">.spec.extendedSettings.mounting[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting[*].path">.spec.extendedSettings.mounting[*].path</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>a directory in pod where you mount external folder</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.mounting[*].pvcName">.spec.extendedSettings.mounting[*].pvcName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name for the pvc</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.nodeSelector">.spec.extendedSettings.nodeSelector</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.replicas">.spec.extendedSettings.replicas</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>number of replicas</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources">.spec.extendedSettings.resources</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>here we can specify resource limitations and allowances for this specific component</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.limits">.spec.extendedSettings.resources.limits</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.limits.cpu">.spec.extendedSettings.resources.limits.cpu</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.limits.memory">.spec.extendedSettings.resources.limits.memory</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.requests">.spec.extendedSettings.resources.requests</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.requests.cpu">.spec.extendedSettings.resources.requests.cpu</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.resources.requests.memory">.spec.extendedSettings.resources.requests.memory</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service">.spec.extendedSettings.service</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP">.spec.extendedSettings.service.clusterIP</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>endpoints for cluster IP service type</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP[*]">.spec.extendedSettings.service.clusterIP[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP[*].containerPort">.spec.extendedSettings.service.clusterIP[*].containerPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>allows to specify Kubernetes port for the pod.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.clusterIP[*].name">.spec.extendedSettings.service.clusterIP[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name of the endpoint</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.enabled">.spec.extendedSettings.service.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>if enabled allows pod to be exposed using ClusterIP</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.ingress">.spec.extendedSettings.service.ingress</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>section for ingress</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.ingress.urlPaths">.spec.extendedSettings.service.ingress.urlPaths</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>url paths</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.ingress.urlPaths[*]">.spec.extendedSettings.service.ingress.urlPaths[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort">.spec.extendedSettings.service.nodePort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>endpoints for cluster IP service type</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*]">.spec.extendedSettings.service.nodePort[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*].containerPort">.spec.extendedSettings.service.nodePort[*].containerPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>allows to specify Kubernetes port for the pod.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*].exposedPort">.spec.extendedSettings.service.nodePort[*].exposedPort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>a port on which the service will be accessible.</p>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort[*].name">.spec.extendedSettings.service.nodePort[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>name of the endpoint</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.sharedMemory">.spec.extendedSettings.sharedMemory</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.sharedMemory.enabled">.spec.extendedSettings.sharedMemory.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>DESCRIPTION NEEDED</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.grpcRouter">.spec.grpcRouter</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>this is custom configuration for router grpc</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.imageName">.spec.imageName</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>docker image repository URL</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.imageVersion">.spec.imageVersion</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>docker image tag</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.loggingConfig">.spec.loggingConfig</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>this is logging configuration</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.mqRouter">.spec.mqRouter</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>this is custom configuration for router mq</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.params">.spec.params</h3>
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
<h3 class="property-path" id="v2-.spec.params[*]">.spec.params[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.params[*].name">.spec.params[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>name of parameter</p>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.params[*].value">.spec.params[*].value</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>
<span class="property-required">Required</span>
</div>

<div class="property-description">
<p>value of parameter</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus">.spec.prometheus</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

<div class="property-description">
<p>custom configuration of prometheus for microservices</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus.enabled">.spec.prometheus.enabled</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>

</div>

<div class="property-description">
<p>By default this is set to true.</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus.host">.spec.prometheus.host</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>host for prometheus</p>

</div>

</div>
</div>

<div class="property depth-2">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.prometheus.port">.spec.prometheus.port</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>

</div>

<div class="property-description">
<p>port for prometheus</p>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.type">.spec.type</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-1">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.versionRange">.spec.versionRange</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

<div class="property-description">
<p>when the appropriate pattern is provided, image-version will be updated with the latest tag from the image repository that satisfies that pattern.</p>

</div>

</div>
</div>

<div class="property depth-0">
<div class="property-header">
<h3 class="property-path" id="v2-.status">.status</h3>
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
<h3 class="property-path" id="v2-.status.conditions">.status.conditions</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*]">.status.conditions[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.status.conditions[*].lastTransitionTime">.status.conditions[*].lastTransitionTime</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].lastUpdateTime">.status.conditions[*].lastUpdateTime</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].message">.status.conditions[*].message</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].reason">.status.conditions[*].reason</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].status">.status.conditions[*].status</h3>
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
<h3 class="property-path" id="v2-.status.conditions[*].type">.status.conditions[*].type</h3>
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
<h3 class="property-path" id="v2-.status.message">.status.message</h3>
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
<h3 class="property-path" id="v2-.status.phase">.status.phase</h3>
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
<h3 class="property-path" id="v2-.status.subResourceName">.status.subResourceName</h3>
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



