package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(developCmd)
	rootCmd.AddCommand(devCmd)
	rootCmd.AddCommand(mergeCmd)
}

var developCmd = &cobra.Command{
	Use:   "develop",
	Short: "create develop branch from master",
	Long:  `git checkout -b develop master`,
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command("git", "checkout", "-b", "develop", "master")
		command.Dir = getCurrentAbsPath()
		command.Stderr = os.Stderr
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			fmt.Println(err)
		}
	},
}

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "checkout develop branch",
	Long:  `git checkout develop`,
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command("git", "checkout", "develop")
		command.Dir = getCurrentAbsPath()
		command.Stderr = os.Stderr
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			fmt.Println(err)
		}
	},
}

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "merge from develop to master",
	Long:  `git checkout master && git merge --no-ff --no-edit develop && git branch -d develop`,
	Run: func(cmd *cobra.Command, args []string) {
		command1 := exec.Command("git", "checkout", "master")
		command1.Dir = getCurrentAbsPath()
		command1.Stderr = os.Stderr
		command1.Stdout = os.Stdout
		if err := command1.Run(); err != nil {
			fmt.Println(err)
		}

		command2 := exec.Command("git", "merge", "--no-ff", "--no-edit", "develop")
		command2.Dir = getCurrentAbsPath()
		command2.Stderr = os.Stderr
		command2.Stdout = os.Stdout
		if err := command2.Run(); err != nil {
			fmt.Println(err)
		}

		command3 := exec.Command("git", "branch", "-d", "develop")
		command3.Dir = getCurrentAbsPath()
		command3.Stderr = os.Stderr
		command3.Stdout = os.Stdout
		if err := command3.Run(); err != nil {
			fmt.Println(err)
		}
	},
}
