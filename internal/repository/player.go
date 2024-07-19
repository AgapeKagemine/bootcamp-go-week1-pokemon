package repository

import (
	"errors"
	"fmt"
	"pokemon/internal/domain"
	"pokemon/internal/interfaces"
)

type PlayerRepositoryInterface interface {
	interfaces.InterfaceRepository[domain.Player]
}

type PlayerRepository struct {
	db map[int]domain.Player
}

func NewPlayerRepository() PlayerRepositoryInterface {
	return PlayerRepository{
		db: make(map[int]domain.Player, 0),
	}
}

func (repo PlayerRepository) GetAll() (players []domain.Player, err error) {
	if len(repo.db) < 1 {
		return nil, errors.New("no Players found")
	}
	for _, Player := range repo.db {
		players = append(players, Player)
	}
	return players, err
}

func (repo PlayerRepository) GetById(id int) (player domain.Player, err error) {
	if _, exists := repo.db[player.ID]; !exists {
		return domain.Player{}, errors.New("player not found")
	}
	return repo.db[player.ID], err
}

func (repo PlayerRepository) Save(player *domain.Player) (err error) {
	if _, exists := repo.db[player.ID]; exists {
		return errors.New("player already exists")
	}
	repo.db[player.ID] = *player
	fmt.Println("player saved successfully")
	return err
}

func (repo PlayerRepository) UpdateById(player *domain.Player) (err error) {
	if _, exists := repo.db[player.ID]; !exists {
		return errors.New("player not found")
	}
	repo.db[player.ID] = *player
	fmt.Println("Player updated successfully")
	return err
}

func (repo PlayerRepository) DeleteById(id int) (err error) {
	if _, exists := repo.db[id]; !exists {
		return errors.New("player not found")
	}
	delete(repo.db, id)
	return err
}
