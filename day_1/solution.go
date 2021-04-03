package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	numbers := make(map[int]int)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if val, ok := numbers[num]; ok {
			fmt.Printf("Solution = %d\n", num*val)
		} else {
			numbers[2020-num] = num
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
