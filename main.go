package main

import (
	"fmt"
	"go1/internal/configHelper"
	"go1/internal/console"
	"go1/internal/gamefield"
	"go1/internal/interfaces"
	"go1/internal/logic"
	"go1/internal/player"
)

func main() {

	var err error
	conf := configHelper.Config{}

	var useEnvironmentVarsConfig bool = false
	var useYamlConfig bool = true

	if useYamlConfig == true {
		if err = conf.ReadFromYaml(); err != nil {
			fmt.Println("ReadFromYaml error: ", err, conf)
			return
		}
		fmt.Println("ReadFromYaml ok: ", conf)
	}

	if useEnvironmentVarsConfig == true {

		conf = configHelper.Config{Port: 8080, Db_url: "https://github.com/AlexeyZXC/go1"}

		configHelper.SetPrefix("myPref")
		err = conf.WriteEnvVar()
		if err != nil {
			fmt.Println("WriteEnvVar error: ", err)
			return
		}

		fmt.Println("WriteEnvVar ok")

		if err := conf.ReadEnvVar(); err != nil {
			fmt.Println("ReadEnvVar error: ", err, conf)
			return
		}
		fmt.Println("ReadEnvVar ok", conf)
	}

	//-----------------------------

	var gmfield gamefield.Gamefield
	//var err error
	var player1, player2 player.Player
	var side string

	var logic logic.Logic

	var irw interfaces.IReadWrite
	var console console.ReadWrite
	irw = console //use console as input/output

	for {
		var size uint
		irw.Write("Enter size of gamefield")
		irw.Read(&size)
		if gmfield, err = gamefield.NewGamefield(size, logic, console); err != nil {
			irw.Write(err.Error())
			continue
		}
		break
	}

	for {
		var name string
		irw.Write("Enter name of player 1")
		irw.Read(&name)
		irw.Write("Enter player 1 chip, 'x' or '0' ")
		_, err := irw.Read(&side)
		if err != nil {
			irw.Write("Side enter error: ", err.Error())
			continue
		}
		if side != "x" && side != "0" {
			irw.Write("x or 0 please.")
			continue
		}
		player1, err = player.NewPlayer(name, side)
		if err != nil {
			irw.Write(err.Error())
			continue
		}
		break
	}

	for {
		var name string
		irw.Write("Enter name of player 2")
		irw.Read(&name)
		var side2 string = "x"
		if side == "x" {
			side2 = "0"
		}
		player2, err = player.NewPlayer(name, side2)
		if err != nil {
			irw.Write(err.Error())
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
			irw.Write(player.GetName() + " make your move")
			irw.Read(&row, &column)
			if status, err := gmfield.SetCell(row, column, player.GetChip()); err != nil {
				irw.Write(err)
				continue
			} else if status != "" {
				irw.Write(status)
				return
			}
			break
		}

		gmfield.Print()
	}
}
