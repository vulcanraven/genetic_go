package genetic

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Fitness interface {
	Eval([]float64) float64
}

type Ega struct {
	info       *GeneticInfo
	population []*Individual
	eval       *Fitness
}

func (d *Ega) Setup(inf *GeneticInfo, evalf *Fitness) error {
	rand.NewSource(time.Now().UnixNano())
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
	// Evaluate last n.
	for i := d.info.population; i < d.info.population*2; i++ {
		d.population[i].aptitude = (*d.eval).Eval(d.population[i].fenotype)
	}
	return nil
}

func (d *Ega) Run(generations int) *Individual {
	// Calculate bits to mutate.
	mutatebits := int(float32(d.info.PopulationBytes()) * 8.0 * d.info.mutation)
	for g := 0; g < generations; g++ {
		// Evaluate individuals.
		for i := 0; i < d.info.population; i++ {
			d.population[i].aptitude = (*d.eval).Eval(d.population[i].fenotype)
		}
		// Sort by aptitude.
		sort.Sort(ByAptitude(d.population))
		// Print besta aptitude individual.
		fmt.Println("Best: ", d.population[0].aptitude)
		// Clone best n into the worst n of the population.
		for i := 0; i < d.info.population; i++ {
			d.population[i] = d.population[d.info.population+i]
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
		d.population[i].aptitude = (*d.eval).Eval(d.population[i].fenotype)
	}
	sort.Sort(ByAptitude(d.population))
	return d.population[0]
}
