package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oxodao/cao/orm"
)

func GetDecks(w http.ResponseWriter, r *http.Request) {
	decks, err := orm.GET.Deck.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Err getting deck list: ", err)
		return
	}

	data, _ := json.Marshal(decks)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func GetFullDeck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if _, ok := vars["deck_id"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var id int64 = -1
	var err error = nil
	if id, err = strconv.ParseInt(vars["deck_id"], 10, 64); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Failed to parse deck_id %v: %v\n", vars["deck_id"], err)
		return
	}

	deck, err := orm.GET.Deck.FindById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error getting the detail of a deck: ", err)
		return
	}

	err = orm.GET.Card.FillDeck(deck)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Failed to fill deck: ", err)
		return
	}

	data, _ := json.Marshal(deck)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
