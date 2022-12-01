package day1

import (
	"strings"
	"testing"
)

var input string = `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000

`

func TestPartOne(t *testing.T) {

	inputData := strings.NewReader(input)

	got := PartOne(inputData)
	var expected int64 = 24000

	if got != expected {
		t.Errorf("Bruh! Not the same %v != %v", got, expected)
	}

}

func TestPartTwo(t *testing.T) {

	inputData := strings.NewReader(input)

	got := PartTwo(inputData)
	var expected int64 = 24000 + 11000 + 10000

	if got != expected {
		t.Errorf("Bruh! Not the same %v != %v", got, expected)
	}

}
