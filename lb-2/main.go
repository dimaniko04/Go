package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readArray() ([]int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)

	intArr := make([]int, len(parts))
	for i, val := range parts {
		num, err := strconv.Atoi(val)

		if err != nil {
			return []int{}, errors.New("invalid input")
		}

		intArr[i] = num
	}

	return intArr, nil
}

func printIntArray(arr []int) {
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

func removeMinElements(arr []int) []int {
	minValue := arr[0]
	for _, val := range arr {
		if val < minValue {
			minValue = val
		}
	}

	result := []int{}
	for _, val := range arr {
		if val != minValue {
			result = append(result, val)
		}
	}
	return result
}

func task1() {
	fmt.Print("Enter initial array: ")
	arr, err := readArray()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	arr = removeMinElements(arr)
	fmt.Print("Array after removing minimum elements: ")
	printIntArray(arr)
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

type PrimeArr []int

/* func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
} */

func (p *PrimeArr) generate(n int) error {
	if n <= 0 {
		return errors.New("number of numbers to generate must be greater that 0")
	}
	*p = append(*p, 2)

	num := 3
	count := 1
	for count < n {
		isPrime := true
		for i := 0; ; i++ {
			if (*p)[i]*(*p)[i] > num {
				break
			}
			if num%(*p)[i] == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
			*p = append(*p, num)
		}

		num += 2
	}
	return nil
}

func task2() {
	fmt.Printf("Enter N: ")
	n, err := readInt()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	primeArr := PrimeArr{}
	err = primeArr.generate(n)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("First %d prime numbers: ", n)
	printIntArray(primeArr)
}

type SequenceGenerator interface {
	generateSequence() []int
}

type SquareMatrix [][]int

func readSquareMatrixRow(n int) ([]int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)

	if len(parts) < n {
		return []int{}, fmt.Errorf("expected %d elements", n)
	}

	intArr := make([]int, n)
	for i := 0; i < n; i++ {
		num, err := strconv.Atoi(parts[i])

		if err != nil {
			return []int{}, errors.New("invalid input")
		}

		intArr[i] = num
	}

	return intArr, nil
}

func readSquareMatrix(n int) (SquareMatrix, error) {
	fmt.Println("Enter matrix values:")
	var matrix = make(SquareMatrix, n)

	for i := 0; i < n; i++ {
		var row, err = readSquareMatrixRow(n)

		if err != nil {
			return SquareMatrix{}, err
		}
		matrix[i] = row
	}

	return matrix, nil
}

func isMonotonic(arr []int) bool {
	if len(arr) < 1 {
		return false
	}

	isIncreasing := arr[0] < arr[1]

	for i := 2; i < len(arr); i++ {
		if arr[i] < arr[i-1] && isIncreasing {
			return false
		}
		if arr[i] > arr[i-1] && !isIncreasing {
			return false
		}
	}

	return true
}

func (m SquareMatrix) generateSequence() []int {
	sequence := make([]int, len(m))

	for i, row := range m {
		if isMonotonic(row) {
			sequence[i] = 1
		} else {
			sequence[i] = 0
		}
	}

	return sequence
}

func task3() {
	fmt.Printf("Enter N: ")
	n, err := readInt()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var sequenceGenerator SequenceGenerator
	sequenceGenerator, err = readSquareMatrix(n)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sequence := sequenceGenerator.generateSequence()
	fmt.Print("Obtained sequence: ")
	printIntArray(sequence)
}

type Sentence string

type SentenceProcessor interface {
	removeEvenWords()
}

func readSentence() (Sentence, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a sentence: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", errors.New("failed to read sentence")
	}

	return Sentence(input), nil
}

func (s *Sentence) removeEvenWords() {
	words := strings.Fields(string(*s))
	var result []string

	for i := 0; i < len(words); i += 2 {
		result = append(result, words[i])
	}

	*s = Sentence(strings.Join(result, " "))
}

func task4() {
	var sentenceProcessor SentenceProcessor
	sentence, err := readSentence()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sentenceProcessor = &sentence
	sentenceProcessor.removeEvenWords()

	fmt.Println("Sentence after removing even words:", sentence)
}

type RuneMatrix [][]rune

type RuneMatrixProcessor interface {
	invertMatrixRows()
	String() string
}

func readRuneMatrix() (RuneMatrix, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter empty string to finish reading the matrix")
	fmt.Println("Enter the rune matrix:")

	runeMatrix := RuneMatrix{}
	for {
		row, err := reader.ReadString('\n')
		row = strings.Trim(row, " \n")

		if err != nil {
			return RuneMatrix{}, err
		}
		if len(row) == 0 {
			break
		}
		runes := []rune(row)
		runeMatrix = append(runeMatrix, runes)
	}

	return runeMatrix, nil
}

func printRuneMatrix(matrix RuneMatrix) {
	for _, row := range matrix {
		fmt.Println(string(row))
	}
}

func (matrix *RuneMatrix) invertMatrixRows() {
	for i := range *matrix {
		row := &(*matrix)[i]
		n := len((*matrix)[i])

		for j := 0; j < n/2; j++ {
			(*row)[j], (*row)[n-j-1] = (*row)[n-j-1], (*row)[j]
		}
	}
}

func (matrix *RuneMatrix) String() string {
	var result string

	for _, row := range *matrix {
		result += string(row)
	}

	return result
}

func task5() {
	var matrixProcessor RuneMatrixProcessor
	matrix, err := readRuneMatrix()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	matrixProcessor = &matrix
	matrixProcessor.invertMatrixRows()

	fmt.Println("Matrix after row inversion")
	printRuneMatrix(matrix)

	convertedString := matrixProcessor.String()
	fmt.Println("Matrix converted to string:", convertedString)
	fmt.Println("Length of the obtained string:", len(convertedString))
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

	fmt.Println("Task 4:")
	task4()
	fmt.Println(delim)

	fmt.Println("Task 5:")
	task5()
	fmt.Println(delim)
}
