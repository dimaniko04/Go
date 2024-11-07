package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntArray []int
type IntMatrix [][]int

func ReadIntArray() (IntArray, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter space-separated integers: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		return []int{}, err
	}

	fields := strings.Fields(input)
	intArray := make(IntArray, len(fields))

	for i, field := range fields {
		if i == 20 {
			break
		}
		num, err := strconv.Atoi(field)

		if err != nil {
			return []int{}, errors.New("invalid input")
		}
		intArray[i] = num
	}

	return intArray, nil
}

func ReadIntMatrix() (IntMatrix, error) {
	fmt.Println("Enter empty string to finish reading the matrix")

	var intMatrix = IntMatrix{}

	for {
		var row, err = ReadIntArray()

		if err != nil {
			return IntMatrix{}, err
		}
		if len(row) == 0 {
			break
		}
		intMatrix = append(intMatrix, row)
	}

	return intMatrix, nil
}

func (intArray IntArray) Process() fmt.Stringer {
	var result IntArray

	for i, element := range intArray {
		if element == 0 {
			result = append(result, i)
		}
	}

	return result
}

func (intArray IntArray) String() string {
	output := ""
	for _, val := range intArray {
		output += strconv.Itoa(val) + " "
	}
	return fmt.Sprint(output)
}

func (intMatrix *IntMatrix) Process() fmt.Stringer {
	for i := 1; i < len(*intMatrix); i += 2 {
		row := (*intMatrix)[i]
		r := len(row) - 1

		for l := 0; l < r; l, r = l+1, r-1 {
			row[l], row[r] = row[r], row[l]
		}
	}
	return intMatrix
}

func (intMatrix IntMatrix) String() string {
	output := ""
	for i, row := range intMatrix {
		output += strconv.Itoa(i+1) + ". "
		for _, val := range row {
			output += strconv.Itoa(val) + " "
		}
		output += "\n"
	}
	return fmt.Sprint(output)
}
