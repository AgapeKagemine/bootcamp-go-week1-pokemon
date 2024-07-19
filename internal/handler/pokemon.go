package handler

import (
	"pokemon/internal/domain"
	"pokemon/internal/interfaces"
	"pokemon/internal/usecase"
)

type PokemonHandlerInterface interface {
	interfaces.InterfaceRepository[domain.Pokemon]
}

type PokemonHandler struct {
	pokemonHandler usecase.PokemonUseCaseInterface
}

func NewPokemonHandler(usecase usecase.PokemonUseCaseInterface) PokemonHandlerInterface {
	return PokemonHandler{
		pokemonHandler: usecase,
	}
}

func (uc PokemonHandler) GetAll() (pokemons []domain.Pokemon, err error) {
	return uc.pokemonHandler.GetAll()
}

func (uc PokemonHandler) GetById(id int) (pokemon domain.Pokemon, err error) {
	return uc.pokemonHandler.GetById(id)
}

func (uc PokemonHandler) Save(pokemon *domain.Pokemon) (err error) {
	return uc.pokemonHandler.Save(pokemon)
}

func (uc PokemonHandler) UpdateById(pokemon *domain.Pokemon) error {
	return uc.pokemonHandler.UpdateById(pokemon)
}

func (uc PokemonHandler) DeleteById(id int) error {
	return uc.pokemonHandler.DeleteById(id)
}
