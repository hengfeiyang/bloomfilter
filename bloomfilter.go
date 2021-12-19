// Package bloomfilter contains a bloomfilter implement with roaring bitmap.
package bloomfilter

// bloomfilter basic type
type bloomfilter struct {
	bs  Bitmaper
	m   uint64
	k   uint64
	h   []Hash
	cfg Config
}

// New creates a new bloomfilter from a given config
func New(cfg Config, bm Bitmaper) *bloomfilter {
	m := M(cfg.N, cfg.P)
	k := K(m, cfg.N)
	return &bloomfilter{
		m:   m,
		k:   k,
		h:   HashFactoryNames[cfg.HashName](k),
		bs:  bm,
		cfg: cfg,
	}
}

// Add an element to bloomfilter
func (b *bloomfilter) Add(elem []byte) {
	for _, h := range b.h {
		for _, x := range h(elem) {
			b.bs.Add(x % b.m)
		}
	}
}

// Check if an element is in the bloomfilter
func (b *bloomfilter) Check(elem []byte) bool {
	for _, h := range b.h {
		for _, x := range h(elem) {
			if !b.bs.Check(x % b.m) {
				return false
			}
		}
	}
	return true
}

// Capacity returns the fill degree of the bloomfilter
func (b *bloomfilter) Capacity() float64 {
	return float64(b.bs.Count()) / float64(b.m)
}
