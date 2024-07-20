package repository

import (
	"errors"
	"fmt"
	"pokemon/internal/contract"
	"pokemon/internal/domain"
	"pokemon/internal/util"
)

type PokemonRepository interface {
	contract.GlobalInterface[domain.Pokemon]
}

type PokemonRepositoryImpl struct {
	db map[int]domain.Pokemon
}

func NewPokemonRepository() PokemonRepository {
	return PokemonRepositoryImpl{
		db: make(map[int]domain.Pokemon, 0),
	}
}

func (repo PokemonRepositoryImpl) GetAll() (pokemons []domain.Pokemon, err error) {
	if util.IsEmpty(repo.db) {
		return nil, errors.New("no pokemon found")
	}
	for _, pokemon := range repo.db {
		pokemons = append(pokemons, pokemon)
	}
	return pokemons, err
}

func (repo PokemonRepositoryImpl) GetById(id int) (pokemon domain.Pokemon, err error) {
	if !util.IsExist(repo.db, id) {
		return domain.Pokemon{}, errors.New("pokemon not found")
	}
	return repo.db[id], err
}

func (repo PokemonRepositoryImpl) Save(pokemon *domain.Pokemon) (err error) {
	if _, exists := repo.db[pokemon.ID]; exists {
		return errors.New("pokemon already exists")
	}
	pokemon.ID = repo.db[len(repo.db)].ID + 1
	repo.db[pokemon.ID] = *pokemon
	fmt.Println("pokemon saved successfully")
	return err
}

func (repo PokemonRepositoryImpl) Update(pokemon *domain.Pokemon) (err error) {
	if !util.IsExist(repo.db, pokemon.ID) {
		return errors.New("pokemon not found")
	}
	repo.db[pokemon.ID] = *pokemon
	fmt.Println("pokemon updated successfully")
	return err
}

func (repo PokemonRepositoryImpl) DeleteById(id int) (err error) {
	if !util.IsExist(repo.db, id) {
		return errors.New("pokemon not found")
	}
	delete(repo.db, id)
	fmt.Println("pokemon deleted successfully")
	return err
}
