package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type DVD struct {
	Title    string
	Director string
	Duration int
	Price    float64
}

func WriteDVDsToFile(filename string, dvds []DVD) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, dvd := range dvds {
		_, err := fmt.Fprintf(writer,
			"%s;%s;%d;%f\n",
			dvd.Title,
			dvd.Director,
			dvd.Duration,
			dvd.Price)

		if err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to save changes to file: %w", err)
	}

	return nil
}

func ReadDVDsFromFile(filename string) ([]DVD, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open the file: %w", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var dvds []DVD
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		parts := strings.Split(line, ";")

		if len(parts) != 4 {
			return nil, errors.New("invalid data format")
		}
		duration, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, fmt.Errorf("invalid duration format: %w", err)
		}
		price, err := strconv.ParseFloat(parts[3], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid price format: %w", err)
		}
		dvd := DVD{
			Title:    parts[0],
			Director: parts[1],
			Duration: duration,
			Price:    price,
		}
		dvds = append(dvds, dvd)
	}

	return dvds, nil
}

func DeleteExpensiveDVDs(filename string, maxPrice float64) error {
	dvds, err := ReadDVDsFromFile(filename)
	if err != nil {
		return err
	}

	filtered := []DVD{}
	for _, dvd := range dvds {
		if dvd.Price <= maxPrice {
			filtered = append(filtered, dvd)
		}
	}

	return WriteDVDsToFile(filename, filtered)
}

func AddDVD(filename string, dvd DVD, k int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to insert element at position %d: %w", k, r.(error))
		}
		return
	}()

	dvds, err := ReadDVDsFromFile(filename)
	if err != nil {
		return err
	}

	dvds = append(dvds[:k+1], dvds[k:]...)
	dvds[k] = dvd

	return WriteDVDsToFile(filename, dvds)
}

func PrintFileContent(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open the file: %w", err)
	}
	defer file.Close()

	io.Copy(os.Stdout, file)

	return nil
}
