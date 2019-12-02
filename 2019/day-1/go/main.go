package main

import (
	"fmt"
	fcu "github.com/bfields06/advent-of-code/2019/day-1/go/fuelcounterupper"
	"github.com/bfields06/advent-of-code/helpers"
	"io"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func fuelRequired(nums []int) int {
	var transform []int
	for _, val := range nums {
		transform = append(transform, fcu.CalculateFuelByMass(val))
	}

	return helpers.SumCollection(transform)
}

func fuelRequiredWithExtra(nums []int) int {
	var transform []int

	for _, val := range nums {
		val = fcu.CalculateFuelByMass(val)

		for val > 0 {
			transform = append(transform, val)
			val = fcu.CalculateFuelByMass(val)
		}
	}

	return helpers.SumCollection(transform)
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
	fmt.Printf("Part 2 Fuel Required: %d \n", fuelRequiredWithExtra(nums))

}
