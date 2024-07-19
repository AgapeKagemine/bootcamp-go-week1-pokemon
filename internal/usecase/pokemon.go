package usecase

import (
	"pokemon/internal/domain"
	"pokemon/internal/interfaces"
	"pokemon/internal/repository"
)

type PokemonUseCaseInterface interface {
	interfaces.InterfaceRepository[domain.Pokemon]
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
	return PokemonUseCase{}
}

func (uc PokemonUseCase) GetAll() (pokemons []domain.Pokemon, err error) {
	return uc.pokemonRepo.GetAll()
}

func (uc PokemonUseCase) GetById(id int) (pokemon domain.Pokemon, err error) {
	return uc.pokemonRepo.GetById(id)
}

func (uc PokemonUseCase) Save(pokemon *domain.Pokemon) (err error) {
	return uc.pokemonRepo.Save(pokemon)
}

func (uc PokemonUseCase) UpdateById(pokemon *domain.Pokemon) error {
	return uc.pokemonRepo.UpdateById(pokemon)
}

func (uc PokemonUseCase) DeleteById(id int) error {
	return uc.pokemonRepo.DeleteById(id)
}
