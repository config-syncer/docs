---
title: Overview | Kubed
description: Kubed Overview
menu:
  product_kubed_0.3.0:
    identifier: kubed-overview
    name: Overview
    parent: getting-started
    weight: 20
product_name: kubed
menu_name: product_kubed_0.3.0
section_menu_id: getting-started
url: /products/kubed/0.3.0/getting-started/
aliases:
  - /products/kubed/0.3.0/
  - /products/kubed/0.3.0/README/
---

[![Go Report Card](https://goreportcard.com/badge/github.com/appscode/kubed)](https://goreportcard.com/report/github.com/appscode/kubed)

# Kubed
Kubed (pronounced Cube-Dee) by AppsCode is a Kubernetes Cluster Operator Daemon. Kubed can do the following things for you:

 - Kubed can take periodic [snapshot](/docs/tutorials/cluster-snapshot.md) of a Kubernetes cluster objects.
 - Kubed provides a [recycle bin](/docs/tutorials/recycle-bin.md) for deleted and/or updated Kubernetes objects.
 - Kubed can keep [ConfigMaps and Secrets synchronized across Namespaces](/docs/tutorials/config-syncer.md).
 - Kubed can [forward cluster events](/docs/tutorials/event-forwarder.md) to various destinations.
 - Kubed can setup [janitors](/docs/tutorials/janitors.md) for Elasticsearch and InfluxDB.
 - Kubed can [send notifications](/docs/tutorials/notifiers.md) via Email, SMS or Chat.
 - Kubed has a built-in [search engine](/docs/tutorials/apiserver.md) for your cluster objects using [bleve](https://github.com/blevesearch/bleve).

## Supported Versions
Please pick a version of Kubed that matches your Kubernetes installation.

| Kubed Version                                                 | Docs                                                            | Kubernetes Version |
|---------------------------------------------------------------|-----------------------------------------------------------------|--------------------|
| [0.3.0](https://github.com/appscode/kubed/releases/tag/0.3.0) | [User Guide](https://github.com/appscode/kubed/tree/0.3.0/docs) | 1.7.x+             |
| [0.2.0](https://github.com/appscode/kubed/releases/tag/0.2.0) | [User Guide](https://github.com/appscode/kubed/tree/0.2.0/docs) | 1.5.x - 1.7.x      |

## Installation
To install Kubed, please follow the guide [here](/docs/install.md).

## Using Kubed
Want to learn how to use Kubed? Please start [here](/docs/tutorials/README.md).

## Contribution guidelines
Want to help improve Kubed? Please start [here](/CONTRIBUTING.md).

## Project Status
Wondering what features are coming next? Please visit [here](/ROADMAP.md).

---

**The kubed operator collects anonymous usage statistics to help us learn how the software is being used and how we can improve it. To disable stats collection, run the operator with the flag** `--analytics=false`.

---

## Support
If you have any questions, you can reach out to us.
* [Slack](https://slack.appscode.com)
* [Twitter](https://twitter.com/AppsCodeHQ)
* [Website](https://appscode.com)
