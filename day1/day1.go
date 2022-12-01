package day1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func PartOne(input io.Reader) int64 {
	fileScanner := bufio.NewScanner(input)
	var biggestCalories int64 = 0
	var calorieCounter int64 = 0

	for fileScanner.Scan() {
		calorie, err := strconv.ParseInt(fileScanner.Text(), 0, 0)

		if err != nil {
			if calorieCounter > biggestCalories {
				biggestCalories = calorieCounter
			}
			calorieCounter = 0

			continue
		}

		calorieCounter += calorie
	}

	return biggestCalories
}

func PartTwo(input io.Reader) int64 {
	fileScanner := bufio.NewScanner(input)
	var calorieCounter int64 = 0

	top3 := [3]int64{0, 0, 0}
	for fileScanner.Scan() {
		calorie, err := strconv.ParseInt(fileScanner.Text(), 0, 0)
		if err != nil {
			if calorieCounter > top3[0] {
				top3[0], top3[1], top3[2] = calorieCounter, top3[0], top3[1]
			} else if calorieCounter > top3[1] {
				top3[1], top3[2] = calorieCounter, top3[1]
			} else if calorieCounter > top3[2] {
				top3[2] = calorieCounter
			}
			calorieCounter = 0

			continue
		}

		calorieCounter += calorie
	}

	return top3[0] + top3[1] + top3[2]
}

func Run() {

	inputFile, err := os.Open("input/day_1.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	inputBytes, err := io.ReadAll(inputFile)
	inputFile.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	input := string(inputBytes)

	fmt.Printf("Day One Part One Result: %d\n", PartOne(strings.NewReader(input)))
	fmt.Printf("Day One Part One Result: %d\n", PartTwo(strings.NewReader(input)))
}
