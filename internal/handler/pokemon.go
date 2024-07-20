package handler

import (
	"pokemon/internal/contract"
	"pokemon/internal/domain"
	"pokemon/internal/usecase"
)

type PokemonHandler interface {
	contract.GlobalInterface[domain.Pokemon]
}

type PokemonHandlerImpl struct {
	pokemonHandler usecase.PokemonUseCase
}

func NewPokemonHandler(uc usecase.PokemonUseCase) PokemonHandler {
	return PokemonHandlerImpl{
		pokemonHandler: uc,
	}
}

func (uc PokemonHandlerImpl) GetAll() (pokemons []domain.Pokemon, err error) {
	return uc.pokemonHandler.GetAll()
}

func (uc PokemonHandlerImpl) GetById(id int) (pokemon domain.Pokemon, err error) {
	return uc.pokemonHandler.GetById(id)
}

func (uc PokemonHandlerImpl) Save(pokemon *domain.Pokemon) (err error) {
	return uc.pokemonHandler.Save(pokemon)
}

func (uc PokemonHandlerImpl) Update(pokemon *domain.Pokemon) error {
	return uc.pokemonHandler.Update(pokemon)
}

func (uc PokemonHandlerImpl) DeleteById(id int) error {
	return uc.pokemonHandler.DeleteById(id)
}
