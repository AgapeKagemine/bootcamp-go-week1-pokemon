package util

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func ConsoleClear() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") // cmd clear
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear") // universal?
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	ConsolePosition()
}

func ConsolePosition() {
	print("\033[50;H")
}

func IsEmpty[T any](list map[int]T) bool {
	return len(list) == 0
}

func IsExist[T any](list map[int]T, id int) bool {
	if _, exists := list[id]; exists {
		return true
	}
	return false
}

// Cheat code
func IsConnected(T any) bool {
	return T != nil
}

var trash string // to catch extra input when asking for enter

func InputString(scanner *bufio.Scanner, text string) (input string) {
	for {
		if scanner.Scan() {
			input = scanner.Text()
			if len(input) < 1 {
				ConsoleClear()
				fmt.Println(text + " need to have at least 1 character\npress enter to continue")
				fmt.Scanln(&trash)
				return "nil"
			}
			return input
		}
	}
}

func InputInt(scanner *bufio.Scanner, text string) (input int) {
	for {
		if scanner.Scan() {
			input, err := strconv.Atoi(scanner.Text())
			if err != nil {
				ConsoleClear()
				fmt.Println(text + " need to be a number\npress enter to continue")
				fmt.Scanln(&trash)
				return -1
			}
			return input
		}
	}
}
