package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gbfapi",
	Short: "Gran Blue Fantasy API",
	Long: `
		An API server and crawler for Gran Blue Fantasy to provide developer can develop helper APP/Webasite
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Empty command, show help to tell user what to do
		cmd.Help()
	},
}
