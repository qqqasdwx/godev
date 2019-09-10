package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var ciTag string

func init() {
	ciCmd.Flags().StringVarP(&ciTag, "tag", "t", getGitCurrentTag(), "git tag")
	rootCmd.AddCommand(ciCmd)
}

var ciCmd = &cobra.Command{
	Use:   "ci",
	Short: "push code",
	Long:  `develop: push, master: tag and push`,
	Run: func(cmd *cobra.Command, args []string) {
		command1 := exec.Command("git", "symbolic-ref", "HEAD")
		command1.Dir = getCurrentAbsPath()
		command1.Stderr = os.Stderr
		out, err := command1.Output()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		currentBranch := strings.TrimPrefix(strings.Trim(string(out), "\n"), "refs/heads/")
		if currentBranch == "master" {
			nextTag, err := getNextTag(ciTag)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
				return
			}
			fmt.Println(fmt.Sprintf("Building version: %s\n", nextTag))
			command2 := exec.Command("git", "tag", "-a", nextTag, "-m", fmt.Sprintf("Released version %s", nextTag))
			command2.Dir = getCurrentAbsPath()
			command2.Stderr = os.Stderr
			command2.Stdout = os.Stdout
			if err := command2.Run(); err != nil {
				fmt.Println(err)
			}

			command3 := exec.Command("git", "push", "origin", nextTag)
			command3.Dir = getCurrentAbsPath()
			command3.Stderr = os.Stderr
			command3.Stdout = os.Stdout
			if err := command3.Run(); err != nil {
				fmt.Println(err)
				os.Exit(0)
				return
			}

			command4 := exec.Command("git", "push", "origin", "master")
			command4.Dir = getCurrentAbsPath()
			command4.Stderr = os.Stderr
			command4.Stdout = os.Stdout
			if err := command4.Run(); err != nil {
				fmt.Println(err)
				os.Exit(0)
				return
			}
		} else {
			command4 := exec.Command("git", "push", "origin", currentBranch)
			command4.Dir = getCurrentAbsPath()
			command4.Stderr = os.Stderr
			command4.Stdout = os.Stdout
			if err := command4.Run(); err != nil {
				fmt.Println(err)
				os.Exit(0)
				return
			}
		}

	},
}
