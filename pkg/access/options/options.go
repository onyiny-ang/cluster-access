package main

const (
	DefaultContext = ""
)

type SubcommandOptions struct {
	ClusterName string
	Context     string
	User        string
	Kubeconfig  string
	Namespace   string
}

// BindCommon adds the common options that are shared by different
// sub-commands to the list of flags.
func (o *SubcommandOptions) BindCommon(flags *pflag.FlagSet) {
	flags.StringVar(&o.Kubeconfig, "kubeconfig", "",
		"Path to the kubeconfig file to use for CLI requests.")
	flags.StringVar(&o.ClusterName, "cluster-access-context", "",
		"Name of the cluster which will be added to/deleted from the kubeconfig file.")
	flags.StringVar(&o.Namespace, "cluster-namespace",
		"",
		"Namespace to be created in the cluster being added to kubeconfig")
	flags.StringVar(&o.User, "user", "admin",
		"User to be used to authorize use of the cluster.")
	flags.StringVar(&o.Context, "context", DefaultContext,
		"The context from which the cluster is created is used for demonstrative purposes.")
}
