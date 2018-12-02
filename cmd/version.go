package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var version = "v0.1"

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number of md2html",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func UserAgent() string {
	return fmt.Sprintf("md2html/%s (%s; %s; %s)", version, runtime.GOOS, runtime.GOARCH, runtime.Version())
}