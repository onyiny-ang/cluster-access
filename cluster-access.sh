#!/bin/bash
#==========================================================
#
#  FILE:  access.sh
#  DESCRIPTION: A prototype tool to allow a cluster in a cluster registry
#               to be easily added or deleted from the kubeconfig file
#  AUTHOR:  Lindsey Tulloch , ltulloch@redhat.com
#  VERSION:  1.0
#  CREATED:  2018-01-24 03:46:46 PM EST
#  REVISION:  ---
#============================================================

#set -e

function usage {
  echo "Usage: '$'$0  [command] [- | -- ][arguments]"
  echo "  Commands:"
  echo "    create    creates an entry for a specified cluster, context and user in KUBECONFIG (requires -k, -c, -u flags)"
  echo "    delete    deletes an entry for the specified cluster in KUBECONFIG (requires -c)"
  echo "  Required arguments:"
  echo "    -c, --cluster-name    specifies the desired cluster to be created/deleted name in KUBECONFIG"
  echo "    -k, --kubeconfig-entry-context  creates an entry for the specified cluster in KUBECONFIG"
  echo "    -u, --user    creates an entry for the specified cluster in KUBECONFIG"
  echo "  Optional Arguments:"
  echo "    -h, --help             Display this usage"
  echo "    -v, --verbose          Increase verbosity for debugging"
  echo "    -l, --kube-location    Indicate location of kube config file"
  echo "    -n, --namespace  creates an entry for the specified cluster in KUBECONFIG"
}

function parse_args {
    req_arg_count=0
    NAMESPACE="default"

      if [[ ${1} == '-h' || ${1} == '--help' ]]; then
          usage
          exit 1
      else
        case "${1}" in
          create)
            CREATE=true
            shift
            ;;
          delete)
            CREATE=false
            shift
            ;;
          *)
            echo "Error: invalid argument '${arg}'"
            usage
            exit 1
            ;;
        esac
      fi

      if [[ ! $CREATE && $# -lt 3 ]]; then
          case "${1}" in
            -l|--kube-location)
              KUBECONFIG_LOCATION="${2}"
              shift
              ;;
            -c|--cluster-name)
              CLUSTER="${2}"
              (( req_arg_count += 1 ))
              shift
              ;;
            *)
              echo "Error: invalid argument '${arg}'"
              usage
              exit 1
              ;;
          esac
          shift
      else
        while [[ ${CREATE} && $# -gt 1 ]]; do
          case "${1}" in
              -k|--kubeconfig-entry-context)
                  CONTEXT="${2}"
                  (( req_arg_count += 1 ))
                  shift
                  ;;
              -l|--kube-location)
                  KUBECONFIG_LOCATION="${2}"
                   shift
                   ;;
              -n|--namespace)
                 NAMESPACE="${2}"
                 (( req_arg_count += 1 ))
                 shift
                 ;;
              -c|--cluster-name)
                 CLUSTER="${2}"
                 (( req_arg_count += 1 ))
                 shift
                 ;;
              -u|--user)
                 USER="${2}"
                 (( req_arg_count += 1 ))
                 shift
                 ;;
              -v|--verbose)
                 set -x
                 ;;
              -h|--help)
                 usage
                 exit 1
                 ;;
              *)
                 echo "Error: invalid argument '${arg}'"
                 usage
                 exit 1
                 ;;
         esac
         shift
       done
     fi

    if [[ $CREATE && ${req_arg_count} -ne 3 ]] | [[ ! $CREATE && ${req_arg_count} -ne 1 ]]; then
    echo "Error: missing required arguments"
        usage
        exit 1
    fi

    if [[ ${KUBECONFIG_LOCATION} == '' ]]; then
      if [[ ! -f "$HOME/.kube/config" ]]; then
        echo "ERROR: Couldn't find Kubeconfig file, please specify with -k flag"
      else
        KUBECONFIG_LOCATION="$HOME/.kube/config"
      fi
    fi
}

function validate_args {
    if ! $(kubectl config get-contexts -o name | grep ${CONTEXT} &> /dev/null); then
        echo "Error: context name '${CONTEXT}' is not valid. Please check the context name and try again."
        usage
        exit 1
    fi
    if ! $(kubectl get clusters --context=${CONTEXT} -o name | grep ${CLUSTER} &> /dev/null); then
        echo "Error: cluster name '${CLUSTER}' is not valid. Please check the cluster name and try again."
        usage
        exit 1
    fi
}

function validate_cluster {
    if ! $(kubectl get clusters --context=${CLUSTER} -o name | grep ${CLUSTER} &> /dev/null); then
        echo "Error: cluster name '${CLUSTER}' is not valid. Please check the cluster name and try again."
        usage
        exit 1
    fi
}

function action {
  if ${CREATE} ; then
    validate_args
    create_kubeentry
  else
    validate_cluster
    delete_kubeentry
  fi
}

function create_kubeentry {

SERVER="$(kubectl get clusters ${CLUSTER} -o json | jq -r '.spec.kubernetesApiEndpoints.serverEndpoints[].serverAddress')"

KUBE_CA="$(kubectl config view -o json | jq -r --arg CONTEXT "$CONTEXT" '.clusters[] | select(.name==$CONTEXT)| .cluster | .["certificate-authority"]')"

KUBE_CERT="$(kubectl config view -o json | jq -r --arg CONTEXT "$CONTEXT" '.users[] | select(.name==$CONTEXT)| .user | .["client-certificate"]')"

KUBE_KEY="$(kubectl config view -o json | jq -r --arg CONTEXT "$CONTEXT" '.users[] | select(.name==$CONTEXT)| .user | .["client-key"]')"

kubectl config --kubeconfig=${KUBECONFIG_LOCATION} set-cluster ${CLUSTER} --cluster=${CLUSTER} --namespace=${NAMESPACE} --user=${CLUSTER} --server=${SERVER} --certificate-authority=${KUBE_CA}

echo "Created cluster '${CLUSTER}' in kubeconfig"

kubectl config --kubeconfig=${KUBECONFIG_LOCATION} set-credentials ${USER} --client-certificate=${KUBE_CERT} --client-key=${KUBE_KEY}

kubectl config --kubeconfig=${KUBECONFIG_LOCATION} set-context ${CLUSTER} --cluster=${CLUSTER} --namespace=${NAMESPACE} --user=${USER}

echo "Created context '${CLUSTER}' in kubeconfig"

}

function delete_kubeentry {

KUBE_USER="$(kubectl config view --flatten -o json | jq -r --arg CLUSTER "$CLUSTER" '.contexts[] | select(.name==$CLUSTER)| .context.user')"

kubectl config unset users.${KUBE_USER}
kubectl config unset clusters.${CLUSTER}
echo "Removed cluster ${CLUSTER} from kubeconfig"
kubectl config unset contexts.${CLUSTER}
echo "Removed context ${CLUSTER} from kubeconfig"


}

function main {
  parse_args $@
  action
}

main $@

