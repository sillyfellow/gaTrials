package main

import "math/rand"

type Individual []int

func (item Individual) Mutate() {
	size := len(item)
	x := rand.Intn(size)
	y := rand.Intn(size)
	item[x], item[y] = item[y], item[x]
}

func (daddy Individual) Procreate() *Individual {
	clone := Individual(append([]int(nil), daddy...))
	clone.Mutate()
	return &clone
}

func (item *Individual) FitnessScore() int {
	// TODO : memoize this  -- (ss@06/28/2016) --

	if item == nil {
		return 0
	}

	individual := *item
	score := 0
	for i := 1; i < len(individual); i++ {
		if individual[i] > individual[i-1] {
			score++
		} else {
			score--
		}
	}
	return score
}
