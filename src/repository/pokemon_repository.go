package repository

import (
	"database/sql"
	"log"

	"github.com/MizukiShigi/go_pokemon/domain"
)

type PokemonRepository struct {
	db *sql.DB
}

func NewPokemonRepository(db *sql.DB) domain.IPokemonRepository {
	return &PokemonRepository{db}
}

func (pr *PokemonRepository) InsertPokeomons(pokemons []domain.Pokemon) (err error) {
	cmd := `
		INSERT INTO pokemons
			(pokemon_number, name, type1, type2, height, weight, base_experience, image_url, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
	`

	tx, err := pr.db.Begin()
	if err != nil {
		log.Fatalln(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	for _, pokemon := range pokemons {
		_, err := stmt.Exec(
			pokemon.PokemonNumber,
			pokemon.Name,
			pokemon.Type1,
			pokemon.Type2,
			pokemon.Height,
			pokemon.Weight,
			pokemon.BaseExperience,
			pokemon.ImageUrl)
		if err != nil {
			log.Fatalln(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (pr *PokemonRepository) GetPokemonByPokemonNumber(pokemonNumber int) (domain.Pokemon, error) {
	var pokemon domain.Pokemon
	cmd := "SELECT pokemon_number, name, type1, type2, height, weight, base_experience, image_url FROM pokemons WHERE pokemon_number = ? LIMIT 1"
	err := pr.db.QueryRow(cmd, pokemonNumber).Scan(
		&pokemon.PokemonNumber,
		&pokemon.Name,
		&pokemon.Type1,
		&pokemon.Type2,
		&pokemon.Height,
		&pokemon.Weight,
		&pokemon.BaseExperience,
		&pokemon.ImageUrl,
	)
	if err != nil {
		log.Fatalln(err)
		return domain.Pokemon{}, err
	}

	return pokemon, err
}
