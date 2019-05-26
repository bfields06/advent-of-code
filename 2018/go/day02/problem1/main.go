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
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)

	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	triple := 0
	double := 0

	for _, line := range lines {
		letterCount := map[rune]int{}
		for _, letter := range line {
			//fmt.Println(string(letter))
			letterCount[letter]++
		}
		fmt.Println(line)
		foundDouble := false
		foundTriple := false
		for k, v := range letterCount {

			if v == 3 {
				fmt.Printf("3 %s's\n", string(k))
				if !foundTriple {
					triple++
					foundTriple = true
				}

			}

			if v == 2 {
				fmt.Printf("2 %s's\n", string(k))
				if !foundDouble {
					double++
					foundDouble = true
				}
			}
		}

	}

	fmt.Printf("Double:%v * Triple:%v = %v\n", double, triple, double*triple)
}
