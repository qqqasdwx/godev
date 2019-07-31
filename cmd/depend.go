package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(commonCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get [repo]",
	Long:  `go get [repo] && go mod vendor -v`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command1 := exec.Command("docker-compose", "exec", "workspace", "/bin/sh", "-c", fmt.Sprintf("\"go get -v %s\"", args[0]))
		command1.Dir = fmt.Sprintf("%s/%s", getCurrentAbsPath(), "workspace")
		command1.Stderr = os.Stderr
		command1.Stdout = os.Stdout
		if err := command1.Run(); err != nil {
			fmt.Println(err)
			os.Exit(0)
			return
		}
		command2 := exec.Command("docker-compose", "exec", "workspace", "/bin/sh", "-c", fmt.Sprintf("\"go mod vendor -v\""))
		command2.Dir = fmt.Sprintf("%s/%s", getCurrentAbsPath(), "workspace")
		command2.Stderr = os.Stderr
		command2.Stdout = os.Stdout
		if err := command2.Run(); err != nil {
			fmt.Println(err)
			os.Exit(0)
			return
		}
		command3 := exec.Command("pkill", "gopls")
		command3.Dir = getCurrentAbsPath()
		command3.Stderr = os.Stderr
		command3.Stdout = os.Stdout
		if err := command3.Run(); err != nil {
			fmt.Println(err)
			os.Exit(0)
			return
		}
	},
}

var commonCmd = &cobra.Command{
	Use:   "common",
	Short: "get common",
	Long:  `get common && go mod vendor -v`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command1 := exec.Command("docker-compose", "exec", "workspace", "/bin/sh", "-c", fmt.Sprintf("\"go get -v %s\"", "git.secok.com/cad/common"))
		command1.Dir = fmt.Sprintf("%s/%s", getCurrentAbsPath(), "workspace")
		command1.Stderr = os.Stderr
		command1.Stdout = os.Stdout
		if err := command1.Run(); err != nil {
			fmt.Println(err)
			os.Exit(0)
			return
		}
		command2 := exec.Command("docker-compose", "exec", "workspace", "/bin/sh", "-c", fmt.Sprintf("\"go mod vendor -v\""))
		command2.Dir = fmt.Sprintf("%s/%s", getCurrentAbsPath(), "workspace")
		command2.Stderr = os.Stderr
		command2.Stdout = os.Stdout
		if err := command2.Run(); err != nil {
			fmt.Println(err)
			os.Exit(0)
			return
		}
		command3 := exec.Command("pkill", "gopls")
		command3.Dir = getCurrentAbsPath()
		command3.Stderr = os.Stderr
		command3.Stdout = os.Stdout
		if err := command3.Run(); err != nil {
			fmt.Println(err)
			os.Exit(0)
			return
		}
	},
}
