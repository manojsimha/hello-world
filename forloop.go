package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Print(strconv.Itoa(i) + " ")
	}

	for j := 0; j <= 5; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Print(j)

	}
}
