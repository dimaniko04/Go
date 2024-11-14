package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

func pause() {
	fmt.Println("Press ENTER to continue")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func task1a() {
	intArray, err := ReadIntArray()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n===Array from indices of elements equal to zero===")
	fmt.Println("Initial array:", intArray)
	fmt.Println("Result:", intArray.Process())
	fmt.Print("==================================================\n\n")

	intMatrix, err := ReadIntMatrix()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n===Reverse even rows of matrix===")
	fmt.Println("Initial matrix:")
	fmt.Println(intMatrix)
	fmt.Println("Result:")
	fmt.Println(intMatrix.Process())
	fmt.Print("==================================\n\n")
}

func task1b() {
	var err error
	var processable Processable

	processable, err = ReadIntArray()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n===Array from indices of elements equal to zero===")
	fmt.Println("Initial array:", processable)
	fmt.Println("Result:", processable.Process())
	fmt.Print("==================================================\n\n")

	matrix, err := ReadIntMatrix()
	if err != nil {
		fmt.Println(err)
		return
	}
	processable = &matrix

	fmt.Println("\n===Reverse even rows of matrix===")
	fmt.Println("Initial matrix:")
	fmt.Println(processable)
	fmt.Println("Result:")
	fmt.Println(processable.Process())
	fmt.Print("==================================\n\n")
}

func f(x float64) float64 {
	return 0.1*math.Pow(x, 2) - x*math.Log(x)
}

func fDerivative(x float64) float64 {
	return 0.2*x - math.Log(x) - 1
}

func birthDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func initializeCharacters() Characters {
	return []Character{
		{
			User: User{
				address:   "123 Maple St, Springfield",
				name:      "Alice",
				birthDate: birthDate(1995, time.January, 10),
			},
			x:      12.34,
			y:      56.78,
			rating: 5,
		},
		{
			User: User{
				address:   "456 Oak Ave, Shelbyville",
				name:      "Bob",
				birthDate: birthDate(1988, time.March, 23),
			},
			x:      23.45,
			y:      67.89,
			rating: 4,
		},
		{
			User: User{
				address:   "789 Pine Blvd, Capital City",
				name:      "Charlie",
				birthDate: birthDate(2000, time.May, 15),
			},
			x:      34.56,
			y:      78.90,
			rating: 3,
		},
		{
			User: User{
				address:   "101 Elm St, Ogdenville",
				name:      "Diana",
				birthDate: birthDate(1992, time.August, 9),
			},
			x:      45.67,
			y:      89.01,
			rating: 4,
		},
		{
			User: User{
				address:   "202 Cedar St, North Haverbrook",
				name:      "Eve",
				birthDate: birthDate(1985, time.December, 5),
			},
			x:      56.78,
			y:      90.12,
			rating: 5,
		},
	}
}

func task2() {
	characters := initializeCharacters()
	fmt.Print("Characters list:\n", characters)
	pause()

	fmt.Println("Distance between characters:")
	for i := 1; i < len(characters); i++ {
		ch1 := characters[i-1]
		ch2 := characters[i]

		fmt.Printf("%s - %s\n", ch1.name, ch2.name)
		PrintDistance(ch1.x, ch1.y, ch2.x, ch2.y)
		pause()
	}

	fmt.Println("Distance between character and default quest:")
	for _, ch := range characters {
		fmt.Println(ch.name)
		PrintDistance(ch.x, ch.y, 0, 0)
		pause()
	}

	fmt.Println("Character with lowest rating:", MinValue(characters...))
	pause()

	fmt.Print("Initial characters list:\n", characters)
	pause()
	characters.Delete("Alice")
	fmt.Println("Deleting character with name Alice...")
	characters.Delete("Not exists")
	fmt.Println("Trying to delete character that does not exist...")
	pause()
	fmt.Print("Obtained character list:\n", characters)

	var eq IEquation = Equation{
		lowerBound: 1,
		upperBound: 2,
	}
	x := eq.newtonMethod(f, fDerivative)

	pause()
	fmt.Println("Equation - 0,1*x^2 - x*ln(x) = 0")
	fmt.Printf("Solution obtained by Newton's method - %.4f\n", x)
}

func main() {
	const delim = "=================="

	fmt.Println("Task 1 part a:")
	task1a()
	fmt.Println(delim)

	fmt.Println("Task 1 part b:")
	task1b()
	fmt.Println(delim)

	fmt.Println("Task 2:")
	task2()
	fmt.Println(delim)
}
