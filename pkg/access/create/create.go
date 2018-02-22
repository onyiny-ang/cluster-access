package create

import (
	"fmt"
	"io"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/cluster-access/pkg/access/options"
	"k8s.io/cluster-access/pkg/access/util"
	crclientset "k8s.io/cluster-registry/pkg/client/clientset_generated/clientset"
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
		Run: func(createCmd *cobra.Command, args []string) {
			pathOptions := clientcmd.NewDefaultPathOptions()
			hostConfig, err := util.GetClientConfig(pathOptions, opts.Kubecontext, opts.KubeLocation).ClientConfig()
			if err != nil {
				glog.Fatalf("error: %v", err)
			}
			err = opts.validateFlags(pathOptions, hostConfig)
			if err != nil {
				glog.Fatalf("error: %v", err)
			}
			pathOptions.LoadingRules.ExplicitPath = opts.KubeLocation
			createRun(cmdOut, opts, hostConfig, pathOptions, createCmd, args)
		},
	}
	flags := createCmd.PersistentFlags()
	opts.BindCommon(flags)
	opts.BindCreate(flags)
	createCmd.MarkPersistentFlagRequired("user")
	createCmd.MarkPersistentFlagRequired("cluster-name")
	createCmd.MarkPersistentFlagRequired("kube-context")
	return createCmd
}

func (o *createOptions) BindCreate(flags *pflag.FlagSet) {
	flags.StringVarP(&o.Namespace, "cluster-namespace", "n", "default", namespaceUsage)
	flags.StringVarP(&o.User, "user", "u", "admin", userUsage)
	flags.StringVarP(&o.Kubecontext, "kube-context", "x", "", kubeUsage)

}

func (o *createOptions) validateFlags(pathOptions *clientcmd.PathOptions, hostConfig *rest.Config) error {
	// ensure context exists
	config, err := pathOptions.GetStartingConfig()
	if err != nil {
		return err
	}
	if _, exists := config.Contexts[o.Kubecontext]; !exists {
		glog.V(4).Info("error: context %v not found", o.Kubecontext)
		return err
	}
	clientset, err := crclientset.NewForConfig(hostConfig)
	if err != nil {
		glog.Fatalf("Unexpected error: %v", err)
	}
	//ensure cluster exists
	if _, err := clientset.ClusterregistryV1alpha1().Clusters().Get(o.ClusterName, metav1.GetOptions{}); err != nil {
		glog.V(4).Info("error: cluster %v not found", o.ClusterName)
		return err
	}
	return nil
}

func createRun(cmdOut io.Writer, o *createOptions, hostConfig *rest.Config, pathOptions *clientcmd.PathOptions, createCmd *cobra.Command, args []string) {
	crclient, err := crclientset.NewForConfig(hostConfig)
	//retrieve server address of cluster
	crcluster, err := crclient.ClusterregistryV1alpha1().Clusters().Get(o.ClusterName, metav1.GetOptions{})
	server := crcluster.Spec.KubernetesAPIEndpoints.ServerEndpoints[0].ServerAddress
	if err != nil {
		glog.Fatalf("Unexpected error: %v", err)
	}

	err = UpdateKubeconfig(cmdOut, o, pathOptions, server)
	if err != nil {
		glog.Fatalf("Unexpected error: th kubeconfig update %v", err)
	}

	fmt.Fprint(cmdOut, "Success")
	glog.V(4).Info("Cluster added to kubeconfig")
}

func UpdateKubeconfig(cmdOut io.Writer, o *createOptions, pathOptions *clientcmd.PathOptions, server string) error {

	kubeconfig, err := pathOptions.GetStartingConfig()
	if err != nil {
		glog.Fatalf("Unexpected error: %v", err)
	}

	//Populate API server endpoint info.
	cluster := clientcmdapi.NewCluster()
	cluster.Server = server
	cluster.CertificateAuthority = kubeconfig.Clusters[o.Kubecontext].CertificateAuthority

	//Populate Auth data (note this is just for demonstrative purposes and is not meant to   be a reasonable way to retrieve auth info in any other context)
	authInfo := clientcmdapi.NewAuthInfo()
	authInfo.ClientCertificate = kubeconfig.AuthInfos[o.Kubecontext].ClientCertificate
	authInfo.ClientKey = kubeconfig.AuthInfos[o.Kubecontext].ClientKey

	// Populate context.
	context := clientcmdapi.NewContext()
	context.Cluster = o.ClusterName
	context.AuthInfo = o.User
	context.Namespace = o.Namespace

	fmt.Fprint(cmdOut, "Updating kubeconfig...")
	glog.V(4).Info("Updating kubeconfig...")

	// Update the config struct with API server endpoint info,
	// credentials and context.
	kubeconfig.Clusters[o.ClusterName] = cluster
	kubeconfig.AuthInfos[o.User] = authInfo
	kubeconfig.Contexts[o.ClusterName] = context

	if err := clientcmd.ModifyConfig(pathOptions, *kubeconfig, true); err != nil {
		return err
	}
	return nil
}
