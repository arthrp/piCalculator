package main

import (
	"flag"
	"fmt"
)

func main() {
	var useConcurrent bool
	var pi float64

	flag.BoolVar(&useConcurrent, "useConcurrent", true, "Whether to use concurrent calculation")
	flag.Parse()

	fmt.Println("Use concurrent is", useConcurrent)

	if useConcurrent {
		pi = EstimatePiConcurrent(100000000)
	} else {
		pi = EstimatePiSync(100000000)
	}

	fmt.Println("Pi is ", pi)
}
