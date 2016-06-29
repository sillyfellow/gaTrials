package main

import "fmt"
import "math/rand"
import "time"
import "container/heap"
import "io"
import "log"
import "flag"

func readIndividual() (*Individual, bool) {
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
	return &nums, len(nums) != 0
}

func main() {
	/* set the stage  (ss@06/28/2016) */
	rand.Seed(time.Now().Unix())

	// parse and get params.
	popSize := *flag.Int("popsize", 1024, "Desired Population Size")
	genCount := *flag.Int("gencount", 128, "Desired Generations to evolve")
	flag.Parse()

	/* get the initial population ready  (ss@06/28/2016) */
	var valid bool
	pop := make(Generation, 1)
	pop[0], valid = readIndividual()
	if !valid {
		log.Fatal("The input is invalid, bbye.\n")
		return
	}
	heap.Init(&pop)

	/* we have a winner?  (ss@06/28/2016) */
	start := time.Now()
	winner, increaseGenCount := GA(&pop, popSize, genCount)
	fmt.Printf("The result is: %v; Need to increase generation count = %v\n", *winner, increaseGenCount)
	fmt.Printf("It took %s to complete.", time.Since(start))
}
