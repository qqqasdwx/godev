package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cliCmd)
}

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "into workspace docker sh",
	Long:  `docker-compose exec workspace sh`,
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command("docker-compose", "exec", "workspace", "sh")
		command.Dir = fmt.Sprintf("%s/%s", getCurrentAbsPath(), "workspace")
		command.Stderr = os.Stderr
		command.Stdout = os.Stdout
		command.Stdin = os.Stdin
		if err := command.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}
