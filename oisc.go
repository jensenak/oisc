package main

import "fmt"

const ws = 3 //wordsize = 3 bytes

var mem map[int8]int8 // memory
var ip int8           // instruction pointer

func initMem() {
	// 10 + 12 - 15 + 3 = 10 :)
	mem[0] = 10 // we're going to do some arithmetic
	mem[1] = 15 // ... and these are our vars
	mem[2] = 12
	mem[3] = 3

	mem[4] = 0 // This is going to serve as our result
	mem[5] = 0 // Finish off the wordsize

	ip = 2
}

func subleq(a, b, c int8) int8 {
	mem[b] = mem[b] - mem[a]        // Perform subtraction
	if mem[b] <= 0 && mem[c] >= 0 { // If condition triggered and c availble, jump
		return c
	}
	if ip >= 127-ws {
		return -1 // End program if we're going to overflow the instruction pointer
	}
	ip += ws
	return ip //otherwise just increment to next instruction
}

func main() {
	initMem() // Fill memory (program, vars, etc)
	for ip >= 0 {
		ip = subleq(mem[ip], mem[ip+1], mem[ip+2])
	}
	for w := range mem {
		fmt.Printf("0x%d %d", w, mem[w])
	}
}
