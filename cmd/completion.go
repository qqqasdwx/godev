package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// var completionOut string
var completionShell string

func init() {
	// completionCmd.Flags().StringVarP(&completionOut, "out", "o", "", "Output directory to write to")
	completionCmd.Flags().StringVarP(&completionShell, "shell", "s", os.Getenv("SHELL"), "Your $SHELL, default read it from system environment $SHELL")
	completionCmd.MarkFlagFilename("shell")
	rootCmd.AddCommand(completionCmd)
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generating bash/zsh completions",
	Long:  `Generating bash/zsh completions`,
	Run: func(cmd *cobra.Command, args []string) {
		oss := strings.Split(completionShell, "/")
		switch oss[len(oss)-1] {
		case "zsh":
			fmt.Println("your shell is zsh")
			if err := rootCmd.GenZshCompletionFile("/usr/share/zsh/functions/Completion/_godev"); err != nil {
				log.Fatal(err)
			}
			fmt.Println(`Generating success~ if completions can't be used, you can run this command: "compinit"`)
		case "bash":
			fmt.Println("your shell is bash")
			if err := rootCmd.GenBashCompletionFile("/etc/bash_completion.d/godev"); err != nil {
				log.Fatal(err)
			}
			command := exec.Command("bash", `/etc/bash_completion`)
			if err := command.Run(); err != nil {
				fmt.Printf("%#v", err)
				log.Fatal(err)
			}
			fmt.Println(`Generating success~ if completions can't be used, you can run this command: ". /etc/bash_completion"`)
		default:
			fmt.Println("This Shell is not supported")
		}
	},
}
