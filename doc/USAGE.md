# Usage

The following assumes you have the plugin installed via

```shell
kubectl krew install kopy
```

## Scan pvs in your current kubecontext

```shell
kubectl kopy
```

## How it works

Gather all PVs and select a pv to download its content. If the content is empty, it means the PV is not used by any workloads.
