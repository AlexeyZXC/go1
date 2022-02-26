package console

import "fmt"

type ReadWrite struct {
}

func (rw ReadWrite) Read(i ...interface{}) (ii int, err error) {
	return fmt.Scan(i...)
}

func (rw ReadWrite) Write(i ...interface{}) (ii int, err error) {
	return fmt.Println(i...)
}
