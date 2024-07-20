package repository

import (
	"errors"
	"fmt"
	"pokemon/internal/domain"
	"pokemon/internal/interfaces"
	"pokemon/internal/util"
)

type PlayerRepositoryInterface interface {
	interfaces.GlobalInterface[domain.Player]
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
	if util.IsEmpty(repo.db) {
		return nil, errors.New("no player found")
	}
	for _, Player := range repo.db {
		players = append(players, Player)
	}
	return players, err
}

func (repo PlayerRepository) GetById(id int) (player domain.Player, err error) {
	if !util.IsExist(repo.db, id) {
		return domain.Player{}, errors.New("player not found")
	}
	return repo.db[player.ID], err
}

func (repo PlayerRepository) Save(player *domain.Player) (err error) {
	if util.IsExist(repo.db, player.ID) {
		return errors.New("player already exists")
	}
	repo.db[player.ID] = *player
	fmt.Println("player saved successfully")
	return err
}

func (repo PlayerRepository) UpdateById(player *domain.Player) (err error) {
	if !util.IsExist(repo.db, player.ID) {
		return errors.New("player not found")
	}
	repo.db[player.ID] = *player
	fmt.Println("player updated successfully")
	return err
}

func (repo PlayerRepository) DeleteById(id int) (err error) {
	if !util.IsExist(repo.db, id) {
		return errors.New("player not found")
	}
	delete(repo.db, id)
	fmt.Println("player deleted successfully")
	return err
}
