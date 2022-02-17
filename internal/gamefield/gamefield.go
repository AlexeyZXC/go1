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
	if size < 2 {
		return Gamefield{}, errors.New("Размер поля не может быть меньше 2")
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

func (g *Gamefield) SetCell(x, y uint, chip string) (string, error) {
	if x > g.size-1 || y > g.size-1 {
		return "", errors.New("Неверная клетка")
	}
	if g.gamefield[x][y] != " " {
		return "", errors.New("Клетка занята")
	}
	g.gamefield[x][y] = chip
	return g.checkGameStatus(chip), nil
}

func (g Gamefield) GetCell(x, y uint) (cell string, err error) {
	if x > g.size-1 || y > g.size-1 {
		return "", errors.New("Нет такой ячейки")
	}
	return g.gamefield[x][y], nil
}

func (g Gamefield) GetGamefield() [][]string {
	return g.gamefield
}

func (g Gamefield) GetSize() uint {
	return g.size
}

func (g Gamefield) Print() {
	for i := range g.gamefield {
		fmt.Println(g.gamefield[i])
	}
}
