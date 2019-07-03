package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(upCmd)
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "workspace docker up",
	Long:  `docker-compose up -d`,
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command("docker-compose", "up", "-d")
		command.Dir = fmt.Sprintf("%s/%s", getCurrentAbsPath(), "workspace")
		command.Stderr = os.Stderr
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}
