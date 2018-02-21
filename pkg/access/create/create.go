package create

import (
	"fmt"
	"io"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	client "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
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
			hostClientset, err := client.NewForConfig(hostConfig)
			if err != nil {
				glog.Fatalf("error: %v", err)
			}
			fmt.Println("unsure if we need this: %v", hostClientset)
			pathOptions.LoadingRules.ExplicitPath = opts.KubeLocation
			opts.UpdateKubeconfig(cmdOut, pathOptions)
			createRun(opts, hostConfig, createCmd, args)
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

func createRun(o *createOptions, hostConfig *rest.Config, createCmd *cobra.Command, args []string) {
	clientset, err := crclientset.NewForConfig(hostConfig)
	//retrieve server address of cluster
	cluster, err := clientset.ClusterregistryV1alpha1().Clusters().Get(o.ClusterName, metav1.GetOptions{})
	svc := cluster.Spec.KubernetesAPIEndpoints.ServerEndpoints[0].ServerAddress
	if err != nil {
		glog.Fatalf("Unexpected error: %v", err)
	}
	//retrieve auth info from context

	//	updateKubeconfig(o)
	fmt.Println("Don't forget to implement or delete me!")
	glog.V(4).Info("Testing some stuff here")
}

// UpdateKubeconfig handles updating the kubeconfig by building up the endpoint
// while printing and logging progress.
//func (o *SubcommandOptions) UpdateKubeconfig(cmdOut io.Writer,
//   pathOptions *clientcmd.PathOptions, svc *v1.Service, ips, hostnames []string,
//   credentials *util.Credentials) error {
//
//   fmt.Fprint(cmdOut, "Updating kubeconfig...")
//   glog.V(4).Info("Updating kubeconfig")
//
//   // Pick the first ip/hostname to update the api server endpoint in kubeconfig
//   // and also to give information to user.
//   // In case of NodePort Service for api server, ips are node external ips.
//   endpoint := ""
//   if len(ips) > 0 {
//   	endpoint = ips[0]
//   } else if len(hostnames) > 0 {
//   	endpoint = hostnames[0]
//   }
//
//   // If the service is nodeport, need to append the port to endpoint as it is
//   // non-standard port.
//   if o.APIServerServiceType == v1.ServiceTypeNodePort {
//   	endpoint = endpoint + ":" + strconv.Itoa(int(svc.Spec.Ports[0].NodePort))
//   }
//
//   err := util.UpdateKubeconfig(pathOptions, o.Name, endpoint, o.Kubeconfig,
//   	credentials, o.DryRun)
//
//   if err != nil {
//   	glog.V(4).Infof("Failed to update kubeconfig: %v", err)
//   	return err
//   }
//
//   fmt.Fprintln(cmdOut, " done")
//   glog.V(4).Info("Successfully updated kubeconfig")
//   return nil
//
//
/// UpdateKubeconfig helper to update the kubeconfig file based on input
/// parameters.
//unc UpdateKubeconfig(o *createOptions, pathOptions *clientcmd.PathOptions, name, endpoint,
//   kubeConfigPath string, credentials *Credentials, dryRun bool) error {
//   pathOptions.LoadingRules.ExplicitPath = kubeConfigPath
//   kubeconfig, err := pathOptions.GetStartingConfig()
//   if err != nil {
//   	return err
//   }
//
//   // Populate API server endpoint info.
//   cluster := clientcmdapi.NewCluster()
//
//   // Prefix "https" as the URL scheme to endpoint.
//   if !strings.HasPrefix(endpoint, "https://") {
//   	endpoint = fmt.Sprintf("https://%s", endpoint)
//   }
//
//   cluster.Server = endpoint
//   cluster.CertificateAuthorityData = certutil.EncodeCertPEM(credentials.CertEntKeyPairs.CA.Cert)
//
//   // Populate credentials.
//   authInfo := clientcmdapi.NewAuthInfo()
//   authInfo.ClientCertificateData = certutil.EncodeCertPEM(credentials.CertEntKeyPairs.Admin.Cert)
//   authInfo.ClientKeyData = certutil.EncodePrivateKeyPEM(credentials.CertEntKeyPairs.Admin.Key)
//   authInfo.Token = credentials.Token
//
//   var httpBasicAuthInfo *clientcmdapi.AuthInfo
//
//   if credentials.Password != "" {
//   	httpBasicAuthInfo = clientcmdapi.NewAuthInfo()
//   	httpBasicAuthInfo.Password = credentials.Password
//   	httpBasicAuthInfo.Username = credentials.Username
//   }
//
//   // Populate context.
//   context := clientcmdapi.NewContext()
//   context.Cluster = o.ClusterName
//   context.AuthInfo = name
//
//   // Update the config struct with API server endpoint info,
//   // credentials and context.
//   kubeconfig.Clusters[name] = cluster
//   kubeconfig.AuthInfos[name] = authInfo
//
//   if httpBasicAuthInfo != nil {
//   	kubeconfig.AuthInfos[fmt.Sprintf("%s-basic-auth", name)] = httpBasicAuthInfo
//   }
//
//   kubeconfig.Contexts[name] = context
//
//   if err := clientcmd.ModifyConfig(pathOptions, *kubeconfig, true); err != nil {
//   	return err
//   }
//   return nil
//
