// Package bloomfilter contains a bloomfilter implement with roaring bitmap.
package bloomfilter

import (
	"math"

	"github.com/RoaringBitmap/roaring/roaring64"
)

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

// bloomfilter basic type
type bloomfilter struct {
	bs  *roaring64.Bitmap
	m   uint64
	k   uint64
	h   []Hash
	cfg Config
}

// New creates a new bloomfilter from a given config
func New(cfg Config) *bloomfilter {
	m := M(cfg.N, cfg.P)
	k := K(m, cfg.N)
	return &bloomfilter{
		m:   m,
		k:   k,
		h:   HashFactoryNames[cfg.HashName](k),
		bs:  roaring64.NewBitmap(),
		cfg: cfg,
	}
}

// Add an element to bloomfilter
func (b bloomfilter) Add(elem []byte) {
	for _, h := range b.h {
		for _, x := range h(elem) {
			b.bs.Add(x % b.m)
		}
	}
}

// Check if an element is in the bloomfilter
func (b bloomfilter) Check(elem []byte) bool {
	for _, h := range b.h {
		for _, x := range h(elem) {
			if !b.bs.Contains(uint64(x % b.m)) {
				return false
			}
		}
	}
	return true
}

// GetSizeInBytes estimates the memory usage of the Bitmap
func (b *bloomfilter) GetSizeInBytes() uint64 {
	return b.bs.GetSizeInBytes()
}
