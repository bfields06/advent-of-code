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
	sum := 0
	for s.Scan() {
		fmt.Sscanf(s.Text(), "%d", &val)
		sum += val
	}

	fmt.Printf("Calibration: %d\n", sum)

	if err = s.Err(); err != nil {
		log.Fatal(err)
	}

}
