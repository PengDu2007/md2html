package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	VersionFlag bool
	cfgFile     string
	file string
)

var RootCmd = &cobra.Command{
	Use:   "md2html",
	Short: "Md2html is a tool that convert markdown to html.",
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		if len(file) == 0 {
			cmd.Help()
			return
		}
		fmt.Println(file)
	},
}

func init() {
	// cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().BoolVarP(&VersionFlag, "version", "v", false, "show version")
	RootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "markdown file")
}
