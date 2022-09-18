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

func (c *Card) GetCards(deckId int64, isBlackCard bool) ([]models.Card, error) {
	rows, err := c.db.Queryx(`
		SELECT id, text, is_black_card
		FROM card
		WHERE deck_id = $1
		  AND is_black_card = $2
	`, deckId, isBlackCard)

	if err != nil {
		return nil, err
	}

	cards := []models.Card{}
	for rows.Next() {
		card := models.Card{}
		err = rows.StructScan(&card)
		if err != nil {
			return nil, err
		}

		cards = append(cards, card)
	}

	return cards, nil
}

func (c *Card) FillDeck(deck *models.Deck) error {
	blackCards, err := c.GetCards(deck.Id, true)
	if err != nil {
		return err
	}

	whiteCards, err := c.GetCards(deck.Id, false)
	if err != nil {
		return err
	}

	deck.BlackCards = blackCards
	deck.WhiteCards = whiteCards

	return nil
}
