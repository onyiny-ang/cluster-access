package util

import "k8s.io/client-go/tools/clientcmd"

func GetClientConfig(pathOptions *clientcmd.PathOptions, context, kubeconfigPath string) clientcmd.ClientConfig {
	loadingRules := *pathOptions.LoadingRules
	loadingRules.Precedence = pathOptions.GetLoadingPrecedence()
	loadingRules.ExplicitPath = kubeconfigPath
	overrides := &clientcmd.ConfigOverrides{
		CurrentContext: context,
	}

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(&loadingRules, overrides)
}
