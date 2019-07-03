package cmd

import (
	"fmt"

	"github.com/qqqasdwx/godev/build"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of godev",
	Long:  `All software has versions. This is godev's`,
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			fmt.Printf("%#v\n", build.Data())
			return
		}
		fmt.Println(build.String())
	},
}
