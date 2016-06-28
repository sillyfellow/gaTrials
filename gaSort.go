package main

import "fmt"
import "math/rand"
import "time"
import "container/heap"
import "io"
import "log"
import "flag"

type Individual []int
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

func readIndividual() *Individual {
	fmt.Printf("Keep giving integers until you are out ;). Press Ctrl-D to finish\n")
	nums := make(Individual, 0)
	var d int
	for {
		_, err := fmt.Scan(&d)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		nums = append(nums, d)
	}
	return &nums
}

func main() {
	/* set the stage  (ss@06/28/2016) */
	rand.Seed(time.Now().Unix())

	// parse and get params.
	popSize := *flag.Int("popsize", 1024, "Desired Population Size")
	genCount := *flag.Int("gencount", 128, "Desired Generations to evolve")
	flag.Parse()

	/* get the initial population ready  (ss@06/28/2016) */
	pop := make(Generation, 1)
	pop[0] = readIndividual()
	heap.Init(&pop)

	/* we have a winner?  (ss@06/28/2016) */
	winner, increaseGenCount := GA(&pop, popSize, genCount)
	fmt.Printf("The result is: %v; Need to increase generation count = %v\n", *winner, increaseGenCount)
}
