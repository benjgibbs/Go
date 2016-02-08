package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var a, b int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prog := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		prog = append(prog, line)
	}

	a = 1
	b = 0

	for i := 0; i < len(prog); {
		inst := prog[i]
		fmt.Println(a, b)
		fmt.Println(i, inst)
		switch inst[:3] {
		case "hlf":
			reg := sel(inst[4:])
			*reg /= 2
			i++
		case "inc":
			reg := sel(inst[4:])
			*reg++
			i++
		case "tpl":
			reg := sel(inst[4:])
			*reg *= 3
			i++
		case "jmp":
			v, _ := strconv.Atoi(inst[4:])
			i += v
		case "jie":
			reg := sel(inst[4:5])
			if *reg%2 == 0 {
				i += offset(inst[7:])
			} else {
				i++
			}
		case "jio":
			reg := sel(inst[4:5])
			if *reg == 1 {
				i += offset(inst[7:])
			} else {
				i++
			}
		}
	}
	fmt.Println("Registers: a=", a, "b=", b)
}

func offset(s string) int {
	if s[0] == '+' {
		s = s[1:]
	}
	num, e := strconv.Atoi(s)
	if e != nil {
		panic("Bad number: " + s)
	}
	return num
}

func sel(s string) *int {
	switch s {
	case "a":
		return &a
	case "b":
		return &b
	}
	panic("Unknown register" + s)
}
