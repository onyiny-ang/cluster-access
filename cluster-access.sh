#!/bin/bash
#==========================================================
#
#  FILE:  access.sh
#  DESCRIPTION:
#  AUTHOR:  Lindsey Tulloch , ltulloch@redhat.com
#  VERSION:  1.0
#  CREATED:  2018-01-24 03:46:46 PM EST
#  REVISION:  ---
#============================================================

set -e

function usage {
  echo "$0: [OPTIONS] [- | -- ]"
  echo "  Optional Arguments:"
  echo "    -h, --help             Display this usage"
  echo "    -v, --verbose          Increase verbosity for debugging"
  echo "  Required arguments:"
  echo "    -c, --create-kubeconfig    creates an entry for the specified cluster in KUBECONFIG"
}

function parse_args {
    req_arg_count=0

    if [[ ${1} == '-h' || ${1} == '--help' ]]; then
        usage
        exit 1
    fi

    while [[ $# -gt 1 ]]; do
        case "${1}" in
            -c|--create-kubeconfig)
                CONTEXT_NAME="${2}"
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

    if [[ ${req_arg_count} -ne 5 ]]; then
    echo "Error: missing required arguments"
        usage
        exit 1
    fi

}

function validate_args {
    if ! $(kubectl config get-contexts -o name | grep ${CONTEXT_NAME} &> /dev/null); then
        echo "Error: cluster name '${CONTEXT_NAME}' is not valid. Please check the cluster name and try again."
        usage
        exit 1
    fi
}

function cleanup {

}

function main {
  parse_args $@
  validate_args
  cleanup
}

main $@

