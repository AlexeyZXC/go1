package gamefield

func (g Gamefield) checkGameStatus(chip string) (status string) {
	size := g.GetSize()
	totalMoves := size * size
	var sumOfmoves uint = 0
	//check rows
	for _, r := range g.GetGamefield() {
		var chipsInLine uint = 0
		for _, c := range r {
			if c != " " {
				sumOfmoves++
			}
			if c == chip {
				if chipsInLine++; chipsInLine == size {
					return "Вы выиграли!"
				}
			}
		}
	}
	//check columns
	var chipsInLine uint = 0
	var i, j uint
	for i = 0; i < size; i++ {
		chipsInLine = 0
		for j = 0; j < size; j++ {
			if c, _ := g.GetCell(j, i); c == chip {
				if chipsInLine++; chipsInLine == size {
					return "Вы выиграли!"
				}
			}
		}
	}

	//check main diagonal
	chipsInLine = 0
	for i = 0; i < size; i++ {
		if c, _ := g.GetCell(i, i); c == chip {
			if chipsInLine++; chipsInLine == size {
				return "Вы выиграли!"
			}
		}
	}

	//check lateral diagonal
	chipsInLine = 0
	for i = 0; i < size; i++ {
		if c, _ := g.GetCell(i, size-i-1); c == chip {
			if chipsInLine++; chipsInLine == size {
				return "Вы выиграли!"
			}
		}
	}

	if totalMoves == sumOfmoves {
		return "Конец игры. Ничья. Все поля заняты"
	}

	return ""
}
