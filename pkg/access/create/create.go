package create

import (
	"fmt"
	"io"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	client "k8s.io/client-go/kubernetes"
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
			err := opts.validateFlags(pathOptions)
			if err != nil {
				glog.Fatalf("error: %v", err)
			}
			hostConfig, err := util.GetClientConfig(pathOptions, opts.Kubecontext, opts.KubeLocation).ClientConfig()
			if err != nil {
				glog.Fatalf("error: %v", err)
			}
			fmt.Println(hostConfig.Cluster.Server)
			hostClientset, err := client.NewForConfig(hostConfig)
			if err != nil {
				glog.Fatalf("error: %v", err)
			}
			fmt.Println(hostClientset)
			pathOptions.LoadingRules.ExplicitPath = opts.KubeLocation
			opts.UpdateKubeconfig(cmdOut, pathOptions)
			createRun(opts, createCmd, args)
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

func (o *createOptions) validateFlags(pathOptions *clientcmd.PathOptions) error {
	config, err := pathOptions.GetStartingConfig()
	if err != nil {
		return err
	}
	if _, exists := config.Contexts[o.Kubecontext]; !exists {
		glog.V(4).Info("error: context %v not found", o.Kubecontext)
		return err
	}
	clientset, err := crclientset.NewForConfig(config)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if cluster, err := clientset.ClusterregistryV1alpha1().Clusters().Get(o.ClusterName, metav1.GetOptions{}); err != nil {
		glog.V(4).Info("error: cluster %v not found", o.ClusterName)
		return err
	}
	return nil
}

func createRun(o *createOptions, createCmd *cobra.Command, args []string) {

	fmt.Println("Don't forget to implement or delete me!")
	glog.V(4).Info("Testing some stuff here")
}
