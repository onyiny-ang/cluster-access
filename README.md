## Cluster access tool (CAT tool)

####  DESCRIPTION
A prototype tool to allow a cluster in a cluster registry to be easily added or deleted from the kubeconfig file
This is not meant to be used in production as there are significant challenges in retrieving accurate authinfo for each cluster.
This prototype uses the context in which the cluster registry and cluster being moved to the kubeconfig file is created to retrieve auth info which works as a proof of concept in this particular case.

#### BUILD:

The cluster access tool is a command line tool that has bash and go implementations. It has been tested with a cluster-registry running in minikube (the `cluster.sh` script deploys the cluster-registry as an aggregated API server in minikube and creates 3 registered test clusters).

##### To run with Golang

Clone the repository into $GOPATH/src/k8s.io/cluster-access.

1.  run `bazel run //:gazelle`
Note: You must have a recent version of [bazel](https://bazel.io) installed.

2. run `bazel build //cmd/access`

This command will create the access tool in `bazel-bin/cmd/access/.../access`. The tool can be used from there or moved to a more easily accessible location.


##### Usage:


```
  Usage: ./access  [command] [- | -- ][arguments]"
    Commands:"
      create    creates an entry for a specified cluster, context and user in KUBECONFIG (requires -k, -c, -u flags)"
      delete    deletes an entry for the specified cluster in KUBECONFIG (requires -c)"
      help
  Flags:
    -c, --cluster-name string        Name of the cluster which will be added to/delete  d from the kubeconfig file.
    -n, --cluster-namespace string   Namespace to be created in the cluster being adde  d to kubeconfig (default "default")
    -h, --help                       help for specified command
    -x, --kube-context string        The context from which the cluster is c  reated is used for demonstrative purposes.
    -k, --kubeconfig string          Path to the kubeconfig file to use for CLI requests. (default "$HOME/.kube/config")
    -u, --user string                User to be used to authorize use of the cluster. (default "admin")

```
See the [example](https://github.com/onyiny-ang/cluster-access/blob/master/example.md) for instructions on setting up

If you see the need to run dep ensure on this repo, you will need to keep a few things in mind.
 1. There is a broken bazel rule when vendoring in client-go/apimachinery and you will need to [Comment this out after](https://github.com/scele/apimachinery/commit/15dc092229cda2ca7ead32667e463b53f4a7c138)
 2. There is another issue with the BUILD file in vendor/k8s.io/client-go/util/cert to have the go_library not reference testdata.
 These are the known issues at this time. There may be others that have not yet been found.
 
 #### Demo of Cluster Access Tool
 
 You can watch a demo of the tool in use [here](https://www.youtube.com/watch?v=V_TIOaVIW8k)
