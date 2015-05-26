package genetic

import (
	"sort"
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
	// Evaluate first n.
	for i := 0; i < inf.population; i++ {
		d.population[i].aptitude = (*d.eval).Eval(d.population[i].fenotype)
	}
	return nil
}

func (d *Ega) Run(generations int) (float64, Individual) {
	for i := 0; i < generations; i++ {
		// Evaluate individuals.
		for i := d.info.population; i < d.info.population*2; i++ {
			d.population[i].aptitude = (*d.eval).Eval(d.population[i].fenotype)
		}
		// Sort by aptitude.
		sort.Sort(ByAptitude(d.population))
		// Clone best n into the worst n of the population.
		for i := 0; i < d.info.population; i++ {
			d.population[i] = d.population[d.info.population+i]
		}
		// Cross over with the eclectic operator.
		for i := 0; i < d.info.population/2; i++ {
			RandomAnnularCrossover(d.population[i], d.population[d.info.population-i-1], d.info)
		}
		// Randomly mutate bits.

		// Update fenotype.
	}
	return 0, Individual{}
}
