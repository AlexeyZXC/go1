package console

import (
	"fmt"
	"strconv"
	"testing"
)

type ReadWrite struct {
}

func (rw ReadWrite) Read(i ...interface{}) (ii int, err error) {
	return fmt.Scan(i...)
}

func (rw ReadWrite) Write(i ...interface{}) (ii int, err error) {
	return fmt.Println(i...)
}

//------------- for testing

type InOutResponce struct {
	InValue  interface{}
	InValue2 interface{}
	OutValue string
}

type ReadWriteTest struct {
	InUintValue uint
	InStrValue  string
	OutStrValue string
	T           *testing.T
	TestFlow    map[int]InOutResponce
	RequestNum  *int
}

func (rw *ReadWriteTest) Init() {
	rw.RequestNum = new(int)
}

func (rw ReadWriteTest) Read(i ...interface{}) (ii int, err error) {

	req := *rw.RequestNum

	switch v := i[0].(type) {
	case *uint:
		*v = rw.TestFlow[req].InValue.(uint)
	case *string:
		*v = rw.TestFlow[req].InValue.(string)
	}

	args := len(i)

	if args > 1 {
		switch v := i[1].(type) {
		case *uint:
			*v = rw.TestFlow[req].InValue2.(uint)
		case *string:
			*v = rw.TestFlow[req].InValue2.(string)
		}
	}

	for _, b := range i {
		switch v := b.(type) {
		case *uint:
			fmt.Println(*v)
		case uint:
			fmt.Println(v)
		case *string:
			fmt.Println(*v)
		}
	}

	*rw.RequestNum++

	return 0, nil
}

func (rw ReadWriteTest) Write(i ...interface{}) (ii int, err error) {

	fmt.Println(i...)

	req := *rw.RequestNum

	if _, ok := i[0].(string); ok == false {
		return
	}

	args := len(i)

	resp := i[0].(string)

	if args > 1 {
		if a, ok := i[1].(uint); ok == true {
			resp += strconv.FormatUint(uint64(a), 10)
		}
	}

	if resp != rw.TestFlow[req].OutValue {
		rw.T.Fatal(i[0].(string) + " instead of " + rw.TestFlow[req].OutValue)
	}

	if req == 1 {
		*rw.RequestNum++
	}

	return 0, nil
}
