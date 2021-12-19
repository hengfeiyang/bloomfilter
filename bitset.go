package bloomfilter

import (
	bitset "github.com/tmthrgd/go-bitset"
)

// bitSet type cotains library bitset and hasher function
type bitSet struct {
	b bitset.Bitset
}

// NewBitSet constructor for bitSet with an array of m bits
func NewBitSet(cfg Config) *bitSet {
	m := M(cfg.N, cfg.P)
	return &bitSet{bitset.New(uint(m))}
}

// Add element to bitset
func (bs *bitSet) Add(x uint64) {
	bs.b.Set(uint(x))
}

// Check element in bitset
func (bs *bitSet) Check(x uint64) bool {
	return bs.b.IsSet(uint(x))
}

// Count returns items num
func (bs *bitSet) Count() uint64 {
	return uint64(bs.b.Count())
}

// Optimize compress the runs of consecutive values found in the bitset
func (r *bitSet) Optimize() {
}
