package main

import (
	"errors"
	"fmt"
)

func tester(value int) (int, error) {
	if value%2 == 0 {
		fmt.Println(errors.New("value must not be a even number"))
		var newValue int
		fmt.Scanf(" %s", &newValue)
		tester(newValue)
	}
	return value, nil
}

func main() {
	var name int
	fmt.Println("Insert a value into")
	fmt.Scanf(" %s", &name)
	fmt.Println(tester(name))
}
