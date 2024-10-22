package main

import (
	"fmt"
	"reflection/basics/assomptotic_notations"
)

func main() {

	runc := assomptotic_notations.New()

	runc.HowLong(countdown, true, 10)
}

func countdown(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("value is %d\n", i)
	}
}
