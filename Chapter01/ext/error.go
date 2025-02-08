package main

import "errors"
import "strconv"

func main() {
	err := doesReturnError("2", "AAA", "BBB", "12345")
	if err != nil {
		panic(err)
	}
}

func doesReturnError(args ...string) error {
	for n, v := range args {
		println(strconv.Itoa(n+1) + "   " + v + "######")
	}
	err := errors.New("this function simply returns an error")
	return err
}
