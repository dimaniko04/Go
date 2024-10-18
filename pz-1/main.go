package main

import (
	"fmt"
)

func areSimilar(a float64, b float64, c float64, d float64) bool {
	return a/c == b/d || a/d == b/c
}

func task1() {
	var a, b, c, d float64

	fmt.Print("First triangle cathetus (enter two numbers): ")
	fmt.Scanf("%f %f\n", &a, &b)

	fmt.Print("Second triangle cathetus (enter two numbers): ")
	fmt.Scanf("%f %f\n", &c, &d)

	fmt.Printf("Triangles are similar: %t", areSimilar(a, b, c, d))
}

func main() {
	fmt.Println("Task 1:")
	task1()
	fmt.Println("==================")
}
