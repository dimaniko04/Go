package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
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

	fmt.Printf("Triangles are similar: %t\n", areSimilar(a, b, c, d))
}

func isEvenTwoDigit(num int) bool {
	if num < 10 || num > 99 {
		return false
	}
	if num%2 != 0 {
		return false
	}

	return true
}

func task2() {
	var num int

	fmt.Print("Enter an integer number ")
	fmt.Scanf("%d\n", &num)

	if isEvenTwoDigit(num) {
		fmt.Printf("Your number is even two-digit\n")
	} else {
		fmt.Printf("Your number is not even two-digit\n")
	}
}

func printMyAge(age int) {
	if age < 0 || age > 99 {
		fmt.Println("Число років має бути у діапазоні від 1 до 99!")
		return
	}

	switch {
	case age%10 == 1 && age != 11:
		fmt.Printf("Мені %d рік\n", age)
	case (age%10 >= 2 && age%10 <= 4) && !(age >= 12 && age <= 14):
		fmt.Printf("Мені %d роки\n", age)
	default:
		fmt.Printf("Мені %d років\n", age)
	}
}

func task3() {
	for age := 1; age < 100; age++ {
		printMyAge(age)
	}
}

func reverseNumber(num int64) (int64, error) {
	var res int64 = 0

	for num != 0 {
		res = (res * 10) + (num % 10)

		if (num > 0 && res < 0) || (num < 0 && res > 0) {
			return 0, errors.New("reversal is impossible due to overflow")
		}

		num /= 10
	}

	return res, nil
}

func task4() {
	numbers := []int64{123, math.MaxInt64, math.MaxInt64 - 8, -123, math.MinInt64 + 9}

	for _, num := range numbers {
		fmt.Printf("Before: %d\n", num)
		reversed, err := reverseNumber(num)
		if err != nil {
			fmt.Println(err)
			fmt.Println()
			continue
		}
		fmt.Printf("Reversed: %d\n\n", reversed)
	}
}

func findMaxEvenInRow(matrix [][]int, row int) int {
	maxEven := math.MinInt32

	for _, value := range matrix[row-1] {
		if value%2 == 0 && value > maxEven {
			maxEven = value
		}
	}

	if maxEven == math.MinInt32 {
		return -1
	}

	return maxEven
}

func task5() {
	matrices := [][][]int{
		{
			{1, 3, 5, 7, 9},
			{2, 4, 6, 8, 10},
			{11, 13, 15, 17, 19},
			{20, 22, 24, 26, 28},
			{31, 32, 33, 34, 35},
		},
		{
			{1, 3, 5, 7, 9},
			{2, 4, 6, 8, 10},
			{11, 13, 15, 17, 19},
			{20, 22, 24, 26, 28},
			{31, 33, 33, 33, 35},
		},
		{
			{1, 3, 5, 7, 9},
			{2, 4, 6, 8, 10},
			{11, 13, 15, 17, 19},
			{20, 22, 24, 26, 28},
			{100, 33, 98, 33, 96},
		},
	}

	const row = 5

	for _, matrix := range matrices {
		fmt.Println("Matrix:")
		maxEven := findMaxEvenInRow(matrix, row)

		if maxEven == -1 {
			fmt.Println("No even numbers in 5th row")
		} else {
			fmt.Println("Max even number in 5th row:", maxEven)
		}
	}
}

type TRAIN struct {
	NAZN string
	NUMR int
	TIME string
}

func (train TRAIN) String() string {
	return fmt.Sprintf("Destination: %s, number: %d, departure time: %s\n", train.NAZN, train.NUMR, train.TIME)
}

func task6() {
	var RASP [8]TRAIN

	for i := 0; i < 8; i++ {
		fmt.Printf("Enter destination for train %d: ", i+1)
		fmt.Scan(&RASP[i].NAZN)

		fmt.Printf("Enter number for train %d: ", i+1)
		fmt.Scan(&RASP[i].NUMR)

		fmt.Printf("Enter departure time (format HH:MM) for train %d: ", i+1)
		fmt.Scan(&RASP[i].TIME)

		fmt.Println()
	}

	sort.Slice(RASP[:], func(i, j int) bool {
		return strings.ToLower(RASP[i].NAZN) < strings.ToLower(RASP[j].NAZN)
	})

	fmt.Println("Sorted alphabetically by destination name:")
	fmt.Println(RASP)

	var inputTime string
	fmt.Print("Enter time (format HH:MM): ")
	fmt.Scan(&inputTime)

	var departAfter []TRAIN

	for _, train := range RASP {
		if train.TIME > inputTime {
			departAfter = append(departAfter, train)
		}
	}

	if len(departAfter) == 0 {
		fmt.Println("No trains depart after", inputTime, ".")
	} else {
		fmt.Println("Trains that depart after", inputTime, ":")
		fmt.Println(departAfter)
	}
}

func main() {
	fmt.Println("Task 1:")
	task1()
	fmt.Println("==================")

	fmt.Println("Task 2:")
	task2()
	fmt.Println("==================")

	fmt.Println("Task 3:")
	task3()
	fmt.Println("==================")

	fmt.Println("Task 4:")
	task4()
	fmt.Println("==================")

	fmt.Println("Task 5:")
	task5()
	fmt.Println("==================")

	fmt.Println("Task 6:")
	task6()
	fmt.Println("==================")
}
