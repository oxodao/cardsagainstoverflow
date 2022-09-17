package cmd

import (
	cmd_deck "github.com/oxodao/cao/cmd/deck"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cao",
	Short: "Cards Against Overflow",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cmd_deck.Register(rootCmd)

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(versionCmd)
}
