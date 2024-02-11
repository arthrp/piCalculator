package main

import (
	"math/rand"
)

func EstimatePiSync(iters int) float64 {
	pointsInCircleQuad := getPointsInCircle(iters)

	ratio := float64(pointsInCircleQuad) / float64(iters)

	return ratio * 4.0
}

func getPointsInCircle(iters int) int {
	pointsInCircleQuad := 0

	for i := 0; i < iters; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if isPointInCircle(x, y) {
			pointsInCircleQuad++
		}
	}

	return pointsInCircleQuad
}

func EstimatePiConcurrent(iters int) float64 {
	const goroutineCount = 5
	pointsInCircleQuad := 0
	channels := make([]chan int, goroutineCount)
	itersPerRoutine := iters / goroutineCount

	//Send
	for i := 0; i < goroutineCount; i++ {
		channels[i] = make(chan int)
		go sendPoints(itersPerRoutine, channels[i])
	}

	//Receive
	for _, ch := range channels {
		newPoints := <-ch
		pointsInCircleQuad += newPoints
	}

	ratio := float64(pointsInCircleQuad) / float64(iters)

	return ratio * 4.0
}

func sendPoints(itersPerRoutine int, ch chan<- int) {
	points := getPointsInCircle(itersPerRoutine)
	ch <- points
	close(ch)
}

func isPointInCircle(x float64, y float64) bool {
	return x*x+y*y < 1.0
}
