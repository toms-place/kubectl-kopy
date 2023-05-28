# kubectl-kopy

A `kubectl` plugin to gather all PVs and select one to download its content.

## Quick Start

Install kopy via [krew](https://krew.sigs.k8s.io/)

```shell
kubectl krew install kopy

# Download the contents of a persistent volume in the current context

kubectl kopy [PV] [LOCAL_DIR]
```
