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
  source_repository_ref: release-v2.0.x-repo-update
  versions:
    - v2
layout: crd
owner:
  - https://github.com/th2-net/th2-infra
aliases:
  - /reference/cp-k8s-api/th2coreboxes.th2.exactpro.com/
technical_name: th2coreboxes.th2.exactpro.com
source_repository: https://github.com/th2-net/th2-infra
source_repository_ref: release-v2.0.x-repo-update
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
<p>NodeSelector is a selector which must be true for the pod to fit on a node.
Selector which must match a node&rsquo;s labels for the pod to be scheduled on that node. More info: <a href="https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector">https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector</a></p>

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
<p>Creates Ingress for the Box</p>

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
<p>Path is a suffix, it is appended to namespace url, e.g for th2-test namespace:</p>

<p>urlPaths:
- /my-box/</p>

<p>service will be available at /th2-test/my-box/</p>

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
<h3 class="property-path" id="v2-.spec.extendedSettings.service.loadBalancer">.spec.extendedSettings.service.loadBalancer</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>endpoints for loadBalancer service type</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.loadBalancer[*]">.spec.extendedSettings.service.loadBalancer[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.extendedSettings.service.loadBalancer[*].containerPort">.spec.extendedSettings.service.loadBalancer[*].containerPort</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.service.loadBalancer[*].name">.spec.extendedSettings.service.loadBalancer[*].name</h3>
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
<h3 class="property-path" id="v2-.spec.extendedSettings.service.nodePort">.spec.extendedSettings.service.nodePort</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>endpoints for nodePort service type</p>

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
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].linkTo">.spec.pins.grpc.client[*].linkTo</h3>
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

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].linkTo[*]">.spec.pins.grpc.client[*].linkTo[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].linkTo[*].box">.spec.pins.grpc.client[*].linkTo[*].box</h3>
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

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.grpc.client[*].linkTo[*].pin">.spec.pins.grpc.client[*].linkTo[*].pin</h3>
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
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.publishers">.spec.pins.mq.publishers</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>section for subscriber mq pins</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*]">.spec.pins.mq.publishers[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].attributes">.spec.pins.mq.publishers[*].attributes</h3>
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

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].attributes[*]">.spec.pins.mq.publishers[*].attributes[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters">.spec.pins.mq.publishers[*].filters</h3>
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

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*]">.spec.pins.mq.publishers[*].filters[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*].message">.spec.pins.mq.publishers[*].filters[*].message</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*].message[*]">.spec.pins.mq.publishers[*].filters[*].message[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-9">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*].message[*].expectedValue">.spec.pins.mq.publishers[*].filters[*].message[*].expectedValue</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*].message[*].fieldName">.spec.pins.mq.publishers[*].filters[*].message[*].fieldName</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*].message[*].operation">.spec.pins.mq.publishers[*].filters[*].message[*].operation</h3>
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

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*].metadata">.spec.pins.mq.publishers[*].filters[*].metadata</h3>
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

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*].metadata[*]">.spec.pins.mq.publishers[*].filters[*].metadata[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-9">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*].metadata[*].expectedValue">.spec.pins.mq.publishers[*].filters[*].metadata[*].expectedValue</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*].metadata[*].fieldName">.spec.pins.mq.publishers[*].filters[*].metadata[*].fieldName</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*].metadata[*].operation">.spec.pins.mq.publishers[*].filters[*].metadata[*].operation</h3>
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

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*].properties">.spec.pins.mq.publishers[*].filters[*].properties</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*].properties[*]">.spec.pins.mq.publishers[*].filters[*].properties[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-9">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*].properties[*].expectedValue">.spec.pins.mq.publishers[*].filters[*].properties[*].expectedValue</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*].properties[*].fieldName">.spec.pins.mq.publishers[*].filters[*].properties[*].fieldName</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].filters[*].properties[*].operation">.spec.pins.mq.publishers[*].filters[*].properties[*].operation</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.publishers[*].name">.spec.pins.mq.publishers[*].name</h3>
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

<div class="property depth-3">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers">.spec.pins.mq.subscribers</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

</div>

<div class="property-description">
<p>section for subscriber mq pins</p>

</div>

</div>
</div>

<div class="property depth-4">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*]">.spec.pins.mq.subscribers[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].attributes">.spec.pins.mq.subscribers[*].attributes</h3>
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

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].attributes[*]">.spec.pins.mq.subscribers[*].attributes[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

</div>

</div>
</div>

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters">.spec.pins.mq.subscribers[*].filters</h3>
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

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*]">.spec.pins.mq.subscribers[*].filters[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*].message">.spec.pins.mq.subscribers[*].filters[*].message</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*].message[*]">.spec.pins.mq.subscribers[*].filters[*].message[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-9">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*].message[*].expectedValue">.spec.pins.mq.subscribers[*].filters[*].message[*].expectedValue</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*].message[*].fieldName">.spec.pins.mq.subscribers[*].filters[*].message[*].fieldName</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*].message[*].operation">.spec.pins.mq.subscribers[*].filters[*].message[*].operation</h3>
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

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*].metadata">.spec.pins.mq.subscribers[*].filters[*].metadata</h3>
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

<div class="property depth-8">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*].metadata[*]">.spec.pins.mq.subscribers[*].filters[*].metadata[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-9">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*].metadata[*].expectedValue">.spec.pins.mq.subscribers[*].filters[*].metadata[*].expectedValue</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*].metadata[*].fieldName">.spec.pins.mq.subscribers[*].filters[*].metadata[*].fieldName</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*].metadata[*].operation">.spec.pins.mq.subscribers[*].filters[*].metadata[*].operation</h3>
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

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*].properties">.spec.pins.mq.subscribers[*].filters[*].properties</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*].properties[*]">.spec.pins.mq.subscribers[*].filters[*].properties[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-9">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*].properties[*].expectedValue">.spec.pins.mq.subscribers[*].filters[*].properties[*].expectedValue</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*].properties[*].fieldName">.spec.pins.mq.subscribers[*].filters[*].properties[*].fieldName</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].filters[*].properties[*].operation">.spec.pins.mq.subscribers[*].filters[*].properties[*].operation</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].linkTo">.spec.pins.mq.subscribers[*].linkTo</h3>
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

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].linkTo[*]">.spec.pins.mq.subscribers[*].linkTo[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

</div>

</div>
</div>

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].linkTo[*].box">.spec.pins.mq.subscribers[*].linkTo[*].box</h3>
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

<div class="property depth-7">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].linkTo[*].pin">.spec.pins.mq.subscribers[*].linkTo[*].pin</h3>
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

<div class="property depth-5">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].name">.spec.pins.mq.subscribers[*].name</h3>
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
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].settings">.spec.pins.mq.subscribers[*].settings</h3>
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

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].settings.overloadStrategy">.spec.pins.mq.subscribers[*].settings.overloadStrategy</h3>
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

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].settings.queueLength">.spec.pins.mq.subscribers[*].settings.queueLength</h3>
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

<div class="property depth-6">
<div class="property-header">
<h3 class="property-path" id="v2-.spec.pins.mq.subscribers[*].settings.storageOnDemand">.spec.pins.mq.subscribers[*].settings.storageOnDemand</h3>
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



