package delete

import (
	"fmt"
	"io"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
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
			pathOptions.LoadingRules.ExplicitPath = opts.KubeLocation
			deleteRun(cmdOut, opts, hostConfig, pathOptions, deleteCmd, args)
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
	if err := clientset.ClusterregistryV1alpha1().RESTClient().Get().Resource("clusters").Name(o.ClusterName).Do().Error(); err != nil {
		glog.V(4).Info("error: cluster %v not found", o.ClusterName)
		fmt.Println(err)
		return err
	}
	return nil
}

func deleteRun(cmdOut io.Writer, opts *deleteOptions, hostConfig *rest.Config, pathOptions *clientcmd.PathOptions, deleteCmd *cobra.Command, args []string) error {

	errCount := 0
	fmt.Fprintf(cmdOut, "Delete kubeconfig entry %s...", opts.ClusterName)
	glog.V(4).Infof("Delete kubeconfig entry %s", opts.ClusterName)
	pathOptions.LoadingRules.ExplicitPath = opts.KubeLocation
	kubeconfig, err := pathOptions.GetStartingConfig()
	if err != nil {
		glog.Fatalf("Unexpected error: %v", err)
	}

	kubeconfigFile := pathOptions.GetDefaultFilename()
	if pathOptions.IsExplicitFile() {
		kubeconfigFile = pathOptions.GetExplicitFile()
	}

	_, ok := kubeconfig.Contexts[opts.ClusterName]
	if !ok {
		glog.V(4).Infof("cannot delete context %s, not in %s", opts.ClusterName, kubeconfigFile)
		errCount++
	} else {
		delete(kubeconfig.Contexts, opts.ClusterName)
	}
	_, ok = kubeconfig.Clusters[opts.ClusterName]
	if !ok {
		glog.V(4).Infof("cannot delete cluster %s, not in %s", opts.ClusterName, kubeconfigFile)
		errCount++
	} else {
		delete(kubeconfig.Clusters, opts.ClusterName)
	}

	_, ok = kubeconfig.AuthInfos[opts.ClusterName]
	if !ok {
		glog.V(4).Infof("cannot delete authinfo %s, not in %s", opts.ClusterName, kubeconfigFile)
		errCount++
	} else {
		delete(kubeconfig.AuthInfos, opts.ClusterName)
	}

	if errCount == 3 {
		fmt.Fprintf(cmdOut, "Could not find any context, cluster, or authinfo information for %s in your kubeconfig.", opts.ClusterName)
		glog.Fatalf("Could not find any context, cluster or authinfo for %s in kubeconfig", opts.ClusterName)
	}
	// Write the updated kubeconfig.
	if err := clientcmd.ModifyConfig(pathOptions, *kubeconfig, true); err != nil {
		glog.Fatalf("Unable to update kubeconfig, %s", err)
	}

	glog.V(4).Infof("deleted kubeconfig entry %s from %s\n", opts.ClusterName, kubeconfigFile)

	if kubeconfig.CurrentContext == opts.ClusterName {
		fmt.Fprintf(cmdOut,
			"warning: this removed your active context, use \"kubectl config use-context\" to select a different one\n")
	}

	fmt.Fprintln(cmdOut, " done")
	glog.V(4).Info("Successfully deleted kubeconfig entry")
	return nil
}
