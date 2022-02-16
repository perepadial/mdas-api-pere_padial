package user

import (
	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

type AddFavoritePokemon struct {
	Repository domain.FavoritePokemonRepository
}

func (useCase AddFavoritePokemon) Execute(userId string, pokemonId string) error {
	user := useCase.Repository.FindUser(domain.CreateUserId(userId))
	favoriteId := domain.CreatePokemonId(pokemonId)
	user.AddFavorite(favoriteId)
	error := useCase.Repository.Save(*user)

	if error != nil {
		return error
	}

	return nil
}
