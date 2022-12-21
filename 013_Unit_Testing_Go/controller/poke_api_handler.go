package controller

import (
	"catching-pokemons/models"
	"catching-pokemons/util"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	ErrPokemonNotFound = errors.New("pokemon not found")
	ErrPokeApiError    = errors.New("api error, try again later")
)

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	apiPokemon, err := GetPokemonFromPokeApi(id)
	if err != nil {
		if errors.Is(err, ErrPokemonNotFound) {
			respondwithJSON(w, http.StatusNotFound, fmt.Sprintf("Pokemon not found with id: %s", id))
		}
		respondwithJSON(w, http.StatusInternalServerError, fmt.Sprintf("error while calling external API: %s", err.Error()))
	}

	parsedPokemon, err := util.ParsePokemon(apiPokemon)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, fmt.Sprintf("error found: %s", err.Error()))
	}

	respondwithJSON(w, http.StatusOK, parsedPokemon)
}

func GetPokemonFromPokeApi(id string) (apiPokemon models.PokeApiPokemonResponse, err error) {
	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	response, err := http.Get(request)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusNotFound {
			err = ErrPokemonNotFound
			return
		}
		err = ErrPokeApiError
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &apiPokemon)
	if err != nil {
		return
	}

	return
}
