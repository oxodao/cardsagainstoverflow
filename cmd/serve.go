package cmd

import (
	"fmt"
	"net/http"

	"github.com/oxodao/cao/routes"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the main game server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cards Against Overflow - Rewritten")

		r := mux.NewRouter()

		api := r.PathPrefix("/api").Subrouter()
		api.HandleFunc("/decks", routes.GetDecks)
		api.HandleFunc("/decks/{deck_id}", routes.GetFullDeck)
		api.HandleFunc("/join/{room_id}", routes.JoinRoom)
		api.HandleFunc("/join/{room_id}/display", routes.JoinRoomAsDisplay)

		fmt.Println(
			"Failed to start the server: ",
			http.ListenAndServe("0.0.0.0:1234", r),
		)
	},
}
