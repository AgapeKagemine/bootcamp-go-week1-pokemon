package usecase

import (
	"errors"
	"pokemon/internal/domain"
	"pokemon/internal/interfaces"
	"pokemon/internal/repository"
	"pokemon/internal/util"
)

type PokemonUseCaseInterface interface {
	interfaces.GlobalInterface[domain.Pokemon]
}

type PokemonUseCase struct {
	pokemonRepo repository.PokemonRepositoryInterface
}

func NewPokemonUseCase(repo repository.PokemonRepositoryInterface) PokemonUseCaseInterface {
	if rp, ok := repo.(repository.PokemonRepository); ok {
		return PokemonUseCase{
			pokemonRepo: rp,
		}
	}
	return PokemonUseCase{
		pokemonRepo: nil,
	}
}

func (uc PokemonUseCase) GetAll() (pokemons []domain.Pokemon, err error) {
	if !util.IsConnected(uc.pokemonRepo) {
		return nil, errors.New("pokemon repository is not initialized")
	}
	return uc.pokemonRepo.GetAll()
}

func (uc PokemonUseCase) GetById(id int) (pokemon domain.Pokemon, err error) {
	if !util.IsConnected(uc.pokemonRepo) {
		return domain.Pokemon{}, errors.New("pokemon repository is not initialized")
	}
	return uc.pokemonRepo.GetById(id)
}

func (uc PokemonUseCase) Save(pokemon *domain.Pokemon) (err error) {
	if !util.IsConnected(uc.pokemonRepo) {
		return errors.New("pokemon repository is not initialized")
	}
	return uc.pokemonRepo.Save(pokemon)
}

func (uc PokemonUseCase) UpdateById(pokemon *domain.Pokemon) error {
	if !util.IsConnected(uc.pokemonRepo) {
		return errors.New("pokemon repository is not initialized")
	}
	return uc.pokemonRepo.UpdateById(pokemon)
}

func (uc PokemonUseCase) DeleteById(id int) error {
	if !util.IsConnected(uc.pokemonRepo) {
		return errors.New("pokemon repository is not initialized")
	}
	return uc.pokemonRepo.DeleteById(id)
}
