package main

import (
	"fmt"
)

func main() {
	var sum int = 0
	for i := 1; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}

	}
	fmt.Printf("sum = %d\n", sum)

}
