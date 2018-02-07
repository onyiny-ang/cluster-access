package delete


import "fmt"


func NewCmdDelete(cmdOut io.Writer)*cobra.Command {

    cmd := &cobra.Command{
        Use: "delete",
        Short: "deletes a specified cluster from KUBECONFIG (requires -c)",
        Long: "Deletes an entry for the specified cluster-registry cluster from KUBECONFIG (requires -c)",
    }

    initCmd := &cobra.Command{
        Use: ""
    }


}
