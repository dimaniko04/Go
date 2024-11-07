package main

import "fmt"

type Processable interface {
	Process() fmt.Stringer
	String() string
}
