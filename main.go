package main

import "fmt"

func main() {
	// pi := EstimatePiSync(100000000)
	pi := EstimatePiConcurrent(100000000)
	fmt.Println("Pi is ", pi)
}
