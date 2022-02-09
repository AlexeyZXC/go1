package main

import (
	"fmt"
	"go1/internal/player"
)

func main() {
	var player1, player2 player.Player
	var sideX bool
	for {
		var name string
		fmt.Println("Enter name of player 1")
		fmt.Scan(&name)
		fmt.Println("Enter your side (enter 'true' to play with X)")
		_, err := fmt.Scan(&sideX)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		player1, err = player.NewPlayer(name, sideX == true)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		break
	}

	for {
		var name string
		var err error
		fmt.Println("Enter name of player 2")
		fmt.Scan(&name)
		player2, err = player.NewPlayer(name, sideX == false)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		break
	}

	for {
		var row, column uint
		fmt.Println(player1.GetName() + " make your move")
		fmt.Scan(&row, &column)
		fmt.Println(player2.GetName() + " make your move")
		fmt.Scan(&row, &column)
	}
}
