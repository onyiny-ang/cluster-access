package delete

import (
	"io"

	"github.com/spf13/cobra"

	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

func NewCmdDelete(cmdOut io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "deletes a specified cluster from KUBECONFIG (requires -c)",
		Long:  "Deletes an entry for the specified cluster-registry cluster from KUBECONFIG (requires -c)",
	}

	cmdutil.AddPrinterFlags(cmd)
	cmd.Flags().String("cluster-name", "", "Name of the cluster to be created/deleted in KUBECONFIG")
	cmd.Flags().String("kube-location", home+"/.kube/config", "Indicate location of kube config file")

}
