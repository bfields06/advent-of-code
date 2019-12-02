package main

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

func readInts(r io.Reader) []int {
	var nums []int
	var num int

	for {
		_, err := fmt.Fscanf(r, "%d", &num)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(errors.Wrap(err, "failed to scan file for int"))
			os.Exit(1)
		}

		nums = append(nums, num)

	}

	return nums
}
func validIndex(length, j, k, l int) bool {
	return (length > j && length > k && length > l)
}
func intcode(nums []int) []int {
	codes := make([]int, len(nums))
	copy(codes, nums)
	index := 0

	length := len(codes)
	var opcode int
	var pos1 int
	var pos2 int
	var pos3 int

	for index < length {
		opcode = codes[index]
		index++
		pos1 = codes[index]
		index++
		pos2 = codes[index]
		index++
		pos3 = codes[index]
		index++

		if !validIndex(length, pos1, pos2, pos3) {
			break
		}

		if opcode == 1 {
			codes[pos3] = codes[pos1] + codes[pos2]
		}

		if opcode == 2 {
			codes[pos3] = codes[pos1] * codes[pos2]
		}

		if opcode == 99 {
			break
		}
	}

	return codes
}
func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to open file"))
		os.Exit(1)
	}

	nums := readInts(f)

	fmt.Println(intcode(nums))
	var r []int
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			nums[1] = a
			nums[2] = b

			r = intcode(nums)
			if r[0] == 19690720 {
				fmt.Printf("Found it! %d, %d, %d\n", r[1], r[2], 100*r[1]+r[2])
				break
			}
		}

	}

	// fmt.Println(r)

}
