package go_say_hello

import (
	"errors"
	"fmt"
)

func SayHello() string {
	return "Hello World"
}

func Addition(value interface{}) (int, error) {
	switch num := value.(type) {
	case int:
		bil := value.(int)
		return bil + 10, errors.New("No error")

	default:
		fmt.Println(num)
		return 0, errors.New("Not an integer\n")
	}
}
