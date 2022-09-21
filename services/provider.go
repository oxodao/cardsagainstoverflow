package services

import (
	"strings"

	"github.com/oxodao/cao/models"
)

var GET *Provider

type Provider struct {
	Rooms []models.Room
}

func (p *Provider) GetRoom(code string) *models.Room {
	for _, r := range p.Rooms {
		if strings.EqualFold(r.Code, code) {
			return &r
		}
	}

	return nil
}

func (p *Provider) CreateRoom() *models.Room {
	r := models.NewRoom()
	p.Rooms = append(p.Rooms, *r)

	return r
}

func Load() error {
	GET = &Provider{}

	return nil
}
