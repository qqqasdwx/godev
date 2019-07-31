package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run with bee",
	Long:  `docker-compose exec workspace /bin/sh -c "GOFLAGS=-mod=vendor bee run`,
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command("docker-compose", "exec", "workspace", "/bin/sh", "-c", fmt.Sprintf("GOFLAGS=-mod=vendor bee run"))
		command.Dir = fmt.Sprintf("%s/%s", getCurrentAbsPath(), "workspace")
		command.Stdin = os.Stdin
		command.Stderr = os.Stderr
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			fmt.Println(err)
		}
	},
}
