package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What will you play? ")
	chip, _ := reader.ReadString('\n')
	fmt.Println("You play with " + chip)
}
