package domain

import (
	"time"
)

type Party struct {
	UserID    int       `json:"user_id"`
	PokemonID int       `json:"pokemon_id"`
	Nickname  string    `json:"nickname"`
	Level     int       `json:"level"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IPartyUsecase interface {
	AddParty(partyList []Party) error
}

type IPartyRepository interface {
	AddParty(partyList []Party) error
}
