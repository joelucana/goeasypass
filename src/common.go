package src

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintSlice(s []string) {
	//fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("slice: %v; len: %d; cap: %d \n", s, len(s), cap(s))
}

func If(statement bool, a, b interface{}) interface{} {
	if statement {
		return a
	}
	return b
}

func Size(size int) []struct{} {
	return make([]struct{}, size)
}

func InitRand() {
	rand.Seed(time.Now().UnixNano())
}
