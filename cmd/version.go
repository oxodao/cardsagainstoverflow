package cmd

import (
	"fmt"

	"github.com/oxodao/cao/utils"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version informations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cards Against Overflow - Rewritten")
		fmt.Printf("Version %v by %v\n", utils.VERSION, utils.AUTHOR)
	},
}
