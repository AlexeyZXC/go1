package logic

const retWin = "Вы выиграли!"
const retDraw = "Конец игры. Ничья. Все поля заняты"

type Logic struct {
}

func (l Logic) CheckGameStatus(g [][]string, emptyFieldValue string, chip string) (status string) {
	size := uint(len(g))
	totalMoves := size * size
	var sumOfmoves uint = 0

	for _, r := range g {
		var chipsInLine uint = 0
		for _, c := range r {
			if c != emptyFieldValue {
				sumOfmoves++
			}
			if c == chip {
				if chipsInLine++; chipsInLine == size {
					return retWin
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
			if c := g[j][i]; c == chip {
				if chipsInLine++; chipsInLine == size {
					return retWin
				}
			}
		}
	}

	//check main diagonal
	chipsInLine = 0
	for i = 0; i < size; i++ {
		if c := g[i][i]; c == chip {
			if chipsInLine++; chipsInLine == size {
				return retWin
			}
		}
	}

	//check lateral diagonal
	chipsInLine = 0
	for i = 0; i < size; i++ {
		if c := g[i][size-i-1]; c == chip {
			if chipsInLine++; chipsInLine == size {
				return retWin
			}
		}
	}

	if totalMoves == sumOfmoves {
		return retDraw
	}

	return ""
}
