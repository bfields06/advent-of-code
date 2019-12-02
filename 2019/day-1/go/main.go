package main

import (
	"fmt"
	"io"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func calculateFuelByMass(mass int) int {
	ret := mass/3 - 2
	if ret > 0 {
		return ret
	}

	return 0
}

func sumCollection(mass []int) int {
	total := 0
	for _, m := range mass {
		total += m
	}

	return total
}

func fuelRequired(nums []int) int {
	var transform []int
	for _, val := range nums {
		transform = append(transform, calculateFuelByMass(val))
	}

	return sumCollection(transform)
}

func fuelRequiredWithExtra(nums []int) int {
	var transform []int

	for _, val := range nums {
		val = calculateFuelByMass(val)

		for val > 0 {
			transform = append(transform, val)
			val = calculateFuelByMass(val)
		}
	}

	return sumCollection(transform)
}

func readInts(r io.Reader) []int {
	var perline int
	var nums []int

	for {
		_, err := fmt.Fscanf(r, "%d\n", &perline)

		if err != nil {

			if err == io.EOF {
				break //stop reading file
			}

			fmt.Println(err)
			os.Exit(1)

		}
		nums = append(nums, perline)
	}

	return nums
}

func main() {
	fileName := "input.txt"

	f, err := os.Open(fileName)
	checkError(err)

	nums := readInts(f)

	//fmt.Println(nums)
	fmt.Printf("Part 1 Fuel Required: %d\n", fuelRequired(nums))
	fmt.Printf("Part 2 Fuel Required: %d \n", sumCollection(fuelRequiredWithExtra(nums)))

}
