package create

import (
	"io"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"k8s.io/cluster-access/pkg/access/options"
)

var (
	createLong = `
	Creates an entry for the specified cluster-registry cluster from KUBECONFIG (requires cluster-name)
	`

	createExample = `
	#Create an entry in kubeconfig for cluster-registry cluster "test-cluster1" existing in the minikube context
	cluster-access create cluster-name=test-cluster1 kube-context=minikube user=tester
	`
)

type createOptions struct {
	options.SubcommandOptions
}

func NewCmdCreate(cmdOut io.Writer) *cobra.Command {
	opts := &createOptions{}

	cmd := &cobra.Command{
		Use:     "create [cluster-name=name] [kube-context=context] [user=user]",
		Short:   "creates a specified cluster from KUBECONFIG (requires -c)",
		Long:    createLong,
		Example: createExample,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 3 {
				glog.Fatalf("Error: missing required arguments")
			}

		},
	}
	flags := cmd.Flags()
	opts.BindCommon(flags)

	//	cmd.Flags().String("cluster-name", "", "Name of the cluster to be created/deleted in KUBECONFIG")
	//	cmd.Flags().String("kube-context", "", "Existing context where cluster-registry and cluster exist")
	//	cmd.Flags().String("user", "", "User name for credential creation")
	//	cmd.Flags().String("kube-location", home+"/.kube/config", "Indicate location of kube config fil")
	//	cmd.Flags().String("namespace", "", "Namespace for specified cluster")
	return cmd
}
