package cmd

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// Verbose Verbose
var Verbose bool

var rootCmd = &cobra.Command{
	Use:   "godev",
	Short: "godev is a golang develop tool",
	Long:  `Complete documentation is available at https://github.com/qqqasdwx/godev`,
	Args:  cobra.ExactValidArgs(1),
	ValidArgs: []string{
		"version",
		"completion",
		"help",
		"up",
		"down",
		"cli",
		"develop",
		"dev",
		"merge",
		"get",
		"common",
		"run",
		"make",
		"push",
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// Execute Execute
func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getCurrentAbsPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return path.Base(dir)
}

func httpGet(path string) string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(fmt.Sprintf("https://raw.githubusercontent.com/qqqasdwx/godev/intwall/%s", path))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bodyBytes)
}

func getGitCurrentTag() string {
	command1 := exec.Command("git", "rev-list", "--tags", "--max-count=1")
	command1.Dir = getCurrentAbsPath()
	command1.Stderr = os.Stderr
	out, err := command1.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	if len(out) == 0 {
		return "1.0.0"
	}

	command2 := exec.Command("git", "describe", "--tags", strings.Trim(string(out), "\n"))
	command2.Dir = getCurrentAbsPath()
	command2.Stderr = os.Stderr
	tag, err := command2.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	return strings.Trim(string(tag), "\n")
}

func getDockerTag(tag string) string {
	tags := strings.Split(tag, ".")
	if len(tags) != 3 {
		return tag
	}
	return fmt.Sprintf("%s.%s", tags[0], tags[1])
}

func getNextTag(currentTag string) (string, error) {
	tags := strings.Split(currentTag, ".")
	if len(tags) != 3 {
		return currentTag, nil
	}
	smalltag, err := strconv.Atoi(tags[2])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s.%d", tags[0], tags[1], smalltag+1), nil
}
