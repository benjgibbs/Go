package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Nodes map[string]Node

var cache map[string]uint
var nodes Nodes

type Node interface {
	apply() uint
}

type And struct {
	in1, in2 string
}

func (a And) apply() uint {
	v1 := deref(a.in1)
	v2 := deref(a.in2)
	return v1 & v2
}

type Or struct {
	in1, in2 string
}

func (a Or) apply() uint {
	v1 := deref(a.in1)
	v2 := deref(a.in2)
	return v1 | v2
}

type LShift struct {
	shift uint
	in    string
}

func (s LShift) apply() uint {
	return deref(s.in) << s.shift
}

type RShift struct {
	shift uint
	in    string
}

func (s RShift) apply() uint {
	return deref(s.in) >> s.shift
}

type Not struct {
	in string
}

func (n Not) apply() uint {
	return deref(n.in) ^ 0xFFFF
}

type Input struct {
	v string
}

func (s Input) apply() uint {
	return deref(s.v)
}

func main() {
	nodes = make(Nodes)
	cache = make(map[string]uint)
	parseInput()
	for k, v := range nodes {
		fmt.Println(k, ":=", v.apply())
	}
	fmt.Println("a :=", deref("a"))

}

func parseInput() {
	assign := regexp.MustCompile(`^(\w+) -> (\w+)$`)
	andOr := regexp.MustCompile(`(\w+) (AND|OR) (\w+) -> (\w+)`)
	shift := regexp.MustCompile(`(\w+) ((L|R)SHIFT) (\d+) -> (\w+)`)
	not := regexp.MustCompile(`NOT (\w+) -> (\w+)`)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		match := assign.FindStringSubmatch(line)
		if len(match) > 0 {
			val := match[1]
			id := match[2]
			nodes[id] = Input{val}
			continue
		}
		match = andOr.FindStringSubmatch(line)
		if len(match) > 0 {
			in1 := match[1]
			in2 := match[3]
			out := match[4]
			if match[2] == "AND" {
				nodes[out] = And{in1, in2}
			} else {
				nodes[out] = Or{in1, in2}
			}
			continue
		}
		match = shift.FindStringSubmatch(line)
		if len(match) > 0 {
			in := match[1]
			s, _ := strconv.Atoi(match[4])
			out := match[5]
			if match[3] == "L" {
				nodes[out] = LShift{uint(s), in}
			} else {
				nodes[out] = RShift{uint(s), in}
			}
			continue
		}
		match = not.FindStringSubmatch(line)
		if len(match) > 0 {
			in := match[1]
			out := match[2]
			nodes[out] = Not{in}
			continue
		}
		fmt.Println("Error.  Unmatched line: ", line)
	}
}

func deref(s string) uint {
	if cLookup, ok := cache[s]; ok {
		return cLookup
	}
	n := nodes[s]
	if n != nil {
		res := n.apply()
		cache[s] = res
		return res
	}
	i, err := strconv.Atoi(s)
	if err == nil {
		res := uint(i)
		cache[s] = res
		return res
	}
	panic("Failed to find: " + s)
}
