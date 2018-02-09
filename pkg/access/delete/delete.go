package delete

import (
	"io"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"k8s.io/cluster-access/pkg/access/options"
)

var (
	deleteLong = `
    Deletes an entry for the specified cluster-registry cluster from KUBECONFIG (requires cluster-name)`

	deleteExample = `
    #Delete an entry in kubeconfig for cluster-registry cluster "test-cluster1" existing in the minikube context
    cluster-access delete cluster-name=test-cluster1
	`
)

type deleteOptions struct {
	options.SubcommandOptions
}

func NewCmdDelete(cmdOut io.Writer) *cobra.Command {
	opts := &deleteOptions{}

	cmd := &cobra.Command{
		Use:     "delete [cluster-name=name]",
		Short:   "deletes a specified cluster from KUBECONFIG (requires -c)",
		Long:    deleteLong,
		Example: deleteExample,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				glog.Fatalf("Error: missing required arguments")
			}
		},
	}
	flags := cmd.Flags()
	opts.BindCommon(flags)

	// cmd.Flags().String("cluster-name", "", "Name of the cluster to be created/deleted in KUBECONFIG")
	// cmd.Flags().String("kube-location", home+"/.kube/config", "Indicate location of kube config file")
	return cmd
}
