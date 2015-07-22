package main

import (
	"fmt"
	"time"
)

const (
	ws   = 3     //wordsize = 3 bytes
	mmax = 16381 // highest allowable ip
	//RESERVED ADDRS
	zero     = 0
	minusOne = 1
	one      = 2
	termS    = 3
	termE    = 4
	termC    = 5
	count    = 6
	speed    = 7
	result   = 8
)

var (
	mem [mmax + ws - 1]int16 // memory
	ip  int16                // instruction pointer
)

func initMem() {
	a := int16(9)  // Multiply value in this address...
	b := int16(10) // By value here
	ma := int16(12)
	mb := int16(13)
	mem = [mmax + ws - 1]int16{0, 0, 0, // Reserved addrs
		0, 0, 0, // Reserved addrs
		0, 0, 0, // Reserved addrs
		9, 5, 0, // Things you want to multiply go in these :)
		0, 0, 0,
		a, ma, 0, //12
		b, mb, 0,
		ma, mb, 36, //18
		ma, ma, 0,
		b, ma, 0, //24
		mb, mb, 0,
		a, mb, 42, //30
		mb, mb, 0,
		b, mb, 0, //36
		minusOne, ma, 0,
		mb, result, 0, //42
		minusOne, ma, 45,
		0, 0, mmax}

	mem[zero] = 0
	mem[minusOne] = -1
	mem[one] = 1
	mem[termS] = 0
	mem[termE] = 63
	mem[termC] = 1
	mem[count] = 0
	mem[speed] = 500
	mem[result] = 0

	ip = 15
}

func subleq(a, b, c int16) int16 {
	mem[count]++
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

func terminal(start, end, cols int16) {
	cols = cols * ws     // Each column contains a word
	fmt.Print("\033[2J") // Clear screen
	line := 2
	col := int16(0)
	space := int16(20)
	for i := start; i <= end; i++ {
		fmt.Printf("\033[%d;%dH", line, col)
		if i >= ip && i <= ip+2 {
			fmt.Print("\033[1;32m")
		}
		if i == mem[ip+1] {
			fmt.Print("\033[1;34m")
		}
		if i == result {
			fmt.Print("\033[1;35m")
		}
		if i == count {
			fmt.Print("\033[1;33m")
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
	time.Sleep(time.Millisecond * time.Duration(mem[speed]))
}

func main() {
	initMem() // Fill memory (program, vars, etc)
	for ip >= 0 {
		terminal(mem[termS], mem[termE], mem[termC])
		ip = subleq(mem[ip], mem[ip+1], mem[ip+2])
		clock()
	}
}
