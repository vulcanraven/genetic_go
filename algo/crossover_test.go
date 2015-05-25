package genetic

import (
	"reflect"
	"testing"
)

func TestByteSwap(t *testing.T) {
	a, b := maskByteSwap(7, 8, 7)
	if a != 0 || b != 15 {
		t.Fatalf("Got: a = %d, b = %d, Expected: a = %d, b = %d", a, b, 0, 15)
	}
	a, b = maskByteSwap(34, 78, 255)
	if a != 78 || b != 34 {
		t.Fatalf("Got: a = %d, b = %d, Expected: a = %d, b = %d", a, b, 78, 34)
	}
}

func TestAnnularCrossover(t *testing.T) {
	a := Individual{genes: []byte{0xAB, 0xCD}}
	b := Individual{genes: []byte{0x12, 0x34}}
	info := GeneticInfo{bytes: 2}
	AnnularCrossover(a, b, info, 4)
	if reflect.DeepEqual(a.genes, []byte{0xA2, 0x3D}) || reflect.DeepEqual(b.genes, []byte{0x1B, 0xC4}) {
		t.Fatalf("Got: a = %v, b = %v Expected: a = %v, b = %v", a.genes, b.genes, []byte{0xA2, 0x3D}, []byte{0x1B, 0xC4})
	}

	a = Individual{genes: []byte{0xAB, 0x11, 0xCD}}
	b = Individual{genes: []byte{0x12, 0x22, 0x34}}
	info = GeneticInfo{bytes: 2}
	AnnularCrossover(a, b, info, 4)
	if reflect.DeepEqual(a.genes, []byte{0xA2, 0x22, 0x3D}) || reflect.DeepEqual(b.genes, []byte{0x1B, 0x11, 0xC4}) {
		t.Fatalf("Got: a = %v, b = %v Expected: a = %v, b = %v", a.genes, b.genes, []byte{0xA2, 0x22, 0x3D}, []byte{0x1B, 0x11, 0xC4})
	}
}
