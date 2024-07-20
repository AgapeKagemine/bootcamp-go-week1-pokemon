package handler

import (
	"errors"
	"pokemon/internal/domain"
	"pokemon/internal/interfaces"
	"pokemon/internal/usecase"
	"pokemon/internal/util"
)

type PlayerHandlerInterface interface {
	interfaces.GlobalInterface[domain.Player]
}

type PlayerHandler struct {
	playerHandler usecase.PlayerUseCaseInterface
}

func NewPlayerHandler(uc usecase.PlayerUseCaseInterface) PlayerHandlerInterface {
	if handler, ok := uc.(usecase.PlayerUseCase); ok {
		return PlayerHandler{
			playerHandler: handler,
		}
	}
	return PlayerHandler{
		playerHandler: nil,
	}
}

func (uc PlayerHandler) GetAll() (players []domain.Player, err error) {
	if !util.IsConnected(uc.playerHandler) {
		return nil, errors.New("player usecase is not initialized")
	}
	return uc.playerHandler.GetAll()
}

func (uc PlayerHandler) GetById(id int) (player domain.Player, err error) {
	if !util.IsConnected(uc.playerHandler) {
		return domain.Player{}, errors.New("player usecase is not initialized")
	}
	return uc.playerHandler.GetById(id)
}

func (uc PlayerHandler) Save(player *domain.Player) (err error) {
	if !util.IsConnected(uc.playerHandler) {
		return errors.New("player usecase is not initialized")
	}
	return uc.playerHandler.Save(player)
}

func (uc PlayerHandler) UpdateById(player *domain.Player) error {
	if !util.IsConnected(uc.playerHandler) {
		return errors.New("player usecase is not initialized")
	}
	return uc.playerHandler.UpdateById(player)
}

func (uc PlayerHandler) DeleteById(id int) error {
	if !util.IsConnected(uc.playerHandler) {
		return errors.New("player usecase is not initialized")
	}
	return uc.playerHandler.DeleteById(id)
}
