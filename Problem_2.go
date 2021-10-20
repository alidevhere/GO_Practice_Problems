package main

import (
	"fmt"
)

func main() {
	//4000000
	var n1 int = 1
	var n2 int = 2
	var n3 int = 0
	var sum int = 2
	for n3 < 4000000 {
		n3 = n1 + n2
		fmt.Println(n3)
		n1 = n2
		n2 = n3
		if n3%2 == 0 {
			sum += n3
		}

	}
	fmt.Printf("sum of even terms %d", sum)

}
