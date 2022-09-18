package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/oxodao/cao/config"
	"github.com/oxodao/cao/orm"
	"github.com/oxodao/cao/routes"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the main game server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := orm.Load(); err != nil {
			fmt.Println("Failed to connect to the database: ", err)
			os.Exit(1)
		}

		fmt.Println("Cards Against Overflow - Rewritten")
		r := mux.NewRouter()

		api := r.PathPrefix("/api").Subrouter()
		api.HandleFunc("/decks", routes.GetDecks)
		api.HandleFunc("/decks/{deck_id}", routes.GetFullDeck)
		api.HandleFunc("/join/{room_id}", routes.JoinRoom)
		api.HandleFunc("/join/{room_id}/display", routes.JoinRoomAsDisplay)

		fmt.Printf("Listening on %v:%v\n", config.GET.Server.Web.Host, config.GET.Server.Web.Port)

		fmt.Println(
			"Failed to start the server: ",
			http.ListenAndServe(
				fmt.Sprintf("%v:%v", config.GET.Server.Web.Host, config.GET.Server.Web.Port),
				r,
			),
		)
	},
}
