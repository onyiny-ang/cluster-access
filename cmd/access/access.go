package main

import (
	"fmt"
	"os"

	"k8s.io/apiserver/pkg/util/logs"
	"k8s.io/cluster-access/pkg/access"
)

func main() {

	logs.InitLogs()
	defer logs.FlushLogs()

	err := access.NewClusterAccessCommand(os.Stdout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
