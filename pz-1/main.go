package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func areSimilar(a float64, b float64, c float64, d float64) (bool, error) {
	if a <= 0 || b <= 0 || c <= 0 || d <= 0 {
		return false, errors.New("сathetus must be of positive length")
	}

	return a/c == b/d || a/d == b/c, nil
}

func readTwoFloat64() (float64, float64, error) {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)

	if len(parts) != 2 {
		return 0, 0, errors.New("expected two real numbers")
	}

	a, err1 := strconv.ParseFloat(parts[0], 64)
	b, err2 := strconv.ParseFloat(parts[1], 64)

	if err1 != nil || err2 != nil {
		return 0, 0, errors.New("expected two real numbers")
	}

	return a, b, nil
}

func task1() {
	var a, b, c, d float64
	var err error

	fmt.Print("First triangle cathetus (enter two numbers): ")
	a, b, err = readTwoFloat64()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("Second triangle cathetus (enter two numbers): ")
	c, d, err = readTwoFloat64()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	res, err := areSimilar(a, b, c, d)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Triangles are similar: %t\n", res)
}

func isEvenTwoDigit(num int) bool {
	absNum := int(math.Abs(float64(num)))
	if absNum < 10 || absNum > 99 {
		return false
	}
	if absNum%2 != 0 {
		return false
	}

	return true
}

func readInt() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)

	if err != nil {
		return 0, errors.New("expected valid integer number")
	}

	return num, nil
}

func task2() {
	var num int
	var err error

	fmt.Print("Enter an integer number ")
	num, err = readInt()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if isEvenTwoDigit(num) {
		fmt.Printf("Your number is even two-digit\n")
	} else {
		fmt.Printf("Your number is not even two-digit\n")
	}
}

func getAgeEnding(age int) string {
	switch {
	case age%10 == 1 && age != 11:
		return "рік"
	case (age%10 >= 2 && age%10 <= 4) && !(age >= 12 && age <= 14):
		return "роки"
	default:
		return "років"
	}
}

func printMyAge(age int) {
	if age < 0 || age > 99 {
		fmt.Println("Число років має бути у діапазоні від 1 до 99!")
		return
	}

	fmt.Printf("Мені %d %s\n", age, getAgeEnding(age))
}

func task3() {
	for age := 1; age < 100; age++ {
		printMyAge(age)
	}
}

func reverseNumber(num int64) (int64, error) {
	if num <= 0 {
		return 0, errors.New("must be a natural number")
	}

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
	numbers := []int64{123, math.MaxInt64, math.MaxInt64 - 8, -123, 0}

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

func findMaxEvenInRow(matrix [][]int, row int) (int, error) {
	maxEven := math.MinInt32

	if len(matrix) < row {
		return maxEven, errors.New("Matrix does not have " + strconv.Itoa(row) + "(th) row")
	}

	for _, value := range matrix[row-1] {
		if value%2 == 0 && value > maxEven {
			maxEven = value
		}
	}

	if maxEven == math.MinInt32 {
		return -1, nil
	}

	return maxEven, nil
}

func printMatrix(matrix [][]int) {
	for i, row := range matrix {
		fmt.Printf("%d. ", i+1)
		for _, elem := range row {
			fmt.Printf("%d ", elem)
		}
		fmt.Println()
	}
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
		{
			{1, 2, 3, 4, 5, 6},
		},
	}

	const row = 5

	for _, matrix := range matrices {
		fmt.Println("Matrix:")
		printMatrix(matrix)

		maxEven, err := findMaxEvenInRow(matrix, row)
		if err != nil {
			fmt.Println(err)
		} else if maxEven == -1 {
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

func validateInput(input string, validators ...func(input string) error) bool {
	for _, validator := range validators {
		validationErr := validator(input)

		if validationErr != nil {
			fmt.Println(validationErr)
			return false
		}
	}

	return true
}

func readUntilValid(validators ...func(input string) error) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if err != nil {
			fmt.Println("Failed to read input")
			continue
		}

		if validateInput(input, validators...) {
			return input
		}
	}
}

func isValidTime(time string) bool {
	pattern := "^(2[0-3]|[01]?[0-9]):[0-5][0-9]$"

	res, err := regexp.Match(pattern, []byte(time))
	if err != nil {
		return false
	}
	return res
}

func readTrainArray() [8]TRAIN {
	var res [8]TRAIN

	for i := 0; i < 8; i++ {
		fmt.Printf("Enter destination for train %d: ", i+1)
		res[i].NAZN = readUntilValid()

		fmt.Printf("Enter number for train %d: ", i+1)
		trainNumberString := readUntilValid(func(input string) error {
			trainNumber, err := strconv.Atoi(input)

			if err != nil {
				return err
			}

			if slices.ContainsFunc(res[:], func(train TRAIN) bool {
				return train.NUMR == trainNumber
			}) {
				return errors.New("provided number is already used")
			}

			return nil
		})
		res[i].NUMR, _ = strconv.Atoi(trainNumberString)

		fmt.Printf("Enter departure time (format HH:MM) for train %d: ", i+1)
		res[i].TIME = readUntilValid(func(input string) error {
			if !isValidTime(input) {
				return errors.New("invalid time format")
			}
			return nil
		})

		fmt.Println()
	}

	return res
}

func sortAlphabetically(trainArr []TRAIN) {
	sort.Slice(trainArr[:], func(i, j int) bool {
		return strings.ToLower(trainArr[i].NAZN) < strings.ToLower(trainArr[j].NAZN)
	})
}

func task6() {
	RASP := readTrainArray()

	fmt.Println("Sorted alphabetically by destination name:")
	sortAlphabetically(RASP[:])
	fmt.Println(RASP)

	var inputTime string
	fmt.Print("Enter time (format HH:MM): ")
	inputTime = readUntilValid(func(input string) error {
		if !isValidTime(input) {
			return errors.New("invalid time format")
		}
		return nil
	})

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

func pause() {
	fmt.Println("Press ENTER to continue")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	const delim = "=================="

	fmt.Println("Task 1:")
	task1()
	fmt.Println(delim)

	fmt.Println("Task 2:")
	task2()
	fmt.Println(delim)

	pause()
	fmt.Println("Task 3:")
	task3()
	fmt.Println(delim)

	pause()
	fmt.Println("Task 4:")
	task4()
	fmt.Println(delim)

	pause()
	fmt.Println("Task 5:")
	task5()
	fmt.Println(delim)

	pause()
	fmt.Println("Task 6:")
	task6()
	fmt.Println(delim)
}
