package cmd

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

// Verbose Verbose
var Verbose bool

var rootCmd = &cobra.Command{
	Use:       "godev",
	Short:     "godev is a golang develop tool",
	Long:      `Complete documentation is available at https://github.com/qqqasdwx/godev`,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"version", "completion", "help", "up", "down", "cli"},
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

	resp, err := client.Get(fmt.Sprintf("https://raw.githubusercontent.com/qqqasdwx/godev/master/%s", path))
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
