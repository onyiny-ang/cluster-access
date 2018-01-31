#!/bin/bash
#==========================================================
#
#  FILE:  clusters.sh
#  DESCRIPTION:  Creates test clusters for demonstration purposes
#  AUTHOR:  Lindsey Tulloch , ltulloch@redhat.com
#  VERSION:  1.0
#  CREATED:  2018-01-26 09:31:31 AM EST
#============================================================

CURRENT_CONTEXT="$(kubectl config current-context)"
SERVER="$(kubectl config view --flatten -o json | jq -r --arg CURRENT_CONTEXT "$CURRENT_CONTEXT" '.clusters[] | select(.name==$CURRENT_CONTEXT)| .cluster.server')"

$HOME/crinit aggregated init cr-test --host-cluster-context=${CURRENT_CONTEXT}

kubectl apply -f - --context ${CURRENT_CONTEXT} <<EOF
kind: Cluster
apiVersion: clusterregistry.k8s.io/v1alpha1
metadata:
  name: test-cluster1
spec:
  kubernetesApiEndpoints:
    serverEndpoints:
      - clientCidr: "0.0.0.0/0"
        serverAddress: "${SERVER}"
EOF

kubectl apply -f - --context ${CURRENT_CONTEXT} <<EOF
kind: Cluster
apiVersion: clusterregistry.k8s.io/v1alpha1
metadata:
  name: test-cluster2
spec:
  kubernetesApiEndpoints:
    serverEndpoints:
      - clientCidr: "0.0.0.0/0"
        serverAddress: "${SERVER}"
EOF

kubectl apply -f - --context ${CURRENT_CONTEXT} <<EOF
kind: Cluster
apiVersion: clusterregistry.k8s.io/v1alpha1
metadata:
  name: test-cluster3
spec:
  kubernetesApiEndpoints:
    serverEndpoints:
      - clientCidr: "0.0.0.0/0"
        serverAddress: "${SERVER}"
EOF
