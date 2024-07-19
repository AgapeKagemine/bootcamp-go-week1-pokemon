package repository

import (
	"errors"
	"fmt"
	"pokemon/internal/domain"
	"pokemon/internal/interfaces"
)

type PokemonRepositoryInterface interface {
	interfaces.InterfaceRepository[domain.Pokemon]
}

type PokemonRepository struct {
	db map[int]domain.Pokemon
}

func NewPokemonRepository() PokemonRepositoryInterface {
	return PokemonRepository{
		db: make(map[int]domain.Pokemon, 0),
	}
}

func (repo PokemonRepository) GetAll() (pokemons []domain.Pokemon, err error) {
	if len(repo.db) < 1 {
		return nil, errors.New("no pokemons found")
	}
	for _, pokemon := range repo.db {
		pokemons = append(pokemons, pokemon)
	}
	return pokemons, err
}

func (repo PokemonRepository) GetById(id int) (pokemon domain.Pokemon, err error) {
	if _, exists := repo.db[pokemon.ID]; !exists {
		return domain.Pokemon{}, errors.New("pokemon not found")
	}
	return repo.db[pokemon.ID], err
}

func (repo PokemonRepository) Save(pokemon *domain.Pokemon) (err error) {
	if _, exists := repo.db[pokemon.ID]; exists {
		return errors.New("pokemon already exists")
	}
	repo.db[pokemon.ID] = *pokemon
	fmt.Println("pokemon saved successfully")
	return err
}

func (repo PokemonRepository) UpdateById(pokemon *domain.Pokemon) (err error) {
	if _, exists := repo.db[pokemon.ID]; !exists {
		return errors.New("pokemon not found")
	}
	repo.db[pokemon.ID] = *pokemon
	fmt.Println("pokemon updated successfully")
	return err
}

func (repo PokemonRepository) DeleteById(id int) (err error) {
	if _, exists := repo.db[id]; !exists {
		return errors.New("pokemon not found")
	}
	delete(repo.db, id)
	return err
}
