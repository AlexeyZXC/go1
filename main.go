package main

import (
	"bufio"
	"fmt"
	"os"
)

func readRowColumn() (row byte, column byte) {
	fmt.Println("Enter row and column")
	fmt.Scan(&row, &column)
	return row, column
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What will you play? ")
	chip, _ := reader.ReadString('\n')
	fmt.Println("You play with: " + chip)

	// enter row and column
	var row, column byte
	row, column = readRowColumn()
	fmt.Printf("Row: %v; Column: %v ", row, column)
}
