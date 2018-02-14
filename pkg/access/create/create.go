package create

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"k8s.io/cluster-access/pkg/access/options"
)

var (
	createLong = `
	Creates an entry for the specified cluster-registry cluster from KUBECONFIG	`

	createExample = `
	#Create an entry in kubeconfig for cluster-registry cluster "test-cluster1" existing in the minikube context
	cluster-access create cluster-name=test-cluster1 kube-context=minikube user=tester
	`
	namespaceUsage = "Namespace to be created in the cluster being adde  d to kubeconfig"
	userUsage      = "User to be used to authorize use of the cluster."
	kubeUsage      = "The context from which the cluster is c  reated is used for demonstrative purposes."
)

type createOptions struct {
	options.SubcommandOptions
	Kubecontext string
	User        string
	Namespace   string
}

func NewCmdCreate(cmdOut io.Writer) *cobra.Command {
	opts := &createOptions{}

	createCmd := &cobra.Command{
		Use:     "create",
		Short:   "creates a specified cluster from KUBECONFIG",
		Long:    createLong,
		Example: createExample,
		Args:    cobra.MinimumNArgs(3),
		Run: func(createCmd *cobra.Command, args []string) {
			createRun(opts, createCmd, args)
		},
	}
	flags := createCmd.PersistentFlags()
	opts.BindCommon(flags)
	opts.BindCreate(flags)
	createCmd.MarkFlagRequired("user")
	createCmd.MarkFlagRequired("cluster-name")
	createCmd.MarkFlagRequired("kube-context")

	return createCmd
}

func (o *createOptions) BindCreate(flags *pflag.FlagSet) {
	flags.StringVarP(&o.Namespace, "cluster-namespace", "n", "default", namespaceUsage)
	flags.StringVarP(&o.User, "user", "u", "admin", userUsage)
	flags.StringVarP(&o.Kubecontext, "kube-context", "x", "", kubeUsage)
	viper.BindPFlag("cluster-namespace", flags.Lookup("cluster-namespace"))
	viper.BindPFlag("kube-context", flags.Lookup("kube-context"))
	viper.BindPFlag("user", flags.Lookup("user"))

}

func createRun(o *createOptions, createCmd *cobra.Command, args []string) {
	if len(args) < 3 {
		fmt.Printf(createCmd.Flags().FlagUsages())
		os.Exit(1)
	}

	glog.V(4).Info("Testing some stuff here")

	fmt.Println(strings.Join(args, " "))
}
