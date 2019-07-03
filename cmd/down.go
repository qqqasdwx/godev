package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(downCmd)
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "workspace docker down",
	Long:  `docker-compose down`,
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command("docker-compose", "down")
		command.Dir = fmt.Sprintf("%s/%s", getCurrentAbsPath(), "workspace")
		command.Stderr = os.Stderr
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}
