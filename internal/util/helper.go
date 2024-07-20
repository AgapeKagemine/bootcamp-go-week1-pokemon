package util

import (
	"os"
	"os/exec"
	"runtime"
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
		return exists
	}
	return false
}

// Cheat code
func IsConnected(T any) bool {
	return T != nil
}
