package main

import (
	"fmt"
	"time"
)

const (
	ws   = 3  //wordsize = 3 bytes
	mmax = 60 // highest allowable ip
)

var mem [64]int8 // memory
var ip int8      // instruction pointer
var count int8

func initMem() {
	mem = [64]int8{0, -1, 0,
		3, 5, 0, // Things you want to multiply go in these :)
		0, 0, 0,
		3, 6, 0,
		4, 7, 0,
		6, 7, 30,
		6, 6, 0,
		4, 6, 0,
		7, 7, 0,
		3, 7, 36,
		7, 7, 0,
		4, 7, 0,
		1, 6, 0,
		7, 2, 0,
		1, 6, 39,
		0, 0, 63}

	ip = 9
}

func subleq(a, b, c int8) int8 {
	count++
	mem[b] = mem[b] - mem[a]  // Perform subtraction
	if mem[b] <= 0 && c > 0 { // If condition triggered and c availble, jump
		ip = c
	} else {
		ip += ws //otherwise just increment to next instruction
	}
	if ip >= mmax {
		return -1 // End program if we're going to overflow the instruction pointer
	}
	return ip
}

func terminal(start, end, cols int8) {
	fmt.Print("\033[2J") // Clear screen
	line := 2
	col := int8(0)
	space := int8(12)
	for i := start; i <= end; i++ {
		fmt.Printf("\033[%d;%dH", line, col)
		if i >= ip && i <= ip+2 {
			fmt.Print("\033[1;32m")
		}
		if i == mem[ip+1] {
			fmt.Print("\033[1;34m")
		}
		fmt.Printf("0x%d:%d", i, mem[i])
		fmt.Print("\033[0m")
		col += space
		if col >= (cols * space) {
			col = 0
			line++
		}
	}
}

func clock() {
	time.Sleep(time.Second / 2)
}

func main() {
	initMem() // Fill memory (program, vars, etc)
	for ip >= 0 {
		terminal(0, 63, 3)
		ip = subleq(mem[ip], mem[ip+1], mem[ip+2])
		clock()
	}
}
