package main

import "fmt"

func task1a() {
	fmt.Println("Task 1, part a:")

	intArray, err := ReadIntArray()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n===Array from indices of elements equal to zero===")
	fmt.Println("Initial array: " + intArray.String())
	indicesOfZeros := intArray.Process()
	fmt.Println("Result: " + indicesOfZeros.String())
	fmt.Print("==================================================\n\n")

	intMatrix, err := ReadIntMatrix()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n===Reverse even rows of matrix===")
	fmt.Print("Initial matrix:\n" + intMatrix.String())
	intMatrix.Process()
	fmt.Print("Result:\n" + intMatrix.String())
	fmt.Print("==================================\n\n")
}

func task1b() {
	fmt.Println("Task 1, part b:")

	var processable Processable
	var err error

	processable, err = ReadIntArray()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n===Array from indices of elements equal to zero===")
	fmt.Println("Initial array: " + processable.String())
	indicesOfZeros := processable.Process()
	fmt.Println("Result: " + indicesOfZeros.String())
	fmt.Print("==================================================\n\n")

	matrix, err := ReadIntMatrix()
	if err != nil {
		fmt.Println(err)
		return
	}
	processable = &matrix

	fmt.Println("\n===Reverse even rows of matrix===")
	fmt.Print("Initial matrix:\n" + processable.String())
	processable.Process()
	fmt.Print("Result:\n" + processable.String())
	fmt.Print("==================================\n\n")
}

func main() {
	task1a()
	task1b()
}
