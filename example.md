# Try it out:

### Deploy cluster-registry and test clusters

A similar set up to that shown below can be instantiated on your kubernetes cluster to test out the cluster-access tool. Use the [clusters.sh](https://github.com/onyiny-ang/cluster-access/blob/master/clusters.sh) script. The script assumes your `current-context` is the one that you want to use for testing. You will need to have the [crinit tool](https://github.com/kubernetes/cluster-registry/blob/master/docs/userguide.md#deploying-a-cluster-registry) installed to deploy an [aggregated API server](https://github.com/kubernetes/cluster-registry/blob/master/docs/userguide.md#aggregated-api-server) to you kubernetes cluster.


[Deploy the cluster registry as an aggregated API Server](https://github.com/kubernetes/cluster-registry/blob/master/docs/userguide.md#aggregated-api-server)

Create a cluster using the example in the [doc](https://github.com/kubernetes/cluster-registry/blob/master/docs/userguide.md#try-it-out) to set up a test cluster.

The following code will create three clusters: `test-cluster1`, `test-cluster2`, `test-cluster3`.

```
kubectl apply -f - --context minikube <<EOF
kind: Cluster
apiVersion: clusterregistry.k8s.io/v1alpha1
metadata:
  name: test-cluster1
spec:
  kubernetesApiEndpoints:
    serverEndpoints:
      - clientCidr: "0.0.0.0/0"
        serverAddress: "100.200.300.4:8443"
EOF

kubectl apply -f - --context minikube <<EOF
kind: Cluster
apiVersion: clusterregistry.k8s.io/v1alpha1
metadata:
  name: test-cluster2
spec:
  kubernetesApiEndpoints:
    serverEndpoints:
      - clientCidr: "0.0.0.0/0"
        serverAddress: "100.200.300.4:8443"
EOF

kubectl apply -f - --context minikube <<EOF
kind: Cluster
apiVersion: clusterregistry.k8s.io/v1alpha1
metadata:
  name: test-cluster3
spec:
  kubernetesApiEndpoints:
    serverEndpoints:
      - clientCidr: "0.0.0.0/0"
        serverAddress: "100.200.300.4:8443"
EOF
```

### Create an entry in Kubeconfig

Using the [cluster-access.sh](https://github.com/onyiny-ang/cluster-access/blob/master/cluster-access.sh) script, you can create an entry in your kubeconfig file for one of the clusters--in this case, `test-cluster1` with the following command:

```
./cluster-access.sh create -k minikube -c test-cluster1 -u tester

```

This command also allows your to specify a namespace for the cluster and the location of your kubeconfig file, otherwise it will default to `$HOME/.kube/config`. The context and cluster names will be the same.

### Delete an entry in Kubeconfig

Using the same `test-cluster1` we can delete it from the kubeconfig file as follows:

```
./cluster-access.sh delete -c test-cluster1

```
