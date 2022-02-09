package gamefield

import "errors"

type Gamefield struct {
	size uint
}

func NewGamefield(size uint) (Gamefield, error) {
	if size < 3 {
		return Gamefield{}, errors.New("Размер поля не может быть меньше 3")
	}
	return Gamefield{size}, nil
}

func (g *Gamefield) SetCell(x, y uint, sideX bool) error {
	if x < 0 || x > g.size || y < 0 || y > g.size {
		return errors.New("Неверный ввод")
	}
	return nil
}
