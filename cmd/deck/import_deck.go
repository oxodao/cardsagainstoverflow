package cmd_deck

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/oxodao/cao/models"
	"github.com/oxodao/cao/orm"
	"github.com/spf13/cobra"
)

var importDeckCmd = &cobra.Command{
	Use:   "import",
	Args:  cobra.ExactArgs(1),
	Short: "Import a deck from a deckfile",
	Run: func(cmd *cobra.Command, args []string) {
		selectedByDefault, err := cmd.Flags().GetBool("selected-by-default")
		if err != nil {
			fmt.Println("Failed to parse flag")
			fmt.Println("THIS SHOULD NOT HAVE HAPPENED, CONTACT THE DEVS !")
			fmt.Println(err)

			os.Exit(1)
		}

		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			fmt.Println("Usage: cao import-deck [deck_file.json]")
			fmt.Println("File not found")

			os.Exit(1)
		}

		data, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Println("Failed to read the file: ", err)

			os.Exit(1)
		}

		if err = orm.Load(); err != nil {
			fmt.Println("Failed to connect to the database: ", err)
			os.Exit(1)
		}

		var deck *models.Deck
		err = json.Unmarshal(data, &deck)

		if err != nil {
			fmt.Println("Invalid file specified")
			fmt.Println("Are you sure it's a deckfile ?")

			os.Exit(1)
		}

		if cmd.Flags().Changed("selected-by-default") {
			deck.SelectedByDefault = selectedByDefault
		}

		tx := orm.GET.Db.MustBeginTx(context.Background(), nil)

		fmt.Printf("Loading %v by %v\n", deck.Name, deck.Author)
		err = orm.GET.Deck.CreateDeck(tx, deck)
		if err != nil {
			fmt.Println("Failed to create deck: ", err)
			tx.Rollback()
			os.Exit(1)
		}

		fmt.Println("- Creating black cards")
		for _, card := range deck.BlackCards {
			err = orm.GET.Card.CreateCard(tx, deck, &card)
			if err != nil {
				fmt.Println("Failed to create card: ", err)
				tx.Rollback()
				os.Exit(1)
			}
		}

		fmt.Println("- Creating white cards")
		for _, card := range deck.WhiteCards {
			err = orm.GET.Card.CreateCard(tx, deck, &card)
			if err != nil {
				fmt.Println("Failed to create card: ", err)
				tx.Rollback()
				os.Exit(1)
			}
		}

		err = tx.Commit()
		if err != nil {
			fmt.Println("Failed to commit deck creation: ", err)
			os.Exit(1)
		}

		fmt.Println("Deck created")
	},
}

func registerImportCmd() {
	importDeckCmd.Flags().Bool("selected-by-default", false, "Select the deck by default in the web-ui")

	deckCmd.AddCommand(importDeckCmd)
}
