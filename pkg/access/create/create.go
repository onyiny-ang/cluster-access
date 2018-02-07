package create

import (
	"io"

	"github.com/spf13/cobra"

	"k8s.io/kubernetes/pkg/kubectl/cmd/templates"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/util/i18n"
)

var (
	createLong = templates.LongDesc(`Creates an entry for the specified cluster-registry cluster from KUBECONFIG (requires -c)`)

	createExample = templates.Examples(i18n.T(`
	#Create an entry in kubeconfig for cluster-registry cluster "test-cluster1" existing in the minikube context
	cluster-access create cluster-name=test-cluster1 kube-context=minikube user=tester
	`))
)

func NewCmdCreate(cmdOut io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:     "create [cluster-name=name] [kube-context=context] [user=user]",
		Short:   "creates a specified cluster from KUBECONFIG (requires -c)",
		Long:    createLong,
		Example: createExample,
	}

	cmdutil.AddPrinterFlags(cmd)
	cmd.Flags().String("cluster-name", "", i18n.T("Name of the cluster to be created/deleted in KUBECONFIG"))

}
