package genetic

import (
	"math"
	"testing"
)

type HansenFunc struct{}

func (h *HansenFunc) Eval(v []float64) float64 {
	p1, p2, result := float64(0), float64(0), float64(0)
	for i := 0; i <= 4; i++ {
		p1 += (float64(i) + 1.0) * math.Cos(float64(i)*v[0]+float64(i)+1.0)
		p2 += (float64(i) + 1) * math.Cos((float64(i)+2)*v[1]+float64(i)+1)
	}
	result = p1 * p2
	return result
}

func TestEga(t *testing.T) {
	info := &GeneticInfo{
		genesize:    2,
		genes:       2,
		intbits:     7,
		decimalbits: 8,
		signbit:     true,
		population:  100,
		mutation:    0.01,
		crossover:   0.9,
	}
	ga := Ega{}
	err := ga.Setup(info, &HansenFunc{})
	if err != nil {
		t.Fatalf(err.Error())
	}

	best := ga.Run(50)

	if best.aptitude != 0 {
		t.Fatalf("Got: %f, Expected: %f", best.aptitude, -176.54)
	}

}
