package genetic

import (
	"math/rand"
)

type Fitness interface {
	Eval([]float32) float64
}

type Ega struct {
	info       GeneticInfo
	population []Individual
	eval       *Fitness
}

func (d *Ega) Setup(inf GeneticInfo, evalf *Fitness) { // , pointer_to_conditional_stop_func
	d.info = inf
	d.population = make([]Individual, inf.population*2) // double the size of the population.
	for i := range d.population {
		d.population[i].genes = make([]byte, inf.bytes)
		for j := 0; j < inf.bytes; j++ {
			d.population[i].genes[j] = byte(rand.Intn(256))
		}
	}
}

func (d *Ega) Run(generations int) (float64, Individual) {
	for i := 0; i < generations; i++ {

	}
	return 0, Individual{}
}
