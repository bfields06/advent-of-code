package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	s := bufio.NewScanner(f)

	var val int
	var vals []int
	for s.Scan() {
		fmt.Sscanf(s.Text(), "%d", &val)
		vals = append(vals, val)
	}

	fmt.Println(vals)

	if err = s.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	seen := map[int]bool{0: true}
	for {
		for _, val = range vals {
			sum += val

			if seen[sum] {
				fmt.Printf("We've found the repeate offender: %d\n", sum)
				return
			}
			seen[sum] = true
		}
	}

}
