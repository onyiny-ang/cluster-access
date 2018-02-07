package create

import (
	"io"

	"github.com/spf13/cobra"

	"k8s.io/client-go/util/homedir"
)

var (
	createLong = `
	Creates an entry for the specified cluster-registry cluster from KUBECONFIG (requires -c)`

	createExample = `
	#Create an entry in kubeconfig for cluster-registry cluster "test-cluster1" existing in the minikube context
	cluster-access create cluster-name=test-cluster1 kube-context=minikube user=tester
	`
)

func NewCmdCreate(cmdOut io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:     "create [cluster-name=name] [kube-context=context] [user=user]",
		Short:   "creates a specified cluster from KUBECONFIG (requires -c)",
		Long:    createLong,
		Example: createExample,
	}

	home := homedir.HomeDir()
	cmd.Flags().String("cluster-name", "", "Name of the cluster to be created/deleted in KUBECONFIG")
	cmd.Flags().String("kube-context", "", "Existing context where cluster-registry and cluster exist")
	cmd.Flags().String("user", "", "User name for credential creation")
	cmd.Flags().String("kube-location", home+"/.kube/config", "Indicate location of kube config fil")
	cmd.Flags().String("namespace", "", "Namespace for specified cluster")
	return cmd
}
