package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileContents, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(fileContents), "\n")
	lines = append(lines, "")
	fmt.Printf("Part 1: %v\n", partOne(lines))
	fmt.Printf("Part 2: %v\n", partTwo(lines))
}

func partOne(lines []string) int {
	totalCalories, maxCalories := 0, 0
	for _, calories := range lines {
		cal, _ := strconv.Atoi(calories)
		totalCalories += cal
		if calories == "" {
			maxCalories = int(math.Max(float64(maxCalories), float64(totalCalories)))
			totalCalories = 0
		}
	}
	return maxCalories
}

func partTwo(lines []string) int {
	totalCalories := 0
	maxCalories := [3]int{0, 0, 0}
	for _, calories := range lines {
		cal, _ := strconv.Atoi(calories)
		totalCalories += cal
		if calories == "" {
			for i := range maxCalories {
				if totalCalories >= maxCalories[i] {
					maxCalories[i], totalCalories = totalCalories, maxCalories[i]
				}
			}
			totalCalories = 0
		}
	}
	return maxCalories[0] + maxCalories[1] + maxCalories[2]
}
