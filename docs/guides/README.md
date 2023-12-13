---
title: Overview
description: Overview of guides
menu:
  docs_{{ .version }}:
    identifier: guides-overview
    name: Overview
    parent: guides
    weight: -1
product_name: kubed
menu_name: docs_{{ .version }}
section_menu_id: guides
url: /docs/{{ .version }}/guides/
aliases:
  - /docs/{{ .version }}/guides/README/
---

# Guides

This section contains guides on how to use Config Syncer. Please visit the links below to learn more:

- Configuration Syncer
  - [Synchronize Configuration across Namespaces](/docs/guides/config-syncer/intra-cluster.md): This tutorial will show you how Config Syncer can sync ConfigMaps/Secrets across Kubernetes namespaces.
  - [Synchronize Configuration across Clusters](/docs/guides/config-syncer/inter-cluster.md): This tutorial will show you how Config Syncer can sync ConfigMaps/Secrets across Kubernetes cluster.
- [Monitoring Config Syncer](/docs/guides/monitoring.md): This article described the various metrics exposed by Config Syncer operator.
