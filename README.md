## Cluster access tool (CAT tool)

####  DESCRIPTION
A prototype tool to allow a cluster in a cluster registry to be easily added or deleted from the kubeconfig file

```
  Usage: ./cluster-access  [command] [- | -- ][arguments]"
    Commands:"
      create    creates an entry for a specified cluster, context and user in KUBECONFIG (requires -k, -c, -u flags)"
      delete    deletes an entry for the specified cluster in KUBECONFIG (requires -c)"
    Required arguments:"
      -c, --cluster-name    Cluster to be created/deleted name in KUBECONFIG"
      -k, --kube-context    Existing context where cluster-registry and cluster exist"
      -u, --user            User name for credential creation"
    Optional Arguments:"
      -h, --help            Displays this usage"
      -v, --verbose         Increase verbosity for debugging"
      -l, --kube-location   Indicate location of kube config file"
      -n, --namespace       Namespace for specified cluster"
```
See the [example](https://github.com/onyiny-ang/cluster-access/blob/master/example.md) for instructions on setting up

