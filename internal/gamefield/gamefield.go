package gamefield

import (
	"errors"
	"go1/internal/interfaces"
)

// type ICheckGameStatus interface {
// 	CheckGameStatus(gf Gamefield) (status string)
// }

const EmptyFieldValue = " "
const errFieldSizeNotLess = "Размер поля не может быть меньше 2"
const errWrongCell = "Неверная клетка"
const errCellOccupied = "Клетка занята"

type Gamefield struct {
	size             uint
	gamefield        [][]string
	icheckGameStatus interfaces.ICheckGameStatus
	irw              interfaces.IReadWrite
}

func NewGamefield(size uint, _iCheckGameStatus interfaces.ICheckGameStatus, _irw interfaces.IReadWrite) (Gamefield, error) {
	if size < 2 {
		return Gamefield{}, errors.New(errFieldSizeNotLess)
	}
	var g Gamefield
	g.size = size
	g.gamefield = make([][]string, size)
	for i := range g.gamefield {
		g.gamefield[i] = make([]string, size)
	}
	for i := range g.gamefield {
		for j := range g.gamefield[i] {
			g.gamefield[i][j] = EmptyFieldValue
		}
	}
	g.icheckGameStatus = _iCheckGameStatus
	g.irw = _irw

	return g, nil
}

func (g *Gamefield) SetCell(x, y uint, chip string) (string, error) {
	if x > g.size-1 || y > g.size-1 {
		return "", errors.New(errWrongCell)
	}
	if g.gamefield[x][y] != EmptyFieldValue {
		return "", errors.New(errCellOccupied)
	}
	g.gamefield[x][y] = chip

	return g.icheckGameStatus.CheckGameStatus(g.GetGamefield(), EmptyFieldValue, chip), nil
}

func (g Gamefield) GetCell(x, y uint) (cell string, err error) {
	if x > g.size-1 || y > g.size-1 {
		return "", errors.New(errWrongCell)
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
		g.irw.Write(g.gamefield[i])
	}
}
