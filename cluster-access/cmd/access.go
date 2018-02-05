package main

import (
	"fmt"
	"os"

	"cluster-access/pkg/access"

	"k8s.io/apiserver/pkg/util/logs"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	err := access.NewClusterAccessCommand().Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
