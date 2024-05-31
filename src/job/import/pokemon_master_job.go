package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/MizukiShigi/go_pokemon/domain"
	"github.com/MizukiShigi/go_pokemon/infrastructure"
	"github.com/MizukiShigi/go_pokemon/util"
)

const logFilePath = "log/job/pokemon_master"
const startID = 1
const endID = 151
const pokemonApiUrl = "https://pokeapi.co/api/v2/pokemon"

// TODO: 複数の構造体に分割したい
type ApiPokemon struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Height         int `json:"height"`
	Weight         int `json:"weight"`
	BaseExperience int `json:"base_experience"`
	Sprites        struct {
		Other struct {
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
			} `json:"official-artwork"`
		} `json:"other"`
	} `json:"sprites"`
}

func init() {
	day := time.Now()
	logFile := "pokemon_master_job" + day.Format("20060102") + ".log"
	util.LoggingSetting(logFilePath + logFile)
}
 
func getPokemonApiById(id int, wg *sync.WaitGroup, results chan<- ApiPokemon) {
	defer wg.Done()

	url := fmt.Sprintf(pokemonApiUrl+"/%d", id)
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var apiPokemon ApiPokemon
	err = json.Unmarshal(body, &apiPokemon)
	if err != nil {
		log.Fatalln(err)
	}

	results <- apiPokemon
}

func main() {
	defer util.CloseLogFile()
	log.Println("start")

	// pokemon api呼び出し（リトライ3回）
	var wg sync.WaitGroup
	//   APIで取得したポケモンを格納するバッファ151個のチャネルを作成
	results := make(chan ApiPokemon, endID-startID+1)

	// TODO: error groupを使ってエラーハンドリングしたい
	for i := startID; i <= endID; i++ {
		wg.Add(1)
		go getPokemonApiById(i, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// DB登録
	infrastructure.ConnectDB()
	defer infrastructure.CloseDb()

	var insertPokemons []domain.Pokemon

	for result := range results {
		insertPokemon := domain.Pokemon{
			PokemonNumber:  result.ID,
			Name:           result.Name,
			Type1:          result.Types[0].Type.Name,
			Height:         float64(result.Height),
			Weight:         float64(result.Weight),
			BaseExperience: result.BaseExperience,
			ImageUrl:       result.Sprites.Other.OfficialArtwork.FrontDefault,
		}

		if len(result.Types) > 1 {
			insertPokemon.Type2 = result.Types[1].Type.Name
		}
		insertPokemons = append(insertPokemons, insertPokemon)
	}

	err := domain.InsertPokeomons(insertPokemons)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("end")
}
