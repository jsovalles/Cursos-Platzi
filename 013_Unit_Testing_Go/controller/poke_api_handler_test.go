package controller

import (
	"catching-pokemons/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPokemonFromPokeApi(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		pokemon, err := GetPokemonFromPokeApi("132")
		assert.NoError(t, err)

		body, err := ioutil.ReadFile("samples/poke_api_readed.json")
		assert.NoError(t, err)

		var expected models.PokeApiPokemonResponse

		err = json.Unmarshal(body, &expected)
		assert.NoError(t, err)

		assert.Equal(t, expected, pokemon)
	})

	t.Run("SuccessWithMock", func(t *testing.T) {
		id := "ditto"
		request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

		body, err := ioutil.ReadFile("samples/pokeapi_response.json")
		assert.NoError(t, err)
		var expected models.PokeApiPokemonResponse
		err = json.Unmarshal(body, &expected)
		assert.NoError(t, err)

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", request, httpmock.NewBytesResponder(200, body))

		pokemon, err := GetPokemonFromPokeApi(id)
		assert.NoError(t, err)
		assert.Equal(t, expected, pokemon)
	})

	t.Run("InternalServerError", func(t *testing.T) {

		id := "ditto"
		request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", request, httpmock.NewBytesResponder(500, nil))

		_, err := GetPokemonFromPokeApi(id)
		assert.Errorf(t, err, ErrPokeApiError.Error())
	})

	t.Run("NotFoundError", func(t *testing.T) {

		id := "ditto"
		request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", request, httpmock.NewBytesResponder(404, nil))

		_, err := GetPokemonFromPokeApi(id)
		assert.Errorf(t, err, ErrPokemonNotFound.Error())
	})
}

func TestGetPokemon(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		r, err := http.NewRequest("GET", "/pokemon/{id}", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		vars := map[string]string{
			"id": "ditto",
		}

		r = mux.SetURLVars(r, vars)

		GetPokemon(w, r)

		body, err := ioutil.ReadFile("samples/api_response.json")
		assert.NoError(t, err)
		var expected models.Pokemon
		err = json.Unmarshal(body, &expected)
		assert.NoError(t, err)

		var actual models.Pokemon
		err = json.Unmarshal(w.Body.Bytes(), &actual)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expected, actual)
	})

	t.Run("PokemonNotFoundError", func(t *testing.T) {
		r, err := http.NewRequest("GET", "/pokemon/{id}", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		vars := map[string]string{
			"id": "ssssssss",
		}

		r = mux.SetURLVars(r, vars)

		GetPokemon(w, r)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("PokemonInternalServerError", func(t *testing.T) {
		r, err := http.NewRequest("GET", "/pokemon/{id}", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		vars := map[string]string{
			"id": "",
		}

		r = mux.SetURLVars(r, vars)

		GetPokemon(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
