package util

import (
	"catching-pokemons/models"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestParsePokemon(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		body, err := ioutil.ReadFile("samples/pokeapi_response.json")
		assert.NoError(t, err)

		var response models.PokeApiPokemonResponse

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		result, err := ParsePokemon(response)
		assert.NoError(t, err)

		apiResponse, err := ioutil.ReadFile("samples/api_response.json")
		assert.NoError(t, err)

		var expected models.Pokemon

		err = json.Unmarshal(apiResponse, &expected)
		assert.NoError(t, err)

		assert.Equal(t, expected, result)
	})

	t.Run("TypeNotFound", func(t *testing.T) {
		body, err := ioutil.ReadFile("samples/pokeapi_response.json")
		assert.NoError(t, err)

		var response models.PokeApiPokemonResponse

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		response.PokemonType = []models.PokemonType{}

		_, err = ParsePokemon(response)
		assert.Error(t, err)
		assert.Errorf(t, err, ErrNotFoundPokemonType.Error())
	})

	t.Run("TypeNameNotFound", func(t *testing.T) {
		body, err := ioutil.ReadFile("samples/pokeapi_response.json")
		assert.NoError(t, err)

		var response models.PokeApiPokemonResponse

		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		response.PokemonType[0].RefType.Name = ""

		_, err = ParsePokemon(response)
		assert.Error(t, err)
		assert.Errorf(t, err, ErrNotFoundPokemonTypeName.Error())
	})
}
