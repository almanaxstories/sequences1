package functions

import (
	"strconv"
)

type Sequence struct {
	Name  string
	Value int
}

type Error interface {
	Error() string
}

type mismatchedLength struct{}

func (m *mismatchedLength) Error() string {
	return "The length of sequence is larger than the input array's one"
}

func CalculateSequences(sequenceLength int, data []int) ([]Sequence, error) {

	if sequenceLength > len(data) {
		return nil, &mismatchedLength{}
	}

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
	return sequences, nil
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
