package main

import (
	"go1/internal/console"
	"testing"
)

func Test_1(t *testing.T) {
	ConsoleTest.T = t
	ConsoleTest.TestFlow = make(map[int]console.InOutResponce)

	ConsoleTest.TestFlow[0] = console.InOutResponce{
		InValue: uint(2), OutValue: "Enter size of gamefield",
	}
	ConsoleTest.TestFlow[1] = console.InOutResponce{
		InValue: "", OutValue: "size is 2",
	}
	ConsoleTest.TestFlow[2] = console.InOutResponce{
		InValue: "name1", OutValue: "Enter name of player 1",
	}
	ConsoleTest.TestFlow[3] = console.InOutResponce{
		InValue: "x", OutValue: "Enter player 1 chip, 'x' or '0' ",
	}
	ConsoleTest.TestFlow[4] = console.InOutResponce{
		InValue: "name2", OutValue: "Enter name of player 2",
	}

	ConsoleTest.TestFlow[5] = console.InOutResponce{
		InValue: uint(0), InValue2: uint(0), OutValue: "name1 make your move",
	}

	ConsoleTest.TestFlow[6] = console.InOutResponce{
		InValue: uint(1), InValue2: uint(0), OutValue: "name2 make your move",
	}

	ConsoleTest.TestFlow[7] = console.InOutResponce{
		InValue: uint(0), InValue2: uint(1), OutValue: "name1 make your move",
	}

	ConsoleTest.TestFlow[8] = console.InOutResponce{
		InValue: uint(0), OutValue: "Вы выиграли!",
	}

	main()

}
