package player

import "errors"

type Player struct {
	name      string
	chip      string
	winsCount uint
}

func NewPlayer(name string, chip string) (Player, error) {
	if err := validateName(name); err != nil {
		return Player{}, err
	}
	return Player{name, chip, 0}, nil
}

func validateName(name string) error {
	if len(name) < 3 {
		return errors.New("Имя должно быть длиной больше 3")
	}
	for _, x := range name {
		if x == ' ' {
			return errors.New("В имени не может быть пробела")
		}
	}
	return nil
}

func (p *Player) Rename(name string) error {
	if err := validateName(name); err != nil {
		return err
	}
	p.name = name
	return nil
}

func (p *Player) Win() {
	p.winsCount++
}

func (p Player) GetWins() uint {
	return p.winsCount
}

func (p Player) GetName() string {
	return p.name
}

func (p Player) GetChip() string {
	return p.chip
}
