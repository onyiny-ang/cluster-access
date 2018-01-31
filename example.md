## For example:
### created cluster:

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

A similar set up will be set up can be instantiated on your kubernetes cluster by using the [clusters.sh](https://github.com/onyiny-ang/cluster-access/blob/master/clusters.sh) script. The script assumes your `current-context` is the one that you want to use for testing.



