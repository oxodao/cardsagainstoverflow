package models

type Deck struct {
	Id                int64  `db:"id" json:"id"`
	Name              string `db:"name" json:"name"`
	Author            string `db:"author" json:"author"`
	SelectedByDefault bool   `db:"selected_by_default" json:"selected_by_default"`
	WhiteCards        []Card `db:"-" json:"white_cards,omitempty"`
	BlackCards        []Card `db:"-" json:"black_cards,omitempty"`

	AmtWhiteCards int `db:"amt_white_cards" json:"amt_white_cards"`
	AmtBlackCards int `db:"amt_black_cards" json:"amt_black_cards"`
}

type Card struct {
	Id          int64  `db:"id" json:"id"`
	Text        string `db:"text" json:"text"`
	IsBlackCard bool   `db:"is_black_card" json:"is_black_card"`
}
