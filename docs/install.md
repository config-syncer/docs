> New to Kubed? Please start [here](/docs/tutorials/README.md).

# Installation Guide

## Using YAML
Kubed can be installed using YAML files includes in the [/hack/deploy](/hack/deploy) folder.

```console
# Install without RBAC roles
$ curl https://raw.githubusercontent.com/appscode/kubed/0.1.0/hack/deploy/without-rbac.yaml \
  | kubectl apply -f -


# Install with RBAC roles
$ curl https://raw.githubusercontent.com/appscode/kubed/0.1.0/hack/deploy/with-rbac.yaml \
  | kubectl apply -f -
```

## Using Helm
Kubed can be installed via [Helm](https://helm.sh/) using the [chart](/chart/kubed) included in this repository. To install the chart with the release name `my-release`:
```bash
$ helm install chart/kubed --name my-release
```
To see the detailed configuration options, visit [here](/chart/kubed/README.md).


## Verify installation
To check if Kubed operator pods have started, run the following command:
```console
$ kubectl get pods --all-namespaces -l app=kubed --watch
```

Once the operator pods are running, you can cancel the above command by typing `Ctrl+C`.

Now, you are ready to [start managing your cluster](/docs/tutorials/README.md) using Kubed.
