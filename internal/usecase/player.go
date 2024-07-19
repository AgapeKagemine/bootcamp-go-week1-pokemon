package usecase

import (
	"pokemon/internal/domain"
	"pokemon/internal/interfaces"
	"pokemon/internal/repository"
)

type PlayerUseCaseInterface interface {
	interfaces.InterfaceRepository[domain.Player]
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
	return nil
}

func (uc PlayerUseCase) GetAll() (players []domain.Player, err error) {
	return uc.playerRepo.GetAll()
}

func (uc PlayerUseCase) GetById(id int) (Player domain.Player, err error) {
	return uc.playerRepo.GetById(id)
}

func (uc PlayerUseCase) Save(Player *domain.Player) (err error) {
	return uc.playerRepo.Save(Player)
}

func (uc PlayerUseCase) UpdateById(Player *domain.Player) error {
	return uc.playerRepo.UpdateById(Player)
}

func (uc PlayerUseCase) DeleteById(id int) error {
	return uc.playerRepo.DeleteById(id)
}
