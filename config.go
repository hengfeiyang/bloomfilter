package bloomfilter

import "math"

// Config for bloomfilter defining the parameters:
// N - number of elements to be stored in the filter
// P - desired false positive probability
// HashName - the name of the particular hashfunction
type Config struct {
	N        uint64  // capacity
	P        float64 // false probability
	HashName string  // hash functions
}

// M function computes the length of the bit array of the bloomfilter as function of n and p
func M(n uint64, p float64) uint64 {
	return uint64(math.Ceil(-(float64(n) * math.Log(p)) / math.Log(math.Pow(2.0, math.Log(2.0)))))
}

// K function computes the number of hashfunctions of the bloomfilter as function of n and p
func K(m, n uint64) uint64 {
	return uint64(math.Ceil(math.Log(2.0) * float64(m) / float64(n)))
}
