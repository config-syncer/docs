---
title: Configuration Syncer | Kubernetes
description: Configuration Syncer for Kubernetes Clusters
menu:
  docs_{{ .version }}:
    identifier: readme-config-syncer
    name: Overview
    parent: config-syncer
    weight: -1
product_name: kubed
menu_name: docs_{{ .version }}
section_menu_id: guides
url: /docs/{{ .version }}/guides/config-syncer/
aliases:
  - /docs/{{ .version }}/guides/config-syncer/README
---

# Configuration Syncer

This section contains guides on how to use Config Syncer to synchronize Configmaps or Secrets across namespaces of a Kubernetes Cluster or across Kubernetes clusters.

- [Synchronize Configuration across Namespaces](/docs/guides/config-syncer/intra-cluster.md): This tutorial will show you how Config Syncer can sync ConfigMaps/Secrets across Kubernetes namespaces.
- [Synchronize Configuration across Clusters](/docs/guides/config-syncer/inter-cluster.md): This tutorial will show you how Config Syncer can sync ConfigMaps/Secrets across Kubernetes cluster.
