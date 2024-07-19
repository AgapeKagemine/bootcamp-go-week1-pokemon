package handler

import (
	"pokemon/internal/domain"
	"pokemon/internal/interfaces"
	"pokemon/internal/usecase"
)

type PlayerHandlerInterface interface {
	interfaces.InterfaceRepository[domain.Player]
}

type PlayerHandler struct {
	playerHandler usecase.PlayerUseCaseInterface
}

func NewPlayerHandler(usecase usecase.PlayerUseCaseInterface) PlayerHandlerInterface {
	return PlayerHandler{
		playerHandler: usecase,
	}
}

func (uc PlayerHandler) GetAll() (players []domain.Player, err error) {
	return uc.playerHandler.GetAll()
}

func (uc PlayerHandler) GetById(id int) (player domain.Player, err error) {
	return uc.playerHandler.GetById(id)
}

func (uc PlayerHandler) Save(player *domain.Player) (err error) {
	return uc.playerHandler.Save(player)
}

func (uc PlayerHandler) UpdateById(player *domain.Player) error {
	return uc.playerHandler.UpdateById(player)
}

func (uc PlayerHandler) DeleteById(id int) error {
	return uc.playerHandler.DeleteById(id)
}
