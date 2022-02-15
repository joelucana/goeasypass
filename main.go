package main

import (
	"fmt"
	"github.com/joelucana/goeasypass/src"
)

func main() {
	src.InitRand()
	groups := 3
	fmt.Println(src.GeneratePass(groups))
}
