package orm

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/oxodao/cao/config"

	_ "github.com/lib/pq"
)

var GET *ORM

type ORM struct {
	Db   *sqlx.DB
	Deck Deck
	Card Card
}

func Load() error {
	cfg := config.GET.Server.Database

	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
			cfg.Host,
			cfg.Port,
			cfg.Username,
			cfg.Password,
			cfg.Database,
		),
	)

	if err != nil {
		return err
	}

	GET = &ORM{
		db,
		Deck{db},
		Card{db},
	}

	return nil
}
