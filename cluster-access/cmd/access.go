package main

import (
	"flag"
	"os"
)

func main() {

	helpCommand := flag.String("help", "", "Usage: '$'$0  [command] [- | -- ][arguments]\nCommands:\ncreate\tCreates an entry for a specified cluster, context and user in KUBECONFIG (requires -k, -c, -u flags)\ndelete\tdeletes an entry for the specified cluster in KUBECONFIG (requires -c)\n Required arguments:\n-c, --cluster-name\tName of the cluster to be created/deleted in KUBECONFIG\n-k, --kube-context\tExisting context where cluster-registry and cluster exist\n-u, --user\tUser name for credential creation\nOptional Arguments:\n-h, --help\t Display this usage\n-v, --verbose\t Increase verbosity for debugging\n-l, --kube-location\tIndicate location of kube config file\n-n, --namespace\tNamespace for specified cluster")
	createCommand := flag.NewFlagSet("create", flag.ExitOnError)
	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)

	//create subcommands
	createCluster := createCommand.String("cluster-name", "", "Name of the cluster to be created in KUBECONFIG")
	createContext := createCommand.String("kube-context", "", "Existing context where cluster-registry and cluster exist")
	createNamespace := createCommand.String("namespace", "default", "Namespace for specified cluster")
	createUser := createCommand.String("user", "", "User name for credential creation")
	createLocation := createCommand.String("kube-location", "", "Indicate location of KUBECONFIG file")

	//delete subcommands
	deleteCluster := deleteCommand.String("cluster-name", "", "Name of the cluster to be deleted in KUBECONFIG")
	deleteLocation := deleteCommand.String("kube-location", "", "Indicate location of KUBECONFIG file")

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "create":
		createCommand.Parse(os.Args[2:])
	case "delete":
		deleteCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
	os.Exit(1)
}
