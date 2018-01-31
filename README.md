## Cluster access tool (CAT tool)

####  DESCRIPTION
A prototype tool to allow a cluster in a cluster registry to be easily added or deleted from the kubeconfig file

```
  Usage: ./cluster-access  [command] [- | -- ][arguments]"
    Commands:"
      create    creates an entry for a specified cluster, context and user in KUBECONFIG (requires -k, -c, -u flags)"
      delete    deletes an entry for the specified cluster in KUBECONFIG (requires -c)"
    Required arguments:"
      -c, --cluster-name    specifies the desired cluster to be created/deleted name in KUBECONFIG"
      -k, --kubeconfig-entry-context  creates an entry for the specified cluster in KUBECONFIG"
      -u, --user    creates an entry for the specified cluster in KUBECONFIG"
    Optional Arguments:"
      -h, --help             Display this usage"
      -v, --verbose          Increase verbosity for debugging"
      -l, --kube-location    Indicate location of kube config file"
      -n, --namespace  creates an entry for the specified cluster in KUBECONFIG"
```


