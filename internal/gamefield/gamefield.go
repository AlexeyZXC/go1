package gamefield

import (
	"errors"
	"fmt"
)

type Gamefield struct {
	size      uint
	gamefield [][]string
}

func NewGamefield(size uint) (Gamefield, error) {
	if size < 3 {
		return Gamefield{}, errors.New("Размер поля не может быть меньше 3")
	}
	var g Gamefield
	g.size = size
	g.gamefield = make([][]string, size)
	for i := range g.gamefield {
		g.gamefield[i] = make([]string, size)
	}
	for i := range g.gamefield {
		for j := range g.gamefield[i] {
			g.gamefield[i][j] = " "
		}
	}

	return g, nil
}

func (g *Gamefield) SetCell(x, y uint, side string) error {
	if x > g.size-1 || y > g.size-1 {
		return errors.New("Неверный ввод")
	}
	g.gamefield[x][y] = side
	return nil
}

func (g Gamefield) Print() {
	for i := range g.gamefield {
		fmt.Println(g.gamefield[i])
	}
}
