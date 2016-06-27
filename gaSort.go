package main

import "fmt"
import "math/rand"
import "time"

type Individual []int

func (item Individual) Mutate() {
	size := len(item)
	x := rand.Intn(size)
	y := rand.Intn(size)
	item[x], item[y] = item[y], item[x]
}

func (daddy Individual) Procreate(mommy Individual) Individual {
	clone := Individual(append([]int(nil), daddy...))
	clone.Mutate()
	return clone
}

func (item Individual) FitnessScore() int {
	score := 0
	for i := 1; i < len(item); i++ {
		if item[i] > item[i-1] {
			score++
		} else {
			score--
		}
	}
	return score
}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("hello world")
}
