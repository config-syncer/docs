---
title: Synchronize Configuration from an OCM Hub to Managed Clusters
description: Synchronize ConfigMaps and Secrets from an Open Cluster Management hub to its managed clusters
menu:
  docs_{{ .version }}:
    identifier: ocm-addon-syncer
    name: OCM Hub to Spokes
    parent: config-syncer
    weight: 20
product_name: kubed
menu_name: docs_{{ .version }}
section_menu_id: guides
---

> New to Config Syncer? Please start [here](/docs/concepts/README.md).

# Synchronize Configuration from an OCM Hub to Managed Clusters

If you manage a fleet of clusters with [Open Cluster Management (OCM)](https://open-cluster-management.io/), Config Syncer can run as an OCM addon. You annotate a ConfigMap or Secret on the **hub** cluster with the names of the target managed clusters, and Config Syncer copies it into those clusters.

Unlike [syncing across clusters with `sync-contexts`](/docs/guides/config-syncer/inter-cluster.md), you don't need a `kubeconfig` with credentials for every target cluster. Each managed cluster runs its own agent, which pulls the objects meant for it from the hub.

## Terms

- **Hub**: the OCM management cluster.
- **Spoke** (or **ManagedCluster**): a cluster registered to the hub.
- **Cluster namespace**: on the hub, OCM creates one namespace per spoke, named after the spoke (e.g. `spoke-1`).
- **ClusterManagementAddOn**: the hub-wide definition of an addon. This addon is named `config-syncer`.
- **ManagedClusterAddOn**: "the addon is enabled on this spoke". It lives in the spoke's cluster namespace on the hub.

## How it works

The addon has two components:

| Component | Runs on | What it does |
|---|---|---|
| Manager | hub, namespace `open-cluster-management-addon` | Deploys the agent to every selected spoke. Copies each annotated source into the cluster namespace of every target spoke (a **mirror**). |
| Agent | each spoke, namespace `kubeops` | A regular Config Syncer that also watches its own cluster namespace on the hub and writes the mirrors into the spoke's namespaces. |

```
                 HUB                                            SPOKE spoke-1
 ┌─────────────────────────────────────────┐        ┌──────────────────────────┐
 │ ns demo                                 │        │ ns kubeops               │
 │   cm/omni                               │        │   deploy/config-syncer   │
 │     sync-to-clusters: spoke-1           │        │     (agent) ────────┐    │
 │        │                                │        │                     │    │
 │        │ manager                        │  watch │                     ▼    │
 │        ▼                                │◀───────┤ ns demo                  │
 │ ns spoke-1  (cluster namespace)         │ (read- │   cm/omni  (copy)        │
 │   cm/demo.omni   (mirror)               │  only) │                          │
 └─────────────────────────────────────────┘        └──────────────────────────┘
```

For example, a ConfigMap `omni` in namespace `demo` on the hub, annotated with `kubed.appscode.com/sync-to-clusters: spoke-1`:

1. The manager creates the mirror `demo.omni` in the hub namespace `spoke-1`. The mirror is named `<source namespace>.<source name>`, so two sources with the same name in different namespaces don't collide.
2. The agent on `spoke-1` sees the mirror and creates the ConfigMap `omni` in the chosen namespaces of `spoke-1`.
3. When the source changes, the mirror and the spoke copies follow. When the source is deleted, or `spoke-1` is removed from the annotation, the mirror and the copies are deleted.

The agent can read only its own cluster namespace on the hub, so a spoke can't read objects meant for another spoke.

## Annotations

Set these on the source ConfigMap/Secret on the hub:

| Annotation | Required | Meaning |
|---|---|---|
| `kubed.appscode.com/sync-to-clusters` | yes | Comma-separated ManagedCluster names, e.g. `spoke-1,spoke-2`. |
| `kubed.appscode.com/sync-to-spoke-namespace` | no | Comma-separated namespace names on the spoke, e.g. `team-a,team-b`. |
| `kubed.appscode.com/sync` | no | Label selector for namespaces on the spoke, e.g. `app=kubed`. Empty value means all namespaces. Same syntax as [intra-cluster sync](/docs/guides/config-syncer/intra-cluster.md#namespace-selector). |

Spoke target namespaces are chosen in this order:

1. `sync-to-spoke-namespace` is set: those namespaces. `sync` is ignored.
2. Otherwise, `sync` is set: every spoke namespace that matches the selector.
3. Otherwise: the namespace of the source (`demo` in the example above).

Config Syncer does not create namespaces on the spoke. A copy appears in a target namespace once that namespace exists. The agent also re-checks all mirrors when a namespace is created or relabelled on the spoke, so a new namespace matching a `sync` selector gets the copy.

> If a Config Syncer also runs on the hub itself, a `kubed.appscode.com/sync` annotation on the source also syncs it across namespaces of the hub. Prefer `sync-to-spoke-namespace` if you don't want that.

## Before You Begin

- An OCM hub with at least one registered spoke. See the [OCM quick start](https://open-cluster-management.io/docs/getting-started/quick-start/).
- `kubectl` contexts for the hub and the spoke. This guide uses `hub` and `spoke-1`, and the ManagedCluster is also named `spoke-1`.
- The agent runs Config Syncer in enterprise mode, and gets its license from the [License Proxy Server](https://github.com/appscode/license-proxyserver) on each spoke. The addon itself does not ship a license to the spokes.
- No other Config Syncer installed in the `kubeops` namespace of the spokes. The agent's Deployment and ClusterRole are both named `config-syncer` and would collide with it.

```bash
$ kubectl --context hub get managedclusters
NAME      HUB ACCEPTED   MANAGED CLUSTER URLS      JOINED   AVAILABLE   AGE
spoke-1   true           https://10.2.0.5:6443     True     True        2d
```

## Install the Addon Manager

Install the `config-syncer-manager` chart on the hub:

```bash
$ helm upgrade -i config-syncer-manager \
  oci://ghcr.io/appscode-charts/config-syncer-manager \
  --version {{< param "info.version" >}} \
  --kube-context hub \
  --namespace open-cluster-management-addon --create-namespace \
  --wait --burst-limit=10000 --debug
```

By default the addon is enabled on every cluster in the `global` ManagedClusterSet, through a Placement named `global` in the release namespace. If that Placement already exists (for example, created by another addon such as the License Proxy Server manager), add `--set placement.create=false` to reuse it.

Useful chart values:

| Value | Default | Meaning |
|---|---|---|
| `placement.create` | `true` | Create the `global` Placement and ManagedClusterSetBinding. |
| `placement.name` | `global` | Placement that selects the spokes which get the addon. |
| `clusterName` | `hub` | Hub name, recorded in the `kubed.appscode.com/origin.cluster` label of mirrors and copies. |
| `registryFQDN` | `ghcr.io` | Registry the spokes pull the agent image from. |

To see all options, visit the [chart](https://github.com/kubeops/installer/tree/{{< param "info.version" >}}/charts/config-syncer-manager).

Check that the addon is available on the spoke:

```bash
$ kubectl --context hub get managedclusteraddon config-syncer -n spoke-1
NAME            AVAILABLE   DEGRADED   PROGRESSING   AGE
config-syncer   True                   False         1m

$ kubectl --context spoke-1 get deploy -n kubeops config-syncer
NAME            READY   UP-TO-DATE   AVAILABLE   AGE
config-syncer   1/1     1            1           1m
```

## Synchronize a ConfigMap

Create a ConfigMap on the hub:

```bash
$ kubectl --context hub create namespace demo
namespace/demo created

$ kubectl --context hub create configmap omni -n demo \
  --from-literal=you=only \
  --from-literal=live=once
configmap/omni created
```

Annotate it with the target cluster:

```bash
$ kubectl --context hub annotate configmap omni -n demo \
  kubed.appscode.com/sync-to-clusters=spoke-1
configmap/omni annotated
```

The manager creates the mirror in the cluster namespace `spoke-1` on the hub. No namespace annotation is set on the source, so the manager records the source namespace `demo` as the target:

```bash
$ kubectl --context hub get configmap demo.omni -n spoke-1 -o yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: demo.omni
  namespace: spoke-1
  labels:
    kubed.appscode.com/hub-mirror: "true"
    kubed.appscode.com/origin.cluster: hub
    kubed.appscode.com/origin.hash: "..."
  annotations:
    kubed.appscode.com/origin.name: omni
    kubed.appscode.com/origin.namespace: demo
    kubed.appscode.com/sync-to-spoke-namespace: demo
data:
  live: once
  you: only
```

Namespace `demo` doesn't exist on the spoke yet, so there's no copy. Create it:

```bash
$ kubectl --context spoke-1 create namespace demo
namespace/demo created

$ kubectl --context spoke-1 get configmap omni -n demo
NAME   DATA   AGE
omni   2      5s
```

The copy keeps the source's name and carries a `kubed.appscode.com/hub-mirror-ref` label, which the agent uses to find and clean up its copies.

If you edit the source on the hub, the copy on the spoke is updated:

```bash
$ kubectl --context hub patch configmap omni -n demo --type merge -p '{"data":{"live":"twice"}}'
configmap/omni patched

$ kubectl --context spoke-1 get configmap omni -n demo -o jsonpath='{.data.live}'
twice
```

## Select Spoke Namespaces by Label

Use `kubed.appscode.com/sync` to copy into every spoke namespace that matches a label selector:

```bash
$ kubectl --context spoke-1 create namespace team-a
$ kubectl --context spoke-1 label namespace team-a app=kubed

$ kubectl --context hub annotate configmap omni -n demo kubed.appscode.com/sync=app=kubed
configmap/omni annotated

$ kubectl --context spoke-1 get configmaps -A | grep omni
team-a        omni     2      4s
```

The copy in `demo` is gone, because the target is now the selector, not the source namespace. On the mirror, the manager stores the selector as `kubed.appscode.com/hub-sync` instead of `kubed.appscode.com/sync`, so a Config Syncer running on the hub doesn't copy the mirror across hub namespaces.

## Select Spoke Namespaces by Name

Use `kubed.appscode.com/sync-to-spoke-namespace` to name the target namespaces. It takes priority over `kubed.appscode.com/sync`:

```bash
$ kubectl --context spoke-1 create namespace team-b

$ kubectl --context hub annotate configmap omni -n demo \
  kubed.appscode.com/sync-to-spoke-namespace=demo,team-b
configmap/omni annotated

$ kubectl --context spoke-1 get configmaps -A | grep omni
demo          omni     2      3s
team-b        omni     2      3s
```

`team-a` lost its copy, because the named list wins over the selector.

## Synchronize a Secret

Secrets work the same way:

```bash
$ kubectl --context hub create secret generic db-pass -n demo --from-literal=password=s3cr3t
secret/db-pass created

$ kubectl --context hub annotate secret db-pass -n demo \
  kubed.appscode.com/sync-to-clusters=spoke-1
secret/db-pass annotated

$ kubectl --context spoke-1 get secret db-pass -n demo -o jsonpath='{.data.password}' | base64 -d
s3cr3t
```

## Remove the Copies

Remove the spoke from `sync-to-clusters`, remove the annotation, or delete the source. The mirror on the hub is deleted, and the agent deletes its copies on the spoke:

```bash
$ kubectl --context hub annotate configmap omni -n demo kubed.appscode.com/sync-to-clusters-
configmap/omni annotated

$ kubectl --context hub get configmap -n spoke-1 demo.omni
Error from server (NotFound): configmaps "demo.omni" not found

$ kubectl --context spoke-1 get configmaps -A | grep omni
```

## Labels and Annotations Added by the Addon

| Key | On | Meaning |
|---|---|---|
| `kubed.appscode.com/hub-mirror: "true"` (label) | hub mirror | Marks a mirror. The agent only watches objects with this label. Mirrors are never treated as sources. |
| `kubed.appscode.com/origin.cluster` (label) | hub mirror, spoke copy | Hub name, from the chart value `clusterName`. |
| `kubed.appscode.com/origin.hash` (label) | hub mirror, spoke copy | Hash of the hub name and the source namespace/name. Used to find all mirrors of one source. |
| `kubed.appscode.com/origin.namespace`, `kubed.appscode.com/origin.name` (annotations) | hub mirror, spoke copy | Namespace and name of the source on the hub. The spoke copy uses the same name. |
| `kubed.appscode.com/hub-sync` (annotation) | hub mirror | The source's `kubed.appscode.com/sync` value. |
| `kubed.appscode.com/hub-mirror-ref` (label) | spoke copy | Hash of `<source namespace>.<source name>`. Links the copy to its mirror. |

## Behavior and Limitations

| Situation | Result |
|---|---|
| A cluster in `sync-to-clusters` doesn't have the addon enabled | Skipped and logged. Mirrored once the addon is enabled. |
| The spoke is offline or the agent is down | Mirrors stay on the hub. The agent catches up when it restarts, and removes copies whose mirror is gone. |
| The manager is down | No new mirrors are created. Agents keep serving the existing ones. |
| The spoke already has an object with the same name in a target namespace | It is overwritten. |
| The addon is removed from a spoke | The mirrors in that cluster namespace are deleted. The spoke copies are deleted only if the agent processes those deletions before it is removed; otherwise they are left on the spoke. |
| The source is deleted while the manager is down | The mirrors and spoke copies may be left behind. |

Other limitations:

- Only ConfigMaps and Secrets are synchronized.
- Target clusters are selected by name only, not by label or Placement.
- The agent's hub identity can read all ConfigMaps and Secrets in its own cluster namespace on the hub, not only the mirrors. Kubernetes RBAC can't restrict access by label. It can't read other cluster namespaces.

To check the agent's access on the hub, find its identity in the `status.registrations` field of the spoke's `config-syncer` ManagedClusterAddOn. It depends on the OCM registration driver; the example below uses the token driver's ServiceAccount:

```bash
$ kubectl --context hub auth can-i list secrets -n spoke-1 \
  --as=system:serviceaccount:spoke-1:config-syncer-agent
yes

$ kubectl --context hub auth can-i list secrets -n spoke-2 \
  --as=system:serviceaccount:spoke-1:config-syncer-agent
no
```

## Cleaning up

```bash
$ kubectl --context hub delete namespace demo
$ kubectl --context spoke-1 delete namespace demo team-a team-b
$ helm uninstall config-syncer-manager --kube-context hub -n open-cluster-management-addon
```

## Next Steps

- Learn how to sync ConfigMaps/Secrets across namespaces of a cluster [here](/docs/guides/config-syncer/intra-cluster.md).
- Learn how to sync ConfigMaps/Secrets across clusters with a `kubeconfig` [here](/docs/guides/config-syncer/inter-cluster.md).
- Want to hack on Config Syncer? Check our [contribution guidelines](/docs/CONTRIBUTING.md).
