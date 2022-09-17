package orm

import (
	"github.com/jmoiron/sqlx"
	"github.com/oxodao/cao/models"
)

type Card struct {
	db *sqlx.DB
}

func (c *Card) CreateCard(tx *sqlx.Tx, deck *models.Deck, card *models.Card) error {
	row := tx.QueryRow(`
		INSERT INTO card (deck_id, text, is_black_card)
		VALUES ($1, $2, $3)
		RETURNING id
	`, deck.Id, card.Text, card.IsBlackCard)

	if row.Err() != nil {
		return row.Err()
	}

	return row.Scan(&card.Id)
}
