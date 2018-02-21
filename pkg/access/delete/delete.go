package delete

import (
	"fmt"
	"io"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	client "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/cluster-access/pkg/access/options"
	"k8s.io/cluster-access/pkg/access/util"
	crclientset "k8s.io/cluster-registry/pkg/client/clientset_generated/clientset"
)

var (
	deleteLong = `
    Deletes an entry for the specified cluster-registry cluster from KUBECONFIG`

	deleteExample = `
    #Delete an entry in kubeconfig for cluster-registry cluster "test-cluster1" existing in the minikube context
    cluster-access delete cluster-name=test-cluster1
	`
)

type deleteOptions struct {
	options.SubcommandOptions
}

func NewCmdDelete(cmdOut io.Writer) *cobra.Command {
	opts := &deleteOptions{}

	deleteCmd := &cobra.Command{
		Use:     "delete [cluster-name=name]",
		Short:   "deletes a specified cluster from KUBECONFIG",
		Long:    deleteLong,
		Example: deleteExample,
		Run: func(deleteCmd *cobra.Command, args []string) {
			pathOptions := clientcmd.NewDefaultPathOptions()
			context := opts.ClusterName
			hostConfig, err := util.GetClientConfig(pathOptions, context, opts.KubeLocation).ClientConfig()
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
			deleteRun(opts, deleteCmd, args)
		},
	}
	flags := deleteCmd.Flags()
	opts.BindCommon(flags)
	deleteCmd.MarkPersistentFlagRequired("cluster-name")

	return deleteCmd

}

func (o *deleteOptions) validateFlags(pathOptions *clientcmd.PathOptions, hostConfig *rest.Config) error {
	//ensure Cluster Name exists as context
	config, err := pathOptions.GetStartingConfig()
	if err != nil {
		return err
	}
	if _, exists := config.Contexts[o.ClusterName]; !exists {
		glog.V(4).Info("error: context %v not found", o.ClusterName)
		return err
	}
	clientset, err := crclientset.NewForConfig(hostConfig)
	if err != nil {
		glog.Fatalf("Unexpected error: %v", err)
	}
	//ensure cluster exists in cluster registry
	if _, err := clientset.ClusterregistryV1alpha1().Clusters().Get(o.ClusterName, metav1.GetOptions{}); err != nil {
		glog.V(4).Info("error: cluster %v not found", o.ClusterName)
		return err
	}
	return nil
}

func deleteRun(opts *deleteOptions, deleteCmd *cobra.Command, args []string) {
	fmt.Println("Don't forget to implement or delete me!")
	glog.V(4).Info("Testing some stuff here")

}
