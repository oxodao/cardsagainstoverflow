package cmd_deck

import (
	"fmt"
	"os"

	"github.com/oxodao/cao/orm"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available decks",
	Run: func(cmd *cobra.Command, args []string) {
		if err := orm.Load(); err != nil {
			fmt.Println("Failed to connect to the database: ", err)
			os.Exit(1)
		}

		decks, err := orm.GET.Deck.List()
		if err != nil {
			fmt.Println("Failed to get deck list: ", err)
			os.Exit(1)
		}

		fmt.Println("Decks with a star are selected by default")
		fmt.Println()
		fmt.Println("== Available decks ==")
		for _, deck := range decks {
			selected := ""
			if deck.SelectedByDefault {
				selected = "*"
			}

			fmt.Printf("\t- [%v%v] %v by %v (%v black / %v white)\n", selected, deck.Id, deck.Name, deck.Author, deck.AmtBlackCards, deck.AmtWhiteCards)
		}
	},
}
