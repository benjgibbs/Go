package main

import (
	"fmt"
	"math"
)

type Item struct {
	cost, damage, armour int
}

type Items []Item

var combos []Item

const myHp = 100
const bossHp = 103
const bossDamage = 9
const bossArmour = 2

func main() {
	cheapest := 0
	found := false
	for is := findNextBest(cheapest); !found; is = findNextBest(cheapest) {
		for _, i := range is {
			myScore := myHp
			bossScore := bossHp
			for {
				bossScore -= i.damage - bossArmour
				if bossScore <= 0 {
					fmt.Println("Won, cost is:", i.cost, myScore, bossScore, i)
					found = true
					break
				}
				myScore -= bossDamage - i.armour
				if myScore <= 0 {
					fmt.Println("Lost", i.cost, myScore, bossScore)
					break
				}
			}
			cheapest = i.cost
		}
	}
}

func init() {
	weaponItems := Items{
		Item{8, 4, 0},
		Item{10, 5, 0},
		Item{25, 6, 0},
		Item{40, 7, 0},
		Item{74, 8, 0}}

	armourItems := Items{
		Item{13, 0, 1},
		Item{31, 0, 2},
		Item{53, 0, 3},
		Item{75, 0, 4},
		Item{102, 0, 5}}

	ringItems := Items{
		Item{25, 1, 0},
		Item{50, 2, 0},
		Item{100, 3, 0},
		Item{20, 0, 1},
		Item{40, 0, 2},
		Item{80, 0, 3}}

	combos = []Item{}
	for _, wd := range weaponItems {
		c1 := wd.cost
		d1 := wd.damage
		a1 := wd.armour
		combos = append(combos, Item{c1, d1, a1})
		for _, ad := range armourItems {
			c2 := c1 + ad.cost
			d2 := d1 + ad.damage
			a2 := a1 + ad.armour
			combos = append(combos, Item{c2, d2, a2})
			for r1n, r1d := range ringItems {
				c3 := c2 + r1d.cost
				d3 := d2 + r1d.damage
				a3 := a2 + r1d.armour
				combos = append(combos, Item{c3, d3, a3})
				for r2n, r2d := range ringItems {
					if r1n != r2n {
						c4 := c3 + r2d.cost
						d4 := d3 + r2d.damage
						a4 := a3 + r2d.armour
						combos = append(combos, Item{c4, d4, a4})
					}
				}
			}
		}
	}
}

func findNextBest(lastCheapest int) Items {
	result := Items{}
	nextCheapest := math.MaxInt32
	for _, i := range combos {
		if i.cost > lastCheapest {
			if i.cost == nextCheapest {
				result = append(result, i)
			} else if i.cost < nextCheapest {
				result = Items{i}
				nextCheapest = i.cost
			}
		}
	}
	return result
}
