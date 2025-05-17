package repository

import (
	"database/sql"
	"github.com/MizukiShigi/go_pokemon/domain"
)

type PartyRepository struct {
	db *sql.DB
}

func NewPartyRepository(db *sql.DB) domain.IPartyRepository {
	return &PartyRepository{db}
}

func (pr *PartyRepository) AddParty(partyList []domain.Party) error {
	cmd := `INSERT INTO parties
		(user_id, pokemon_id, nickname, level, created_at, updated_at)
		VALUES (?, ?, ?, ?, now(), now())
	`
	tx, err := pr.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, party := range partyList {
		_, err := tx.Exec(cmd, party.UserID, party.PokemonID, party.Nickname, party.Level)
		if err != nil {
			return err
		}
	}
	tx.Commit()

	return nil
}