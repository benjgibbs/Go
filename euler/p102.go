package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func area(a, b, c point) float64 {
	return 0.5 * math.Abs(float64(a.x*b.y-a.x*c.y+b.x*c.y-b.x*a.y+c.x*a.y-c.x*b.y))
}

func containsX(a, b, c, x point) bool {
	const eps = 1e-12
	a1 := area(a, b, c)
	a2 := area(a, b, x)
	a3 := area(a, x, c)
	a4 := area(x, b, c)
	return math.Abs(a1-a2-a3-a4) < eps
}

func conv(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		log.Fatal(e)
	}
	return i
}

func toPoints(s string) (point, point, point) {
	p := strings.Split(s, ",")
	return point{conv(p[0]), conv(p[1])},
		point{conv(p[2]), conv(p[3])},
		point{conv(p[4]), conv(p[5])}
}

func p102() {
	file, err := os.Open("p102_triangles.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(bufio.NewReader(file))
	origin := point{0, 0}
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		a, b, c := toPoints(line)
		if containsX(a, b, c, origin) {
			count++
		}
	}
	fmt.Println("Num containing the origin: ", count)
}
