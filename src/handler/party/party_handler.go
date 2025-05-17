package party

import (
	"net/http"

	"github.com/MizukiShigi/go_pokemon/domain"
	"github.com/labstack/echo/v4"
)

type IPartyHandler interface {
	AddParty(c echo.Context) error
}

type PartyHandler struct {
	pu domain.IPartyUsecase
}

type PartyRequest struct {
	UserID      int            `json:"user_id"`
	PokemonList []PartyPokemon `json:"pokemons"`
}

type PartyPokemon struct {
	PokemonID int    `json:"pokemon_id"`
	Nickname  string `json:"nickname"`
	Level     int    `json:"level"`
}

type PartyResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
}

func NewPartyHandler(pu domain.IPartyUsecase) IPartyHandler {
	return &PartyHandler{pu}
}

func NewPartyResponse() *PartyResponse {
	return &PartyResponse{"ok", 200}
}

func (ph *PartyHandler) AddParty(c echo.Context) error {
	var partyRequest PartyRequest
	if err := c.Bind(&partyRequest); err != nil {
		return err
	}

	var partyList []domain.Party
	for _, pokemon := range partyRequest.PokemonList {
		party := domain.Party{
			UserID:    partyRequest.UserID,
			PokemonID: pokemon.PokemonID,
			Nickname:  pokemon.Nickname,
			Level:     pokemon.Level,
		}
		partyList = append(partyList, party)
	}
	err := ph.pu.AddParty(partyList)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, NewPartyResponse())
}
