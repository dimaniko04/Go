package utils

import (
	"math"
	"sync"
)

func equation(x float64) float64 {
	return x*x*x - x*x - 6*x + 6
}

func GoldenSectionSearch(a, b, tol float64, resultChan chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	phi := (1 + math.Sqrt(5)) / 2
	iteration := 0

	for b-a > tol {
		iteration++
		x1 := b - (b-a)/phi
		x2 := a + (b-a)/phi

		if equation(x1) < equation(x2) {
			a = x1
		} else {
			b = x2
		}
	}
	resultChan <- (a + b) / 2
}

func ParabolicMethod(a, b float64, n int, resultChan chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	h := (b - a) / float64(n)
	root := equation(a) + equation(b)

	for i := 1; i < n; i++ {
		k := 2 + 2*(i%2)
		root += float64(k) * equation(a+float64(i)*h)
	}
	resultChan <- root * h / 3
}
