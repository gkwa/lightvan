package cmd

import (
	"github.com/gkwa/lightvan/core"
	"github.com/spf13/cobra"
)

var urlFromClipboardCmd = &cobra.Command{
	Use:     "url-from-clipboard",
	Aliases: []string{"ufc"},
	Short:   "Extract URL components from clipboard",
	Long:    `Extracts and displays components of a URL stored in the clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		provider := core.ClipboardURLProvider{}
		err := core.ExtractURL(cmd.Context(), provider)
		if err != nil {
			cmd.PrintErrf("Error: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(urlFromClipboardCmd)
}
