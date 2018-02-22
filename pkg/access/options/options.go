package options

import (
	"github.com/spf13/pflag"

	"k8s.io/client-go/util/homedir"
)

var (
	home         = homedir.HomeDir()
	kubeDefault  = home + "/.kube/config"
	kubeUsage    = "Path to the kubeconfig file to use for CLI requests."
	clusterUsage = "Name of the cluster which will be added to/delete  d from the kubeconfig file."
)

type SubcommandOptions struct {
	ClusterName  string
	KubeLocation string
}

// BindCommon adds the common options that are shared by different
// sub-commands to the list of flags.
func (o *SubcommandOptions) BindCommon(flags *pflag.FlagSet) {
	flags.StringVarP(&o.KubeLocation, "kubeconfig", "k", kubeDefault, kubeUsage)
	flags.StringVarP(&o.ClusterName, "cluster-name", "c", "", clusterUsage)
}
