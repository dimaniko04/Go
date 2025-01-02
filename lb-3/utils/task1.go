package utils

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func CreateFileOfSquareRoots(filename string, numbers []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create the file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, num := range numbers {
		if num < 0 {
			return fmt.Errorf("invalid number: square root of %d is not possible", num)
		}
		root := math.Sqrt(float64(num))
		_, err := writer.WriteString(fmt.Sprintf("%f ", root))
		if err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}
	writer.WriteRune('\n')

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to save changes to file: %w", err)
	}

	return nil
}

func sumStringComponents(line string, sum *float64) error {
	parts := strings.Fields(line)

	for _, part := range parts {
		num, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return fmt.Errorf("failed to convert number: %w", err)
		}
		*sum += num
	}
	return nil
}

func SumFileComponents(filename string) (float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to open the file: %w", err)
	}
	defer file.Close()

	sum := 0.0
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return sum, err
			}
		}
		if err = sumStringComponents(line, &sum); err != nil {
			return sum, err
		}
	}

	return sum, nil
}
