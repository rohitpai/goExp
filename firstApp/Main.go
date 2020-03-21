package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 42
	fmt.Printf("%T\n", i)
	j := strconv.Itoa(i)
	fmt.Printf("%v, %T", j, j)
}
