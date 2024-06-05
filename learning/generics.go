package learning

import "fmt"

func Generics() {
	fmt.Println(multipleTypes[int32](32))
}

func multipleTypes[T int32 | int16](slice T) T {
	sum := slice + 22

	return sum
}
