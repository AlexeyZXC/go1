package main

import (
	"fmt"
	"go1/internal/gamefield"
	"go1/internal/player"
)

func main() {
	var gmfield gamefield.Gamefield
	var err error
	var player1, player2 player.Player
	var side string

	for {
		var size uint
		fmt.Println("Enter size of gamefield")
		fmt.Scan(&size)
		if gmfield, err = gamefield.NewGamefield(size); err != nil {
			fmt.Println(err.Error())
			continue
		}
		break
	}

	for {
		var name string
		fmt.Println("Enter name of player 1")
		fmt.Scan(&name)
		fmt.Println("Enter player 1 chip, 'x' or '0' ")
		_, err := fmt.Scan(&side)
		if err != nil {
			fmt.Println("Side enter error: ", err.Error())
			continue
		}
		if side != "x" && side != "0" {
			fmt.Println("x or 0 please.")
			continue
		}
		player1, err = player.NewPlayer(name, side)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		break
	}

	for {
		var name string
		fmt.Println("Enter name of player 2")
		fmt.Scan(&name)
		var side2 string = "x"
		if side == "x" {
			side2 = "0"
		}
		player2, err = player.NewPlayer(name, side2)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		break
	}

	var player player.Player
	var row, column uint

	for i := 1; ; i++ {
		player = player1
		if i%2 == 0 {
			player = player2
		}

		for {
			fmt.Println(player.GetName() + " make your move")
			fmt.Scan(&row, &column)
			if status, err := gmfield.SetCell(row, column, player.GetChip()); err != nil {
				fmt.Println(err)
				continue
			} else if status != "" {
				fmt.Println(status)
				return
			}
			break
		}

		gmfield.Print()
	}
}
