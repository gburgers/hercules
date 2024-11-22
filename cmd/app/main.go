package main

import (
	"fmt"

	"github.com/gburgers/hercules/internal/math"
)

func main() {
	fmt.Println("Hello, Go!")
	p := math.Numbers{Age: 5, Number: 5}
	fmt.Println(math.Times(p))
}
