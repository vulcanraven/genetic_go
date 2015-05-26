package genetic

import (
	"fmt"
	"math"
	"math/rand"
)

type GeneticInfo struct {
	bytes       int     // Bytes each individual has.
	genesize    int     // Bytes each gene has.
	genes       int     // Genes each individual has.
	intbits     int     // Bits used for the integer part of each gene.
	decimalbits int     // Bits used for the decimal part of each gene.
	signbit     bool    // Genes have a sign bit.
	population  int     // Number of individuals in the population.
	mutation    float32 // Mutation rate.
	crossover   float32 // Crossover rate.
}

type Individual struct {
	genes    []byte    // genetic information
	fenotype []float64 // gene values
	info     *GeneticInfo
	aptitude float64
}

// Return the size of the entire population in bytes
func (g GeneticInfo) PopulationBytes() int {
	return g.population * g.bytes
}

// ByAptitude implements sort.Interface for []*Individual based on
// the aptitude field.
type ByAptitude []*Individual

func (a ByAptitude) Len() int           { return len(a) }
func (a ByAptitude) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAptitude) Less(i, j int) bool { return a[i].aptitude < a[j].aptitude }

func CreateIndividual(inf *GeneticInfo) (*Individual, error) {
	if (inf.signbit && (inf.decimalbits+inf.intbits+1)%8 != 0) || (!inf.signbit && (inf.decimalbits+inf.intbits)%8 != 0) {
		err := fmt.Errorf("Gene size has to be dvisible by 8. Current values decimalbits: %d intbits: %d signbit: %b", inf.decimalbits, inf.intbits, inf.signbit)
		return nil, err
	}
	if inf.bytes != inf.genes*inf.genesize {
		err := fmt.Errorf("The genes and genesize don't match the amount of bytes available. genes: %d genesize: %d bytes: %d", inf.genes, inf.genesize, inf.bytes)
		return nil, err
	}
	ind := Individual{info: inf, genes: make([]byte, inf.bytes), fenotype: make([]float64, inf.genes)}
	for i := range ind.genes {
		ind.genes[i] = byte(rand.Intn(256))
	}
	ind.UpdateFenotype()
	return &ind, nil
}

// Update the fenotype values from the genotype.
func (i *Individual) UpdateFenotype() {
	for g := 0; g < i.info.genes; g++ {
		power := -i.info.decimalbits
		acum := float64(0.0)
		for b := 0; b < i.info.genesize; b++ {
			// iterate over each byte
			for t := 0; t < 8; t++ {
				mask := byte(1) << uint(7-t)
				// check if it has that bit turned on and accumulate it's value
				if mask&i.genes[g*i.info.genesize+b] > 0 {
					if i.info.signbit && b == i.info.genesize-1 && t == 7 {
						acum *= -1
					} else {
						acum += math.Pow(2, float64(power))
					}
				}
				power++
			}
		}
		i.fenotype[g] = acum
	}
}
