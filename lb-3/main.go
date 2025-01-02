package main

import (
	"fmt"
	"lb-3/utils"
	"os"
	"sync"
)

func generateSquaresArray(n int) []int {
	squares := make([]int, n)

	for i := 1; i <= n; i++ {
		squares[i-1] = i * i
	}

	return squares
}

func task1() {
	filename := "squares.txt"
	numbers := generateSquaresArray(100)

	fmt.Println("Numbers:", numbers)

	if err := utils.CreateFileOfSquareRoots(filename, numbers); err != nil {
		fmt.Println(err)
		return
	}

	sum, err := utils.SumFileComponents(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("File components sum: %.2f\n", sum)
}

func task2() {
	filename := "dvds.txt"
	dvds := []utils.DVD{
		{
			Title:    "Inception",
			Director: "Christopher Nolan",
			Duration: 148,
			Price:    9.99,
		},
		{
			Title:    "The Matrix",
			Director: "The Wachowskis",
			Duration: 136,
			Price:    14.99,
		},
		{
			Title:    "Interstellar",
			Director: "Christopher Nolan",
			Duration: 169,
			Price:    19.99,
		},
		{
			Title:    "The Godfather",
			Director: "Francis Ford Coppola",
			Duration: 175,
			Price:    24.99,
		},
	}

	if err := utils.WriteDVDsToFile(filename, dvds); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File content:")
	if err := utils.PrintFileContent(filename); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nDeleting dvds costing more that 15...")
	if err := utils.DeleteExpensiveDVDs(filename, 15.00); err != nil {
		fmt.Println(err)
		return
	}

	newDVD := utils.DVD{
		Title:    "Avatar",
		Director: "James Cameron",
		Duration: 162,
		Price:    12.99,
	}
	fmt.Println("Adding new dvd to file at position 10...")
	if err := utils.AddDVD(filename, newDVD, 10); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Adding new dvd to file at position 1...")
	if err := utils.AddDVD(filename, newDVD, 1); err != nil {
		fmt.Println(err)
	}

	fmt.Println("File content:")
	if err := utils.PrintFileContent(filename); err != nil {
		fmt.Println(err)
		return
	}
}

func task3() {
	n := 1000
	tol := 1e-5
	a, b := -3.0, 3.0
	resultChan := make(chan float64, 2)
	var wg sync.WaitGroup

	wg.Add(2)
	go utils.GoldenSectionSearch(a, b, tol, resultChan, &wg)
	go utils.ParabolicMethod(a, b, n, resultChan, &wg)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	file, err := os.Create("results.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to write to file: %w", r.(error))
		}
		return
	}()

	for r := range resultChan {
		_, err := file.WriteString(fmt.Sprintf("Root: %.4f\n", r))
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Results have been written to result.txt")
}

func main() {
	const delim = "=================="

	fmt.Println("Task 1:")
	task1()
	fmt.Println(delim)

	fmt.Println("Task 2:")
	task2()
	fmt.Println(delim)

	fmt.Println("Task 3:")
	task3()
	fmt.Println(delim)
}
