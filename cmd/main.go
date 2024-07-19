package main

import (
	"fmt"
	"pokemon/internal/domain"
	"pokemon/internal/handler"
	"pokemon/internal/repository"
	"pokemon/internal/usecase"
	"time"
)

func autowired() (handler.PokemonHandlerInterface, handler.PlayerHandlerInterface) {
	return handler.NewPokemonHandler(
		usecase.NewPokemonUseCase(repository.NewPokemonRepository())), handler.NewPlayerHandler(usecase.NewPlayerUseCase(repository.NewPlayerRepository()))
}

func main() {
	hPokemon, hPlayer := autowired()

	x1 := domain.Pokemon{
		ID:             1,
		Name:           "Pikachu",
		Type:           "Electric",
		CatchRate:      10,
		IsRare:         true,
		RegisteredDate: time.Now().Format(time.UnixDate),
	}

	x2 := domain.Pokemon{
		ID:             2,
		Name:           "Pikachu",
		Type:           "Electric",
		CatchRate:      10,
		IsRare:         true,
		RegisteredDate: time.Now().Format(time.UnixDate),
	}

	x3 := domain.Pokemon{
		ID:             3,
		Name:           "Pikachu",
		Type:           "Electric",
		CatchRate:      10,
		IsRare:         true,
		RegisteredDate: time.Now().Format(time.UnixDate),
	}

	hPokemon.Save(&x1)
	hPokemon.Save(&x2)
	hPokemon.Save(&x3)

	list, _ := hPokemon.GetAll()
	for _, v := range list {
		fmt.Println(v)
	}

	fmt.Println()

	p1 := domain.Player{
		ID:   1,
		Name: "Player 1",
		Bag:  []domain.Pokemon{x1, x2},
	}

	p2 := domain.Player{
		ID:   2,
		Name: "Player 2",
		Bag:  []domain.Pokemon{x3},
	}

	p3 := domain.Player{
		ID:   3,
		Name: "Player 3",
		Bag:  []domain.Pokemon{},
	}

	hPlayer.Save(&p1)
	hPlayer.Save(&p2)
	hPlayer.Save(&p3)

	list2, _ := hPlayer.GetAll()
	for _, v := range list2 {
		fmt.Println(v)
	}
}
