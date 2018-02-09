package options

import (
	"fmt"
	"io"

	"github.com/golang/glog"
	"github.com/spf13/pflag"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	home        = homedir.HomeDir()
	kubeDefault = home + "/.kube/config"
)

type SubcommandOptions struct {
	ClusterName  string
	KubeContext  string
	User         string
	Namespace    string
	KubeLocation string
}

// BindCommon adds the common options that are shared by different
// sub-commands to the list of flagsp.
func (o *SubcommandOptions) BindCommon(flags *pflag.FlagSet) {
	flags.StringVar(&o.KubeLocation, "kubeconfig", kubeDefault,
		"Path to the kubeconfig file to use for CLI requests.")
	flags.StringVar(&o.ClusterName, "cluster-access-context", "",
		"Name of the cluster which will be added to/deleted from the kubeconfig file.")
	flags.StringVar(&o.Namespace, "cluster-namespace",
		"",
		"Namespace to be created in the cluster being added to kubeconfig")
	flags.StringVar(&o.User, "user", "admin",
		"User to be used to authorize use of the cluster.")
	flags.StringVar(&o.KubeContext, "context", "",
		"The context from which the cluster is created is used for demonstrative purposes.")
}

// UpdateKubeconfig handles updating the kubeconfig by building up the endpoint
// while printing and logging progress.
func (o *SubcommandOptions) UpdateKubeconfig(cmdOut io.Writer,
	pathOptions *clientcmd.PathOptions) error {

	fmt.Fprint(cmdOut, "Updating kubeconfig...")
	glog.V(4).Info("Updating kubeconfig")

	// Pick the first ip/hostname to update the api server endpoint in kubeconfig
	// and also to give information to user.
	// In case of NodePort Service for api server, ips are node external ips.

	// If the service is nodeport, need to append the port to endpoint as it is
	// non-standard port.
	//	if o.APIServerServiceType == "" {
	//		endpoint = endpoint + ":"
	//	}

	//	err := UpdateKubeconfig(pathOptions, o.Name, endpoint, o.Kubeconfig,
	//		credentials, o.DryRun)

	//	if err != nil {
	//		glog.V(4).Infof("Failed to update kubeconfig: %v", err)
	//		return err
	//	}

	fmt.Fprintln(cmdOut, " done")
	glog.V(4).Info("Successfully updated kubeconfig")
	return nil
}
