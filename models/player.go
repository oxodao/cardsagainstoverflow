package models

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type User interface {
	GetLastPing() time.Time
	IsPlayer() bool
	Send()
}

type Player struct {
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"` // Not a "REAL" password, so no safety on this
	IsAdmin  bool   `db:"is_admin" json:"is_admin"`

	Room  *Room `db:"-" json:"-"`
	Score int   `db:"-" json:"score"`

	MutexWS    *sync.Mutex     `db:"-" json:"-"`
	Connection *websocket.Conn `db:"-" json:"-"`

	LastPing      time.Time `db:"-" json:"-"`
	RerollTimeout time.Time `db:"-" json:"reroll_timeout"`
	WizzTimeout   time.Time `db:"-" json:"wizz_timeout"`
}

func (p *Player) GetLastPing() time.Time {
	return p.LastPing
}

func (p *Player) IsPlayer() bool {
	return true
}

func (p *Player) Send(payload interface{}) {
	p.MutexWS.Lock()
	p.Connection.WriteJSON(payload)
	p.MutexWS.Unlock()
}

type Display struct {
	Room       *Room           `json:"-"`
	Connection *websocket.Conn `json:"-"`
	MutexWS    *sync.Mutex     `json:"-"`
	LastPing   time.Time       `json:"-"`
}

func (d *Display) GetLastPing() time.Time {
	return d.LastPing
}

func (d *Display) IsPlayer() bool {
	return false
}

func (d *Display) Send(payload interface{}) {
	d.MutexWS.Lock()
	d.Connection.WriteJSON(payload)
	d.MutexWS.Unlock()
}
