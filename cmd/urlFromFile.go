package cmd

import (
	"os"

	"github.com/gkwa/lightvan/core"
	"github.com/spf13/cobra"
)

var urlFromFileCmd = &cobra.Command{
	Use:     "url-from-file",
	Aliases: []string{"uff"},
	Short:   "Extract URL from file path",
	Long:    `Extract url from file path by looping over file finding first regex match`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Usage()
			if err != nil {
				cmd.PrintErrf("Error: %v\n", err)
			}
			cmd.PrintErrln("Error: File path is required")
			os.Exit(1)
		}

		provider := &core.FileURLProvider{
			Path: args[0],
		}

		err := core.ExtractURL(cmd.Context(), provider)
		if err != nil {
			cmd.PrintErrf("Error: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(urlFromFileCmd)
}
