package main

import "math/rand"
import "container/heap"

type Generation []*Individual

func (gen Generation) Len() int { return len(gen) }

func (gen Generation) Less(i, j int) bool {
	return gen[i].FitnessScore() > gen[j].FitnessScore()
}

func (gen Generation) Swap(i, j int) {
	gen[i], gen[j] = gen[j], gen[i]
}

func (gen *Generation) Push(x interface{}) {
	item := x.(*Individual)
	*gen = append(*gen, item)
}

func (gen *Generation) Pop() interface{} {
	old := *gen
	n := len(old)
	item := old[n-1]
	*gen = old[0 : n-1]
	return item
}

func (gen *Generation) ResizeTo(n int) {
	orig := *gen
	*gen = orig[0:n]
}

func (gen *Generation) update() {
	heap.Fix(gen, gen.Len()-1)
}

func (gen *Generation) randomIndividual() *Individual {
	return (*gen)[rand.Intn(gen.Len())]
}

func (gen *Generation) AverageFitness() float64 {
	n := gen.Len()

	if n == 0.0 {
		return 0.0
	}

	total := float64(0)
	for i := 0; i < n; i++ {
		total += float64((*gen)[i].FitnessScore())
	}
	return total / float64(n)
}
