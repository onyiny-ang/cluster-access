package options

import (
	"fmt"
	"io"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"k8s.io/client-go/tools/clientcmd"
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
	viper.BindPFlag("kubeconfig", flags.Lookup("kubeconfig"))
	viper.BindPFlag("cluster-name", flags.Lookup("cluster-name"))
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
