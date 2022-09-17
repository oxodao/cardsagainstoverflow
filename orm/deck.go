package orm

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/oxodao/cao/models"
)

type Deck struct {
	db *sqlx.DB
}

func (d *Deck) CreateDeck(tx *sqlx.Tx, deck *models.Deck) error {
	row := tx.QueryRow(`
			INSERT INTO deck (name, author, selected_by_default)
			VALUES ($1, $2, $3)
			RETURNING id
		`, deck.Name, deck.Author, deck.SelectedByDefault)

	if row.Err() != nil {
		return row.Err()
	}

	return row.Scan(&deck.Id)
}

func (d *Deck) List() ([]models.Deck, error) {
	rows, err := d.db.Queryx(`
		SELECT 
		 	deck.id,
			deck.name,
			deck.author,
			COUNT(card.id) FILTER (WHERE card.is_black_card = FALSE) as amt_white_cards,
			COUNT(card.id) FILTER (WHERE card.is_black_card = TRUE) as amt_black_cards,
			deck.selected_by_default
		FROM deck
		LEFT JOIN card ON card.deck_id = deck.id
		GROUP BY deck.id, deck.name, deck.author, deck.selected_by_default
		ORDER BY deck.id
	`)

	if err != nil {
		return nil, err
	}

	decks := []models.Deck{}
	for rows.Next() {
		deck := models.Deck{}
		err = rows.StructScan(&deck)
		if err != nil {
			return nil, err
		}

		decks = append(decks, deck)
	}

	return decks, nil
}

func (d *Deck) FindById(id int64) (*models.Deck, error) {
	row := d.db.QueryRowx(`
		SELECT 
			deck.id,
			deck.name,
			deck.author,
			COUNT(card.id) FILTER (WHERE card.is_black_card = FALSE) as amt_white_cards,
			COUNT(card.id) FILTER (WHERE card.is_black_card = TRUE) as amt_black_cards,
			deck.selected_by_default
		FROM deck
		LEFT JOIN card ON card.deck_id = deck.id
		WHERE deck.id = $1
		GROUP BY deck.id, deck.name, deck.author, deck.selected_by_default
	`, id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var deck models.Deck
	err := row.StructScan(&deck)

	return &deck, err
}

func (d *Deck) Delete(deck *models.Deck) error {
	_, err := d.db.Exec(`DELETE FROM deck WHERE id = $1`, deck.Id)
	return err
}

func (d *Deck) SetSelectedByDefault(deckId int64, selectedByDefault bool) error {
	resp, err := d.db.Exec(`UPDATE deck SET selected_by_default = $1 WHERE id = $2`, selectedByDefault, deckId)
	if err != nil {
		return err
	}

	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
