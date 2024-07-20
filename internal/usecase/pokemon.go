package usecase

import (
	"errors"
	"pokemon/internal/contract"
	"pokemon/internal/domain"
	"pokemon/internal/repository"
	"pokemon/internal/util"
)

type PokemonUseCase interface {
	contract.GlobalInterface[domain.Pokemon]
}

type PokemonUseCaseImpl struct {
	pokemonRepo repository.PokemonRepository
}

func NewPokemonUseCase(repo repository.PokemonRepository) PokemonUseCase {
	return PokemonUseCaseImpl{
		pokemonRepo: repo,
	}
}

func (uc PokemonUseCaseImpl) GetAll() (pokemons []domain.Pokemon, err error) {
	if !util.IsConnected(uc.pokemonRepo) {
		return nil, errors.New("pokemon repository is not initialized")
	}
	return uc.pokemonRepo.GetAll()
}

func (uc PokemonUseCaseImpl) GetById(id int) (pokemon domain.Pokemon, err error) {
	if !util.IsConnected(uc.pokemonRepo) {
		return domain.Pokemon{}, errors.New("pokemon repository is not initialized")
	}
	return uc.pokemonRepo.GetById(id)
}

func (uc PokemonUseCaseImpl) Save(pokemon *domain.Pokemon) (err error) {
	if !util.IsConnected(uc.pokemonRepo) {
		return errors.New("pokemon repository is not initialized")
	}
	return uc.pokemonRepo.Save(pokemon)
}

func (uc PokemonUseCaseImpl) Update(pokemon *domain.Pokemon) error {
	if !util.IsConnected(uc.pokemonRepo) {
		return errors.New("pokemon repository is not initialized")
	}
	return uc.pokemonRepo.Update(pokemon)
}

func (uc PokemonUseCaseImpl) DeleteById(id int) error {
	if !util.IsConnected(uc.pokemonRepo) {
		return errors.New("pokemon repository is not initialized")
	}
	return uc.pokemonRepo.DeleteById(id)
}
