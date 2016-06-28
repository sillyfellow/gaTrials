package main

import "container/heap"

func GA(population *Generation, popSize, count int) (*Individual, bool) {
	averageFitness := -999.0
	for count > 0 {
		for i := 0; i < popSize; i++ {
			randIndividual := population.randomIndividual()
			newChild := randIndividual.Procreate()
			heap.Push(population, newChild)
		}

		population.update()
		population.ResizeTo(popSize)

		newAverageFitness := population.AverageFitness()
		if newAverageFitness == averageFitness {
			break
		}
		averageFitness = newAverageFitness
		count--
	}
	return (*population)[0], count == 0
}
