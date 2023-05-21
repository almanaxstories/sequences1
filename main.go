package main

import (
	"fmt"

	"github.com/almanaxstories/sequences1/functions"
)

func main() {
	var array = []int{23, 123, 523, 231, 150, 250, 7}
	n := 3

	sequences, err := functions.CalculateSequences(n, array)
	if err != nil {
		fmt.Printf("An error occurred: %s", err)
	}

	result := functions.FindSequenceWithMaxSum(sequences)
	fmt.Println(result)
}
