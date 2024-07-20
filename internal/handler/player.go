package handler

import (
	"errors"
	"pokemon/internal/contract"
	"pokemon/internal/domain"
	"pokemon/internal/usecase"
	"pokemon/internal/util"
)

type PlayerHandler interface {
	contract.GlobalInterface[domain.Player]
}

type PlayerHandlerImpl struct {
	playerHandler usecase.PlayerUseCase
}

func NewPlayerHandler(uc usecase.PlayerUseCase) PlayerHandler {
	return PlayerHandlerImpl{
		playerHandler: uc,
	}
}

func (uc PlayerHandlerImpl) GetAll() (players []domain.Player, err error) {
	if !util.IsConnected(uc.playerHandler) {
		return nil, errors.New("player usecase is not initialized")
	}
	return uc.playerHandler.GetAll()
}

func (uc PlayerHandlerImpl) GetById(id int) (player domain.Player, err error) {
	if !util.IsConnected(uc.playerHandler) {
		return domain.Player{}, errors.New("player usecase is not initialized")
	}
	return uc.playerHandler.GetById(id)
}

func (uc PlayerHandlerImpl) Save(player *domain.Player) (err error) {
	if !util.IsConnected(uc.playerHandler) {
		return errors.New("player usecase is not initialized")
	}
	return uc.playerHandler.Save(player)
}

func (uc PlayerHandlerImpl) Update(player *domain.Player) error {
	if !util.IsConnected(uc.playerHandler) {
		return errors.New("player usecase is not initialized")
	}
	return uc.playerHandler.Update(player)
}

func (uc PlayerHandlerImpl) DeleteById(id int) error {
	if !util.IsConnected(uc.playerHandler) {
		return errors.New("player usecase is not initialized")
	}
	return uc.playerHandler.DeleteById(id)
}
