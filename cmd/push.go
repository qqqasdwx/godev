package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var pushTag string

func init() {
	pushCmd.Flags().StringVarP(&pushTag, "tag", "t", getGitCurrentTag(), "git tag")
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push docker images and push git tag",
	Long:  `make push && git checkout master && git push origin [tag] && git push`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("Building version: %s\nDocker images tag: %s", pushTag, getDockerTag(pushTag)))

		// make push
		command1 := exec.Command("make", "push", "-e", fmt.Sprintf("\"DOCKER_TAG_NAME=%s\"", getDockerTag(pushTag)))
		command1.Dir = getCurrentAbsPath()
		command1.Stderr = os.Stderr
		command1.Stdout = os.Stdout
		if err := command1.Run(); err != nil {
			fmt.Println(err)
			os.Exit(0)
			return
		}

		command2 := exec.Command("git", "checkout", "master")
		command2.Dir = getCurrentAbsPath()
		command2.Stderr = os.Stderr
		command2.Stdout = os.Stdout
		if err := command2.Run(); err != nil {
			fmt.Println(err)
			os.Exit(0)
			return
		}

		command3 := exec.Command("git", "push", "origin", pushTag)
		command3.Dir = getCurrentAbsPath()
		command3.Stderr = os.Stderr
		command3.Stdout = os.Stdout
		if err := command3.Run(); err != nil {
			fmt.Println(err)
			os.Exit(0)
			return
		}

		command4 := exec.Command("git", "push", "origin", pushTag)
		command4.Dir = getCurrentAbsPath()
		command4.Stderr = os.Stderr
		command4.Stdout = os.Stdout
		if err := command4.Run(); err != nil {
			fmt.Println(err)
			os.Exit(0)
			return
		}

	},
}
