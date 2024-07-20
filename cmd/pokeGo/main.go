package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"pokemon/internal/domain"
	"pokemon/internal/handler"
	"pokemon/internal/repository"
	"pokemon/internal/usecase"
	"pokemon/internal/util"
	"time"
)

var trash string

func autowired() (handler.PlayerHandler, handler.PokemonHandler) {
	return handler.NewPlayerHandler(
			usecase.NewPlayerUseCase(
				repository.NewPlayerRepository(),
			),
		),
		handler.NewPokemonHandler(
			usecase.NewPokemonUseCase(
				repository.NewPokemonRepository(),
			),
		)
}

func fillPokemons(hPokemons handler.PokemonHandler) {
	pokemons := []domain.Pokemon{
		{Name: "Pikachu", Type: "Electric", CatchRate: 0.7, IsRare: false, RegisteredDate: time.Now().Local().Format(time.DateOnly)},
		{Name: "Charmander", Type: "Fire", CatchRate: 0.4, IsRare: true, RegisteredDate: time.Now().Local().AddDate(0, -1, 0).Format(time.DateOnly)},
		{Name: "Bulbasaur", Type: "Grass/Poison", CatchRate: 0.1, IsRare: true, RegisteredDate: time.Now().Local().AddDate(0, -6, 0).Format(time.DateOnly)},
		{Name: "Dragonite", Type: "Dragon/Flying", CatchRate: 0.3, IsRare: true, RegisteredDate: time.Now().Local().AddDate(0, -8, 0).Format(time.DateOnly)},
		{Name: "Mew", Type: "Psychic", CatchRate: 0.01, IsRare: true, RegisteredDate: time.Now().Local().AddDate(0, -10, 0).Format(time.DateOnly)},
	}
	for _, v := range pokemons {
		hPokemons.Save(&v)
	}
}

func pokeGo(player domain.Player, players handler.PlayerHandler, pokemons handler.PokemonHandler, scanner *bufio.Scanner) {
	for {
		util.ConsoleClear()
		fmt.Println("welcome to Poke Go, ", player.Name)
		if len(player.Bag) < 1 {
			fmt.Println("you have no pokemon")
		} else {
			fmt.Println("your pokemon: ")
			for _, pokemon := range player.Bag {
				fmt.Printf("%s %s\n", pokemon.Name, pokemon.Type)
			}
		}
		fmt.Println()
		fmt.Println("select which pokemon you want to catch: ")
		pokeList, err := pokemons.GetAll()
		if err != nil {
			fmt.Println("no pokemon to catch")
			fmt.Println(err)
			fmt.Scanln(&trash)
			break
		}
		for _, pokemon := range pokeList {
			fmt.Printf("%d - called: %s with type: %s with chance: %d%%\n", pokemon.ID, pokemon.Name, pokemon.Type, int(pokemon.CatchRate*100))
		}
		exit := len(pokeList) + 1
		fmt.Printf("%d to exit\n", exit)
		fmt.Println()
		fmt.Print("input: ")
		menu := util.InputInt(scanner, "Menu")
		if menu == -1 {
			continue
		}
		if menu == exit {
			return
		}
		poke, err := pokemons.GetById(menu)
		if err != nil {
			util.ConsoleClear()
			fmt.Println(err)
			fmt.Println("press enter to continue")
			fmt.Scanln(&trash)
			continue
		}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		prob := r.Float64()
		if prob < poke.CatchRate {
			util.ConsoleClear()
			fmt.Printf("you failed to catch "+poke.Name+" with %d%% chance - you rolled: %d\npress enter to continue\n", int(poke.CatchRate*100), int(prob*100))
			fmt.Scanln(&trash)
			continue
		}
		util.ConsoleClear()
		fmt.Printf("you catched "+poke.Name+" with %d%% chance - you rolled: %d\npress enter to continue\n", int(poke.CatchRate*100), int(prob*100))
		fmt.Scanln(&trash)
		player.Bag = append(player.Bag, poke)
		players.Update(&player)
	}
}

func login(scanner *bufio.Scanner) domain.Player {
	util.ConsoleClear()
	fmt.Print("input player name: ")
	name := util.InputString(scanner, "Name")

	if name == "nil" {
		return login(scanner)
	}

	return domain.Player{
		Name: name,
		Bag:  []domain.Pokemon{},
	}
}

func gameLoop(players handler.PlayerHandler, pokemons handler.PokemonHandler, scanner *bufio.Scanner) {
	for {
		util.ConsoleClear()
		fmt.Println("Poke Go")
		fmt.Println()
		fmt.Println("1. Login")
		fmt.Println("2. Exit")
		fmt.Println()
		menu := util.InputInt(scanner, "Menu")
		if menu == -1 {
			continue
		}
		if menu < 1 || menu > 2 {
			util.ConsoleClear()
			fmt.Println("input need to be between 1 until 2\npress enter to continue")
			fmt.Scanln(&trash)
			continue
		}
		switch menu {
		case 1:
			player := login(scanner)
			players.Save(&player)
			pokeGo(player, players, pokemons, scanner)
		case 2:
			util.ConsoleClear()
			fmt.Println("thanks for playing")
			time.Sleep(time.Duration(3 * time.Second)) // 3 secs delay
			util.ConsoleClear()
			os.Exit(0)
		}
	}
}

func main() {
	players, pokemons := autowired()
	scanner := bufio.NewScanner(os.Stdin)
	fillPokemons(pokemons)
	gameLoop(players, pokemons, scanner)
}
