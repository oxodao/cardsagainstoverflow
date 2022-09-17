package cmd_deck

import "github.com/spf13/cobra"

var deckCmd = &cobra.Command{
	Use:   "deck",
	Short: "Commands related to the decks",
}

func Register(root *cobra.Command) {
	deckCmd.AddCommand(listCmd)
	deckCmd.AddCommand(setDefaultCmd)
	registerImportCmd()
	registerDeleteCmd()

	root.AddCommand(deckCmd)
}
