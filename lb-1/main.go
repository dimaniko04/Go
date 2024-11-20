package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readThreeFloat64() (float64, float64, float64, error) {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)

	if len(parts) != 3 {
		return 0, 0, 0, errors.New("expected three real numbers")
	}

	a, err1 := strconv.ParseFloat(parts[0], 64)
	b, err2 := strconv.ParseFloat(parts[1], 64)
	c, err3 := strconv.ParseFloat(parts[2], 64)

	if err1 != nil || err2 != nil || err3 != nil {
		return 0, 0, 0, errors.New("expected three real numbers")
	}

	return a, b, c, nil
}

func u(x, y, z float64) float64 {
	var max float64 = x
	var min float64 = x

	if y > max {
		max = y
	} else if y < min {
		min = y
	}
	if z > max {
		max = z
	} else if z < min {
		min = z
	}

	if min == 0 {
		return math.NaN()
	}

	numerator := math.Pow(max, 2) - math.Pow(2, x)*min
	denominator := math.Sin(2*x) + max/min

	if denominator == 0 {
		return math.NaN()
	}
	return numerator / denominator
}

func task1() {
	fmt.Print("Enter values for x, y, z: ")
	x, y, z, err := readThreeFloat64()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("u =", u(x, y, z))
}

func convertToMeters(unit int, length float64) (float64, error) {
	unitToMeters := map[int]float64{
		1: 0.001, //millimeters -> meters
		2: 0.01,  //centimeters -> meters
		3: 1,     //meters -> meters
		4: 0.1,   //decimeters -> meters
		5: 1000,  //kilometers -> meters
	}

	if coefficient, exists := unitToMeters[unit]; exists {
		return length * coefficient, nil
	}
	return -1, errors.New("unknown unit of measurement")
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

func readFloat64() (float64, error) {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, errors.New("expected valid real number")
	}

	return num, nil
}

func task2() {
	fmt.Println("Enter the unit number:")
	fmt.Println("\t1 - millimeters")
	fmt.Println("\t2 - centimeters")
	fmt.Println("\t3 - meters")
	fmt.Println("\t4 - decimeters")
	fmt.Println("\t5 - kilometers")

	unit, err := readInt()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("Enter the length: ")

	length, err := readFloat64()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result, err := convertToMeters(unit, length)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Length in meters: %.4f m\n", result)
}

func equationFactory(a, b float64) func(float64) float64 {
	return func(x float64) float64 {
		if x < b {
			return math.Sin(math.Abs(a*x + math.Pow(b, a)))
		} else {
			return math.Cos(math.Abs(a*x - math.Pow(b, a)))
		}
	}
}

func task3() {
	tableBorder := strings.Repeat("-", 17)

	step := 0.1
	leftBound := 1.0
	rightBound := 3.0
	equation := equationFactory(1.5, 2)

	fmt.Println(tableBorder)
	fmt.Println("|  x  |    y    |")
	for i := leftBound; i < rightBound+step; i += step {
		fmt.Printf("| %.1f | %7.4f |\n", i, equation(i))
	}
	fmt.Println(tableBorder)
}

func removeMinElement(arr []int) []int {
	minIndex := 0

	for i := range arr {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}
	arr = append(arr[:minIndex], arr[minIndex+1:]...)

	return arr
}

func insertElements(arr []int, k int, elements ...int) []int {
	if k > len(arr) {
		k = len(arr)
	}
	arr = append(arr[:k], append(elements, arr[k:]...)...)
	return arr
}

func swapEvenOddIndices(arr []int) {
	for i := 0; i < len(arr)-1; i += 2 {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}
}

func findIndex(arr []int, val int) int {
	for i, el := range arr {
		if el == val {
			return i
		}
	}
	return -1
}

func readIntArray() ([]int, error) {
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

func task4() {
	fmt.Print("Enter initial array: ")
	arr, err := readIntArray()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	arr = removeMinElement(arr)
	fmt.Println("After removing the minimum element:", arr)

	fmt.Print("Enter elements to insert: ")
	elements, err := readIntArray()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("Enter the position where to insert the elements: ")
	k, err := readInt()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	arr = insertElements(arr, k, elements...)
	fmt.Println("After inserting elements:", arr)

	swapEvenOddIndices(arr)
	fmt.Println("After swapping even and odd indices:", arr)

	fmt.Print("Enter the value you want to find: ")
	key, err := readInt()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	index := findIndex(arr, key)
	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", key, index)
	} else {
		fmt.Printf("Element %d not found\n", key)
	}
}

type ProcessedColumn struct {
	columnIndex   int
	geometricMean float64
}

func processMatrix(matrix [][]int) []ProcessedColumn {
	var result = []ProcessedColumn{}

	for j := 0; j < len(matrix[0]); j++ {
		prod := 1
		prodEven := 1
		for i := 0; i < len(matrix); i++ {
			if matrix[i][j]%2 == 0 {
				prodEven *= matrix[i][j]
			}
			if matrix[i][j]%3 == 0 {
				prod *= matrix[i][j]
			}
		}
		if prodEven%4 != 0 && prod != 1 {
			geometricMean := math.Pow(float64(prod), 1.0/float64(len(matrix)))

			result = append(result, ProcessedColumn{
				columnIndex:   j,
				geometricMean: geometricMean,
			})
		}
	}
	return result
}

func task5() {
	matrix := [][]int{
		{6, 7, 12, 5},
		{9, 14, 3, 9},
		{15, 2, 18, 11},
		{3, 4, 6, 15},
		{5, 16, 15, 30},
	}

	fmt.Println("Initial matrix:")
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%2d ", val)
		}
		fmt.Println()
	}

	processedMatrix := processMatrix(matrix)

	if len(processedMatrix) == 0 {
		fmt.Println("No column with product of even elements not divisible by 4 found")
	} else {
		for _, column := range processedMatrix {
			fmt.Printf("Column %d - Geometric Mean of multiples of elements divisible by 3: %.4f\n", column.columnIndex+1, column.geometricMean)
		}
	}

}

func sortWordLetters(str string) string {
	words := strings.Fields(str)

	for i, word := range words {
		letters := []rune(word)
		sort.Slice(letters, func(i, j int) bool {
			return letters[i] < letters[j]
		})
		words[i] = string(letters)
	}

	return strings.Join(words, " ")
}

func task6() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a sentence: ")
	input, _ := reader.ReadString('\n')

	sortedString := sortWordLetters(input)
	fmt.Println("Letters in words sorted alphabetically:", sortedString)
}

type Person struct {
	name string
	age  int
}

func filterUnder18(people []Person) []Person {
	var result []Person
	for _, person := range people {
		if person.age < 18 {
			result = append(result, person)
		}
	}
	return result
}

func removeAtIndex(arr []string, index int) ([]string, error) {
	if index < 0 || index >= len(arr) {
		return arr, errors.New("index out of range")
	}
	return append(arr[:index], arr[index+1:]...), nil
}

func task7() {
	people := []Person{
		{"John", 17},
		{"Alice", 22},
		{"Bob", 16},
		{"Charlie", 19},
	}

	fmt.Println("Initial array:")
	for _, person := range people {
		fmt.Printf("%s, Age: %d\n", person.name, person.age)
	}

	under18 := filterUnder18(people)

	fmt.Println("\nUnder 18:")
	for _, person := range under18 {
		fmt.Printf("%s, Age: %d\n", person.name, person.age)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a sentence: ")
	input, _ := reader.ReadString('\n')

	words := strings.Fields(input)
	fmt.Println("Words:", words)

	fmt.Print("Enter the index of word to delete: ")
	k, err := readInt()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	words, err = removeAtIndex(words, k)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Sentence after word removal:", strings.Join(words, " "))
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

	pause()
	fmt.Println("Task 7:")
	task7()
	fmt.Println(delim)
}
