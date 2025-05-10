package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func Avg(nums []float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total / float64(len(nums))
}

func Mediana(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	sort.Float64s(nums)
	i := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[i]
	}
	return (nums[i-1] + nums[i]) / 2
}

func Variance(nums []float64) float64 {
	avg := Avg(nums)
	variance := 0.0
	for _, num := range nums {
		variance += (num - avg) * (num - avg)
	}
	return variance
}

func Standard_Deviation(nums []float64) float64 {
	return math.Sqrt(Variance(nums))
}

func main() {
	filename := "data.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	defer file.Close()

	var numbers []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal("Error parsing number: ", err)
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file: ", err)
	}

	fmt.Printf("Avg: %.2f\n", Avg(numbers))
	fmt.Printf("Mediana: %.2f\n", Mediana(numbers))
	fmt.Printf("Variance: %.2f\n", Variance(numbers))
	fmt.Printf("Standard Deviation: %.2f\n", Standard_Deviation(numbers))
}
