package usecase

import (
	"errors"
	"pokemon/internal/contract"
	"pokemon/internal/domain"
	"pokemon/internal/repository"
	"pokemon/internal/util"
)

type PlayerUseCase interface {
	contract.GlobalInterface[domain.Player]
}

type PlayerUseCaseImpl struct {
	playerRepo repository.PlayerRepository
}

func NewPlayerUseCase(repo repository.PlayerRepository) PlayerUseCaseImpl {
	return PlayerUseCaseImpl{
		playerRepo: repo,
	}
}

func (uc PlayerUseCaseImpl) GetAll() (players []domain.Player, err error) {
	if !util.IsConnected(uc.playerRepo) {
		return nil, errors.New("player repository is not initialized")
	}
	return uc.playerRepo.GetAll()
}

func (uc PlayerUseCaseImpl) GetById(id int) (player domain.Player, err error) {
	if !util.IsConnected(uc.playerRepo) {
		return domain.Player{}, errors.New("player repository is not initialized")
	}
	return uc.playerRepo.GetById(id)
}

func (uc PlayerUseCaseImpl) Save(player *domain.Player) (err error) {
	if !util.IsConnected(uc.playerRepo) {
		return errors.New("player repository is not initialized")
	}
	return uc.playerRepo.Save(player)
}

func (uc PlayerUseCaseImpl) Update(player *domain.Player) (err error) {
	if !util.IsConnected(uc.playerRepo) {
		return errors.New("player repository is not initialized")
	}
	return uc.playerRepo.Update(player)
}

func (uc PlayerUseCaseImpl) DeleteById(id int) (err error) {
	if !util.IsConnected(uc.playerRepo) {
		return errors.New("player repository is not initialized")
	}
	return uc.playerRepo.DeleteById(id)
}
