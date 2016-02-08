package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	//Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
	pattern := regexp.MustCompile(`(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds`)
	totalTime, _ := strconv.Atoi(os.Args[1])
	reader := bufio.NewScanner(os.Stdin)
	max := 0
	winner := ""
	for reader.Scan() {
		line := reader.Text()
		match := pattern.FindStringSubmatch(line)
		who := match[1]
		speed, _ := strconv.Atoi(match[2])
		fTime, _ := strconv.Atoi(match[3])
		rTime, _ := strconv.Atoi(match[4])

		fullPeriods := totalTime / (fTime + rTime)
		partPeriods := totalTime % (fTime + rTime)
		d := speed * (fullPeriods*fTime + min(fTime, partPeriods))
		fmt.Printf("%s travels %d km in %d seconds\n", who, d, totalTime)
		if d > max {
			max = d
			winner = who

		}
	}
	fmt.Println("Winner is: ", winner, " distance: ", max)

}
