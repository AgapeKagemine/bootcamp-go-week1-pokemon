package usecase

import (
	"errors"
	"pokemon/internal/domain"
	"pokemon/internal/interfaces"
	"pokemon/internal/repository"
	"pokemon/internal/util"
)

type PlayerUseCaseInterface interface {
	interfaces.GlobalInterface[domain.Player]
}

type PlayerUseCase struct {
	playerRepo repository.PlayerRepositoryInterface
}

func NewPlayerUseCase(repo repository.PlayerRepositoryInterface) PlayerUseCaseInterface {
	if rp, ok := repo.(repository.PlayerRepository); ok {
		return PlayerUseCase{
			playerRepo: rp,
		}
	}
	return PlayerUseCase{
		playerRepo: nil,
	}
}

func (uc PlayerUseCase) GetAll() (players []domain.Player, err error) {
	if !util.IsConnected(uc.playerRepo) {
		return nil, errors.New("player repository is not initialized")
	}
	return uc.playerRepo.GetAll()
}

func (uc PlayerUseCase) GetById(id int) (Player domain.Player, err error) {
	if !util.IsConnected(uc.playerRepo) {
		return domain.Player{}, errors.New("player repository is not initialized")
	}
	return uc.playerRepo.GetById(id)
}

func (uc PlayerUseCase) Save(Player *domain.Player) (err error) {
	if !util.IsConnected(uc.playerRepo) {
		return errors.New("player repository is not initialized")
	}
	return uc.playerRepo.Save(Player)
}

func (uc PlayerUseCase) UpdateById(Player *domain.Player) error {
	if !util.IsConnected(uc.playerRepo) {
		return errors.New("player repository is not initialized")
	}
	return uc.playerRepo.UpdateById(Player)
}

func (uc PlayerUseCase) DeleteById(id int) error {
	if !util.IsConnected(uc.playerRepo) {
		return errors.New("player repository is not initialized")
	}
	return uc.playerRepo.DeleteById(id)
}
