// Package baseBloomfilter implements a bloomfilter based on an m-bit bit array, k hashfilters and configuration.
//
// It creates a bloomfilter based on bitset and an injected hasher, along with the
// following operations: add an element to the bloomfilter, check the existence of an element
// in the bloomfilter, the union of two bloomfilters, along with the serialization and
// deserialization of a bloomfilter: http://llimllib.github.io/bloomfilter-tutorial/
package baseBloomfilter

import (
	"log"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/devopsfaith/bloomfilter"
)

// Bloomfilter basic type
type Bloomfilter struct {
	bs  *roaring64.Bitmap
	m   uint
	k   uint
	h   []bloomfilter.Hash
	cfg bloomfilter.Config
}

// New creates a new bloomfilter from a given config
func New(cfg bloomfilter.Config) *Bloomfilter {
	m := bloomfilter.M(cfg.N, cfg.P)
	k := bloomfilter.K(m, cfg.N)
	log.Println(m, k)
	return &Bloomfilter{
		m:   m,
		k:   k,
		h:   bloomfilter.HashFactoryNames[cfg.HashName](k),
		bs:  roaring64.NewBitmap(),
		cfg: cfg,
	}
}

// Add an element to bloomfilter
func (b Bloomfilter) Add(elem []byte) {
	for _, h := range b.h {
		for _, x := range h(elem) {
			b.bs.Add(uint64(x % b.m))
		}
	}
}

// Check if an element is in the bloomfilter
func (b Bloomfilter) Check(elem []byte) bool {
	for _, h := range b.h {
		for _, x := range h(elem) {
			if !b.bs.Contains(uint64(x % b.m)) {
				return false
			}
		}
	}
	return true
}

// Capacity returns the fill degree of the bloomfilter
func (b *Bloomfilter) GetSizeInBytes() uint64 {
	return b.bs.GetSizeInBytes()
}
