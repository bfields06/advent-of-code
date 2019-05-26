package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scnr := bufio.NewScanner(f)

	lines := []string{}
	for scnr.Scan() {
		lines = append(lines, scnr.Text())
	}

	for i, line := range lines {
		if i < len(lines) {
			for k := i + 1; k < len(lines); k++ {
				if found, key := singleDiff(lines[k], line); found {
					fmt.Printf("Found the key: %v\n", key)
				}
			}
		}

	}

}

func singleDiff(a, b string) (bool, string) {
	ar := bufio.NewScanner(strings.NewReader(a))
	ar.Split(bufio.ScanBytes)
	br := bufio.NewScanner(strings.NewReader(b))
	br.Split(bufio.ScanBytes)

	index := 0
	diff := 1
	diffIndex := 0

	for ar.Scan() {
		if string(b[index]) != ar.Text() {
			diff--
			diffIndex = index
		}

		if diff < 0 {
			return false, ""
		}
		index++

	}

	return true, a[:diffIndex] + a[diffIndex+1:]
}
