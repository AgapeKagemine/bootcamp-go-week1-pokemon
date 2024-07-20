package repository

import (
	"errors"
	"fmt"
	"pokemon/internal/contract"
	"pokemon/internal/domain"
	"pokemon/internal/util"
)

type PlayerRepository interface {
	contract.GlobalInterface[domain.Player]
}

type PlayerRepositoryImpl struct {
	db map[int]domain.Player
}

func NewPlayerRepository() PlayerRepository {
	return PlayerRepositoryImpl{
		db: make(map[int]domain.Player, 0),
	}
}

func (repo PlayerRepositoryImpl) GetAll() (players []domain.Player, err error) {
	if util.IsEmpty(repo.db) {
		return nil, errors.New("no player found")
	}
	for _, Player := range repo.db {
		players = append(players, Player)
	}
	return players, err
}

func (repo PlayerRepositoryImpl) GetById(id int) (player domain.Player, err error) {
	if !util.IsExist(repo.db, id) {
		return domain.Player{}, errors.New("player not found")
	}
	return repo.db[id], err
}

func (repo PlayerRepositoryImpl) Save(player *domain.Player) (err error) {
	if util.IsExist(repo.db, player.ID) {
		return errors.New("player already exists")
	}
	player.ID = repo.db[len(repo.db)].ID + 1
	repo.db[player.ID] = *player
	fmt.Println("player saved successfully")
	return err
}

func (repo PlayerRepositoryImpl) Update(player *domain.Player) (err error) {
	if !util.IsExist(repo.db, player.ID) {
		return errors.New("player not found")
	}
	repo.db[player.ID] = *player
	fmt.Println("player updated successfully")
	return err
}

func (repo PlayerRepositoryImpl) DeleteById(id int) (err error) {
	if !util.IsExist(repo.db, id) {
		return errors.New("player not found")
	}
	delete(repo.db, id)
	fmt.Println("player deleted successfully")
	return err
}
