package models

type CardUsage struct {
	RemainingCards []Card `json:"-"`
	AvailableCards []Card `json:"-"`
}

type RoomConfig struct {
	Password      string `json:"-"`
	ZenMode       bool   `json:"zen_mode"`
	RerollTimeout int    `json:"reroll_timeout"`
	MaxTurns      int    `json:"max_turns"`
	TimePerTurn   int    `json:"time_per_turn"`
	Decks         []Deck `json:"-"`
	DecksJson     []int  `json:"decks"`
}

type RoomState struct {
	Started       bool `json:"started"`
	Turn          int  `json:"turn"`
	Countdown     int  `json:"countdown"`
	BlackCard     Card `json:"-"`
	BlackCardJson int  `json:"black_card"`

	BlackCardUsage CardUsage `json:"-"`
	WhiteCardUsage CardUsage `json:"-"`
}

type Room struct {
	Code         string     `db:"code" json:"code"`
	Participants []User     `db:"-" json:"participants"`
	Config       RoomConfig `db:"-" json:"config"`
	State        RoomState  `db:"-" json:"state"`
}

func NewRoom() *Room {
	// r := Room{}
	return nil
}
