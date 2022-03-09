package interfaces

type ICheckGameStatus interface {
	CheckGameStatus(gamefield [][]string, emptyFieldValue string, chip string) (status string)
}

type IReadWrite interface {
	Read(...interface{}) (int, error)
	Write(...interface{}) (int, error)
}
