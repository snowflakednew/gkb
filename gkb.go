package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var playerwins, computerwins, round uint
	validmoves := [3]string{"scissor", "paper", "rock"}
	var move, computer string
	weakness := map[string]string{
		"scissor": "rock",
		"paper":   "scissor",
		"rock":    "paper",
	}
	fmt.Printf("\033[2J") // ANSI clear screen
	for {
		computer = validmoves[rand.Intn(3)]
		fmt.Printf("Round %d\n", round)
		fmt.Print("Scissor, paper or rock! ")
		_, err := fmt.Scanln(&move)
		if err != nil {
			panic(err)
		}
		move = strings.ToLower(move)
		switch true {
		case computer == move:
			fmt.Printf("It's a tie. %q (you) vs %q (computer)\n", move, computer)
			round++
		case weakness[computer] == move:
			fmt.Printf("Hoorah, You win! %q (you) vs %q (computer)\n", move, computer)
			playerwins++
			round++
		case weakness[move] == computer:
			fmt.Printf("Aww man, You lost. %q (you) vs %q (computer)\n", move, computer)
			computerwins++
			round++
		default:
			fmt.Println("Invalid input.")
		}
		time.Sleep(750 * time.Millisecond)
	}
}
