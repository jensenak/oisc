package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	ws    = 3             //wordsize = 3 bytes
	mmax  = 16383         // highest allowable addr
	ipmax = mmax - ws + 1 // highest allowable ip
	//RESERVED ADDRS
	zero     = 0
	minusOne = 1
	one      = 2
	termS    = 3
	termE    = 4
	termC    = 5
	count    = 6
	speed    = 7
	begin    = 8
)

var (
	mem [mmax]int16 // memory
	ip  int16       // instruction pointer
)

func initMem() {
	dec := json.NewDecoder(os.Stdin) // As if an OISC would magically understand JSON...
	err := dec.Decode(&mem)
	if err != nil {
		panic(err)
	}
	ip = mem[begin]
}

func subleq(a, b, c int16) int16 {
	mem[count]++
	mem[b] = mem[b] - mem[a]  // Perform subtraction
	if mem[b] <= 0 && c > 0 { // If condition triggered and c availble, jump
		ip = c
	} else {
		ip += ws //otherwise just increment to next instruction
	}
	if ip >= ipmax {
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
