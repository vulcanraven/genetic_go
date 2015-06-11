package genetic

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"time"
)

type Fitness interface {
	Eval([]float64) float64
}

type Ega struct {
	info       *GeneticInfo
	population []*Individual
	eval       Fitness
}

func (d *Ega) Setup(inf *GeneticInfo, evalf Fitness) error {
	rand.Seed(time.Now().UnixNano())
	d.info = inf
	d.eval = evalf
	// Double the size of the population.
	d.population = make([]*Individual, inf.population*2)
	for i := range d.population {
		var err error
		d.population[i], err = CreateIndividual(inf)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Ega) Run(generations int, printg bool) *Individual {
	// Calculate bits to mutate.
	mutatebits := int(float32(d.info.PopulationBytes()) * 8.0 * d.info.mutation)
	if printg {
		fmt.Printf("Mutating %d bits each generation.\n", mutatebits)
	}
	// Evaluate last n.
	for i := d.info.population; i < d.info.population*2; i++ {
		d.population[i].aptitude = d.eval.Eval(d.population[i].fenotype)
	}
	for g := 0; g < generations; g++ {
		// Evaluate individuals.
		for i := 0; i < d.info.population; i++ {
			d.population[i].aptitude = d.eval.Eval(d.population[i].fenotype)
		}
		// Sort by aptitude.
		sort.Sort(ByAptitude(d.population))
		// Print best aptitude individual.
		if printg {
			fmt.Println(g+1, "- Best: ", d.population[0].aptitude)
		}
		// Clone best n into the worst n of the population.
		for i := 0; i < d.info.population; i++ {
			*d.population[d.info.population+i] = *d.population[i]
		}
		// Cross over with the eclectic operator.
		for i := 0; i < d.info.population/2; i++ {
			if rand.Float32() < d.info.crossover {
				RandomAnnularCrossover(d.population[i], d.population[d.info.population-i-1], d.info)
			}
		}
		// Randomly mutate bits.
		for i := 0; i < mutatebits; i++ {
			// Mutate one random bit of an individual in the population.
			d.population[rand.Intn(d.info.population)].RandomMutateBit()
		}

		// Update fenotype.
		for i := 0; i < d.info.population; i++ {
			d.population[i].UpdateFenotype()
		}
	}
	// Evaluate individuals one last time.
	for i := 0; i < d.info.population; i++ {
		d.population[i].aptitude = d.eval.Eval(d.population[i].fenotype)
	}
	sort.Sort(ByAptitude(d.population))
	if printg {
		fmt.Println("Global Best: ", d.population[0].aptitude)
	}
	return d.population[0]
}

func evalInd(ind *Individual, evalf Fitness, c chan bool) {
	ind.aptitude = evalf.Eval(ind.fenotype)
	c <- true
}

func (d *Ega) RunConcurrent(generations, threads int, printg bool) *Individual {
	// Calculate bits to mutate.
	mutatebits := int(float32(d.info.PopulationBytes()) * 8.0 * d.info.mutation)
	if printg {
		fmt.Printf("Mutating %d bits each generation.\n", mutatebits)
	}
	runtime.GOMAXPROCS(threads)
	// Evaluate last n.
	c := make(chan bool, threads)
	for i := d.info.population; i < d.info.population*2; i++ {
		go evalInd(d.population[i], d.eval, c)
		<-c
	}
	for g := 0; g < generations; g++ {
		// Evaluate individuals.
		for i := 0; i < d.info.population; i++ {
			go evalInd(d.population[i], d.eval, c)
			<-c
		}
		// Sort by aptitude.
		sort.Sort(ByAptitude(d.population))
		// Print best aptitude individual.
		if printg {
			fmt.Println(g+1, "- Best: ", d.population[0].aptitude)
		}
		// Clone best n into the worst n of the population.
		for i := 0; i < d.info.population; i++ {
			*d.population[d.info.population+i] = *d.population[i]
		}
		// Cross over with the eclectic operator.
		for i := 0; i < d.info.population/2; i++ {
			if rand.Float32() < d.info.crossover {
				RandomAnnularCrossover(d.population[i], d.population[d.info.population-i-1], d.info)
			}
		}
		// Randomly mutate bits.
		for i := 0; i < mutatebits; i++ {
			// Mutate one random bit of an individual in the population.
			d.population[rand.Intn(d.info.population)].RandomMutateBit()
		}

		// Update fenotype.
		for i := 0; i < d.info.population; i++ {
			d.population[i].UpdateFenotype()
		}
	}
	// Evaluate individuals one last time.
	for i := 0; i < d.info.population; i++ {
		go evalInd(d.population[i], d.eval, c)
		<-c
	}
	sort.Sort(ByAptitude(d.population))
	if printg {
		fmt.Println("Global Best: ", d.population[0].aptitude)
	}
	return d.population[0]
}
