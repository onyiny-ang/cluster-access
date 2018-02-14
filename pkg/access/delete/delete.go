package delete

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cluster-access/pkg/access/options"
)

var (
	deleteLong = `
    Deletes an entry for the specified cluster-registry cluster from KUBECONFIG`

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

	deleteCmd := &cobra.Command{
		Use:     "delete [cluster-name=name]",
		Short:   "deletes a specified cluster from KUBECONFIG",
		Long:    deleteLong,
		Example: deleteExample,
		Run: func(deleteCmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Printf(deleteCmd.Flags().FlagUsages())
				os.Exit(1)

			}
		},
	}
	flags := deleteCmd.Flags()
	opts.BindCommon(flags)
	deleteCmd.MarkPersistentFlagRequired("cluster-name")

	return deleteCmd

}
