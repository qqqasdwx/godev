package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"text/template"

	"github.com/spf13/cobra"
)

// Project Project
type Project struct {
	ProjectName string
}

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init project",
	Long:  `init project`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(path.Join(".", "workspace")); !os.IsNotExist(err) {
			fmt.Println("please remove your workspace dir!")
			return
		}
		if err := os.MkdirAll(path.Join(".", "workspace"), 0755); err != nil {
			fmt.Println(err)
			return
		}

		if err := tpl("file/bashrc", "./workspace/bashrc"); err != nil {
			fmt.Println(err)
			return
		}

		if err := tpl("file/docker-compose.yml", "./workspace/docker-compose.yml"); err != nil {
			fmt.Println(err)
			return
		}

		if err := tpl("file/Dockerfile", "./workspace/Dockerfile"); err != nil {
			fmt.Println(err)
			return
		}

		if err := tpl("file/env", "./workspace/.env"); err != nil {
			fmt.Println(err)
			return
		}

		if err := tpl("file/git-credentials", "./workspace/git-credentials"); err != nil {
			fmt.Println(err)
			return
		}

		if err := tpl("file/Makefile", "./Makefile"); err != nil {
			fmt.Println(err)
			return
		}
		if err := tpl("file/prod.Dockerfile", "./Dockerfile"); err != nil {
			fmt.Println(err)
			return
		}

		command := exec.Command("go", "mod", "init", fmt.Sprintf("git.secok.com/cad/%s", getCurrentPath()))
		command.Dir = getCurrentAbsPath()
		command.Stderr = os.Stderr
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func tpl(url string, filename string) error {
	var project Project
	project.ProjectName = getCurrentPath()

	t := template.Must(template.New("tpl").Parse(httpGet(url)))
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	if err = t.Execute(file, project); err != nil {
		return err
	}
	return nil
}
