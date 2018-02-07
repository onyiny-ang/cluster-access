package access

import (
	"flag"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewClusterAccessCommand(out io.Writer) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "cluster-access",
		Short: "cluster-access adds/deletes a cluster registry cluster to the kubeconfig file",
		Long:  "cluster-access adds/deletes a cluster in the cluster registry to the kubeconfig file for easy interaction",
	}

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	rootCmd.PersistentFlags().AddFlagSet(pflag.CommandLine)
	rootCmd.AddCommand(create.NewCmdCreate(out))
	rootCmd.AddCommand(delete.NewCmdDelete(out))

	return rootCmd
}
