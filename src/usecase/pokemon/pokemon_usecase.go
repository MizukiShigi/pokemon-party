package pokemon

import (
	"github.com/MizukiShigi/go_pokemon/domain"
)

type PokemonUsecase struct {
	pr domain.IPokemonRepository
}

func NewPokemonUsecase(pr domain.IPokemonRepository) domain.IPokemonUsecase {
	return &PokemonUsecase{pr}
}

func (pu *PokemonUsecase) GetPokemon(pokemon *domain.Pokemon) error {
	p, err := pu.pr.GetPokemonByPokemonNumber(pokemon.PokemonNumber)
	if err != nil {
		return err
	}
	*pokemon = p
	return nil
}
