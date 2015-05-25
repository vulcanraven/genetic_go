package genetic

import (
	"math/rand"
)

// Generic crossover function
type CrossFunc func(Individual, Individual, GeneticInfo)

// Swap the bytes of a and b where the mask has ones.
func maskByteSwap(a, b, mask byte) (byte, byte) {
	i_a := mask & a
	i_b := mask & b
	n_mask := ^mask
	a = a & n_mask
	b = b & n_mask
	a = a | i_b
	b = b | i_a
	return a, b
}

// Randomized version of annular crossover.
func RandomAnnularCrossover(a, b Individual, info GeneticInfo) {
	AnnularCrossover(a, b, info, rand.Intn(info.bytes*4))
}

// Annular crossover.
func AnnularCrossover(a, b Individual, info GeneticInfo, start int) {
	bytestart := start / 8
	singles := start % 8
	fullsbyteswap := (info.bytes / 2)
	// swap first bits
	if singles > 0 {
		fullsbyteswap -= 1
		mask := byte(1) << uint(8-singles)
		mask = (mask - 1) & mask
		a.genes[bytestart], b.genes[bytestart] = maskByteSwap(a.genes[bytestart], b.genes[bytestart], mask)
		bytestart++
	}

	if fullsbyteswap > 0 {
		for i := bytestart; i < bytestart+fullsbyteswap; i++ {
			a.genes[i], b.genes[i] = b.genes[i], a.genes[i]
		}
	}

	// swap last bits
	if singles > 0 {
		mask := byte(1) << uint(8-singles)
		mask = ^((mask - 1) & mask)
		last := bytestart + fullsbyteswap
		a.genes[last], b.genes[last] = maskByteSwap(a.genes[last], b.genes[last], mask)
	}
}
