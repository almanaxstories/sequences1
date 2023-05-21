package functions

import (
	"strconv"
)

type Sequence struct {
	Name  string
	Value int
}

func CalculateSequences(sequenceLength int, data []int) []Sequence {
	sequences := []Sequence{}
	for mainIterator := 0; mainIterator <= len(data)-sequenceLength; mainIterator++ {
		sequenceName := "Sequence:"
		sequenceSum := 0
		for subIterator := mainIterator; subIterator < mainIterator+sequenceLength; subIterator++ {
			sequenceName += " " + strconv.Itoa(data[subIterator])
			sequenceSum += data[subIterator]
		}
		sequenceName += " " + "With total sum: "
		sequences = append(sequences, Sequence{Name: sequenceName, Value: sequenceSum})
	}
	return sequences
}

func FindSequenceWithMaxSum(data []Sequence) Sequence {
	maxVal := data[0]
	for _, value := range data {
		if value.Value > maxVal.Value {
			maxVal = value
		}
	}
	return maxVal
}
