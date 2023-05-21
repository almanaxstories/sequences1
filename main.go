package main

import (
	"fmt"

	"github.com/almanaxstories/sequences1/functions"
)

func main() {
	var array = []int{23, 123, 523, 231, 150, 250, 7}
	n := 3
	sequences := functions.CalculateSequences(n, array)
	result := functions.FindSequenceWithMaxSum(sequences)
	fmt.Println(result)
}
