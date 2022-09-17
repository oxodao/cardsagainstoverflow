package cmd_deck

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/oxodao/cao/orm"
	"github.com/spf13/cobra"
)

var setDefaultCmd = &cobra.Command{
	Use:   "set-default [id] [true|false]",
	Short: "Set whether a deck should be selected by default in the web-ui",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(2)(cmd, args); err != nil {
			return err
		}

		if _, err := strconv.ParseInt(args[0], 10, 64); err != nil {
			return errors.New("deck id invalid")
		}

		if _, err := strconv.ParseBool(args[1]); err != nil {
			return errors.New("enabled by default value can only be true or false")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := orm.Load(); err != nil {
			fmt.Println("Failed to connect to the database: ", err)
			os.Exit(1)
		}

		deckId, _ := strconv.ParseInt(args[0], 10, 64)
		enable, _ := strconv.ParseBool(args[1])

		err := orm.GET.Deck.SetSelectedByDefault(deckId, enable)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No deck with this id")
			} else {
				fmt.Println("Failed to set the status of the deck: ", err)
			}

			os.Exit(1)
		}

		if enable {
			fmt.Println("Deck will be selected by default !")
		} else {
			fmt.Println("Deck will not be selected by default !")
		}
	},
}
