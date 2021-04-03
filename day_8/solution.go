package main

/*
	Solution for: https://adventofcode.com/2020/day/8 (part 1)
*/

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	instructions := strings.Split(string(content), "\n")
	alreadyExecutedInstructions := make(map[int]bool)
	accumulator := 0
	instructionPointer := 0

	for {

		if _, ok := alreadyExecutedInstructions[instructionPointer]; ok {
			fmt.Printf("Instruction Pointer: %d - Accumulator: %d\n", instructionPointer, accumulator)
			break
		} else {
			alreadyExecutedInstructions[instructionPointer] = true
		}

		currentInstr := instructions[instructionPointer]
		opParts := strings.Split(currentInstr, " ")
		op := opParts[0]
		value, _ := strconv.Atoi(opParts[1])

		switch op {
		case "jmp":
			instructionPointer += value
		case "nop":
			instructionPointer += 1
		case "acc":
			instructionPointer += 1
			accumulator += value
		}
	}

}
