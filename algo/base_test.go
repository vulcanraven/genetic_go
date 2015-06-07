package genetic

import (
	"reflect"
	"testing"
)

func TestFenotype(t *testing.T) {
	inf := GeneticInfo{bytes: 2, genes: 1, genesize: 2, decimalbits: 8, intbits: 8}
	ind, _ := CreateIndividual(&inf)
	ind.genes = []byte{0x03, 0x80}
	ind.UpdateFenotype()
	if ind.fenotype[0] != 1.75 {
		t.Fatalf("Got: %f, Expected: %f", ind.fenotype[0], 1.75)
	}

	inf = GeneticInfo{bytes: 2, genes: 1, genesize: 2, decimalbits: 8, intbits: 7, signbit: true}
	ind, _ = CreateIndividual(&inf)
	ind.genes = []byte{0x03, 0x81}
	ind.UpdateFenotype()
	if ind.fenotype[0] != -1.75 {
		t.Fatalf("Got: %f, Expected: %f", ind.fenotype[0], -1.75)
	}

	inf = GeneticInfo{bytes: 4, genes: 2, genesize: 2, decimalbits: 8, intbits: 7, signbit: true}
	ind, _ = CreateIndividual(&inf)
	ind.genes = []byte{0x03, 0x81, 0x03, 0x80}
	ind.UpdateFenotype()
	if !reflect.DeepEqual(ind.fenotype, []float64{-1.75, 1.75}) {
		t.Fatalf("Got: %v, Expected: %v", ind.fenotype, []float64{-1.75, 1.75})
	}

	// Force byte config error
	inf = GeneticInfo{bytes: 2, genes: 1, genesize: 2, decimalbits: 8, intbits: 8, signbit: true}
	_, err := CreateIndividual(&inf)
	if err == nil {
		t.Fatalf("Error should not be nil.")
	}

	// Force gene error
	inf = GeneticInfo{bytes: 2, genes: 2, genesize: 2, decimalbits: 8, intbits: 7, signbit: true}
	_, err = CreateIndividual(&inf)
	if err == nil {
		t.Fatalf("Error was not generated with faulty genetic info config.")
	}
}

func TestMutateBit(t *testing.T) {
	inf := GeneticInfo{bytes: 2, genes: 1, genesize: 2, decimalbits: 8, intbits: 8}
	ind, _ := CreateIndividual(&inf)
	ind.genes = []byte{8, 8}
	ind.RandomMutateBit()
	if reflect.DeepEqual(ind.genes, []byte{8, 8}) {
		t.Fatalf("Got: %v, Expected something other than: %v", ind.genes, []byte{8, 8})
	}
}
