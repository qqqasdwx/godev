package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var makeTag string
var dockerfile string

func init() {
	makeCmd.Flags().StringVarP(&dockerfile, "dockerfile", "f", "./Dockerfile", "Dockerfile path")
	makeCmd.Flags().StringVarP(&makeTag, "tag", "t", getGitCurrentTag(), "git tag")
	rootCmd.AddCommand(makeCmd)
}

var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "build docker images",
	Long:  `docker build -f ./Dockerfile -t [tag] .`,
	Run: func(cmd *cobra.Command, args []string) {
		nextTag, err := getNextTag(makeTag)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
			return
		}
		fmt.Println(fmt.Sprintf("Building version: %s\nDocker images tag: %s", nextTag, getDockerTag(makeTag)))

		command1 := exec.Command("make", "-e", fmt.Sprintf("\"DOCKERFILE_PATH=%s\"", dockerfile), "-e", fmt.Sprintf("\"DOCKER_TAG_NAME=%s\"", getDockerTag(makeTag)))
		command1.Dir = getCurrentAbsPath()
		command1.Stderr = os.Stderr
		command1.Stdout = os.Stdout
		if err := command1.Run(); err != nil {
			fmt.Println(err)
			os.Exit(0)
			return
		}

		command2 := exec.Command("git", "tag", "-a", nextTag, "-m", fmt.Sprintf("Released version %s", nextTag))
		command2.Dir = getCurrentAbsPath()
		command2.Stderr = os.Stderr
		command2.Stdout = os.Stdout
		if err := command2.Run(); err != nil {
			fmt.Println(err)
		}
	},
}
