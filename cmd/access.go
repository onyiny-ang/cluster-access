package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/util/homedir"
)

func main() {

	helpCommand := pflag.String("help", "", "Usage: '$'$0  [command] [- | -- ][arguments]\nCommands:\ncreate\tCreates an entry for a specified cluster, context and user in KUBECONFIG (requires -k, -c, -u pflags)\ndelete\tdeletes an entry for the specified cluster in KUBECONFIG (requires -c)\n Required arguments:\n-c, --cluster-name\tName of the cluster to be created/deleted in KUBECONFIG\n-k, --kube-context\tExisting context where cluster-registry and cluster exist\n-u, --user\tUser name for credential creation\nOptional Arguments:\n-h, --help\t Display this usage\n-v, --verbose\t Increase verbosity for debugging\n-l, --kube-location\tIndicate location of kube config file\n-n, --namespace\tNamespace for specified cluster")
	createCommand := pflag.NewFlagSet("create", pflag.ExitOnError)
	deleteCommand := pflag.NewFlagSet("delete", pflag.ExitOnError)

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
		fmt.Printf("%s", *helpCommand)
		os.Exit(1)
	}
	switch os.Args[1] {
	case "create":
		createCommand.Parse(os.Args[2:])
	case "delete":
		deleteCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("%s", *helpCommand)
		os.Exit(1)
	}

	if createCommand.Parsed() {
		if *createContext == "" {
			createCommand.PrintDefaults()
			os.Exit(1)
		}
		if *createCluster == "" {
			createCommand.PrintDefaults()
			os.Exit(1)
		}
		if *createUser == "" {
			createCommand.PrintDefaults()
			os.Exit(1)
		}
		client := fake.NewSimpleClientset()
		stringy := *createContext
		info, _ := client.CoreV1().ConfigMaps(stringy).Get(stringy, metav1.GetOptions{})
		if info.Name != *createContext {
			createCommand.PrintDefaults()
			fmt.Println("wrong name")
			os.Exit(1)
		}
		if *createLocation == "" {
			home := homedir.HomeDir()
			createCommand.Set(*createLocation, home+"/.kube/config")
		}

		fmt.Printf("%s", *createNamespace)

	}

	if deleteCommand.Parsed() {
		if *deleteCluster == "" {
			deleteCommand.PrintDefaults()
			os.Exit(1)
		}
		if *deleteLocation == "" {
			home := homedir.HomeDir()
			createCommand.Set(*deleteLocation, home+"/.kube/config")
		}

	}

	os.Exit(1)
}
