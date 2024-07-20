package main

import (
	"pokemon/internal/handler"
	"pokemon/internal/repository"
	"pokemon/internal/usecase"
	"pokemon/internal/util"
)

func init() {
	util.ConsoleClear()
}

func autowired() (handler.PokemonHandlerInterface, handler.PlayerHandlerInterface) {
	return handler.NewPokemonHandler(
			usecase.NewPokemonUseCase(
				repository.NewPokemonRepository(),
			),
		),
		handler.NewPlayerHandler(
			usecase.NewPlayerUseCase(
				repository.NewPlayerRepository(),
			),
		)
}

func main() {
	util.ConsoleClear()
}
