package main

import (
	"fmt"
	"math"
	"slices"
	"time"
)

type User struct {
	name      string
	address   string
	birthDate time.Time
}

type Character struct {
	User
	x      float64
	y      float64
	rating int
}

func (ch Character) String() string {
	return fmt.Sprintf("%s, x:%f, y:%f, rating:%d", ch.name, ch.x, ch.y, ch.rating)
}

var questCoordinates = struct {
	x float64
	y float64
}{
	x: 101,
	y: 204,
}

func PrintDistance(x1, y1, x2, y2 float64) {
	if x2 == 0 && y2 == 0 {
		x2 = questCoordinates.x
		y2 = questCoordinates.y
	}
	distance := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	fmt.Printf("Distance between (%.2f:%.2f) and (%.2f:%.2f) - %f\n", x1, y1, x2, y2, distance)
}

func MinValue(characters ...Character) Character {
	return slices.MinFunc(characters, func(l, r Character) int {
		return l.rating - r.rating
	})
}

type Characters []Character

func (characters Characters) String() string {
	output := ""
	for _, ch := range characters {
		output += fmt.Sprintf("\t%s\n", ch)
	}
	return output
}

func (characters *Characters) Delete(name string) {
	index := -1

	for i, character := range *characters {
		if character.name == name {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}

	*characters = append((*characters)[:index], (*characters)[index+1:]...)
}

type Equation struct {
	lowerBound float64
	upperBound float64
}

type IEquation interface {
	newtonMethod(f func(float64) float64, derivative func(float64) float64) float64
}

func (eq Equation) newtonMethod(f func(float64) float64, derivative func(float64) float64) float64 {
	x0 := (eq.lowerBound + eq.upperBound) / 2

	for {
		x := x0 - f(x0)/derivative(x0)

		if math.Abs(x-x0) < 1e-5 {
			return x
		}
		x0 = x
	}
}
