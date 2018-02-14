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

	cmd := &cobra.Command{
		Use:     "delete [cluster-name=name]",
		Short:   "deletes a specified cluster from KUBECONFIG",
		Long:    deleteLong,
		Example: deleteExample,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Printf(cmd.Flags().FlagUsages())
				os.Exit(1)

			}
		},
	}
	flags := cmd.Flags()
	opts.BindCommon(flags)
	cmd.MarkFlagRequired("cluster-context")
	flags.Parse(os.Args)

	return cmd
}
