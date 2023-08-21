package models

import (
	"log"
	"time"
)

type Pokemon struct {
	Id int
	PokemonNumber int
    Name string
    Type1 string
    Type2 string
    Height float64
    Weight float64
    BaseExperience int
    ImageUrl string
    CreatedAt time.Time
    UpdatedAt time.Time
}

func InsertPokeomons(pokemons []Pokemon) (err error) {
	cmd := `
		INSERT INTO pokemons
			(pokemon_number, name, type1, type2, height, weight, base_experience, image_url, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
	`

	tx, err := Db.Begin()
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

func GetPokemons() (pokemons []Pokemon, err error) {
	cmd := "SELECT * FROM pokemons"
	rows, err := Db.Query(cmd)
	log.Println(rows)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pokemon Pokemon
		err := rows.Scan(
			&pokemon.Id,
			&pokemon.PokemonNumber,
			&pokemon.Name,
			&pokemon.Type1,
			&pokemon.Type2,
			&pokemon.Height,
			&pokemon.Weight,
			&pokemon.BaseExperience,
			&pokemon.ImageUrl,
			&pokemon.CreatedAt,
			&pokemon.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
		}
		pokemons = append(pokemons, pokemon)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	return pokemons, err
}