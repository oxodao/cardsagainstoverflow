package cmd_deck

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/oxodao/cao/orm"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a given deck",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := orm.Load(); err != nil {
			fmt.Println("Failed to connect to the database: ", err)
			os.Exit(1)
		}

		deckId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Printf("Failed to find deck id %v\n", args[0])
			os.Exit(1)
		}

		deck, err := orm.GET.Deck.FindById(deckId)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No deck with this id")
			} else {
				fmt.Println("Failed to query the database: ", err)
			}

			os.Exit(1)
		}

		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			fmt.Println("Failed to get flag: ", err)
			fmt.Println("THIS SHOULD NOT HAPPEND. REPORT TO THE DEV !")
			os.Exit(1)
		}

		if !force {
			fmt.Println("Removing a deck will also remove the cards and the game history in which it has been played")
			fmt.Println("If you are sure, use the --force / -f flag")
			os.Exit(0)
		}

		err = orm.GET.Deck.Delete(deck)
		if err != nil {
			fmt.Println("Failed to delete the deck: ", err)
			os.Exit(1)
		}

		fmt.Printf("Deck %v by %v deleted\n", deck.Name, deck.Author)
	},
}

func registerDeleteCmd() {
	deleteCmd.Flags().BoolP("force", "f", false, "Force the deletion")

	deckCmd.AddCommand(deleteCmd)
}
