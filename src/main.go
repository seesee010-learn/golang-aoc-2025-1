package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type needs struct {
	direction string // ether "L" or "R"
	number    int    // 0 to 99
}

var global int = 50
var gCount int

func mod(a, b int) int {
	return ((a % b) + b) % b
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// create a list of needs
	var needsList []needs
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" { // change this if you are using windows
			continue
		}
		trimmed := strings.TrimSpace(line)

		n, err := strconv.Atoi(trimmed[1:])
		if err != nil {
			fmt.Println("Error: invalid number")
			fmt.Println(trimmed)
			return
		}

		switch trimmed[0] {
		case 'L', 'R':
			needsList = append(needsList, needs{direction: trimmed[:1], number: n})
		default:
			fmt.Println("Error: unknown direction")
			fmt.Println(trimmed)
			return
		}
	}
	for _, n := range needsList {
		switch n.direction {
		case "R":
			gCount += (global + n.number) / 100
			global = mod(global+n.number, 100)
		case "L":
			gCount += (mod(100-global, 100) + n.number) / 100
			global = mod(global-n.number, 100)
		}
	}
	fmt.Println(gCount)
}
