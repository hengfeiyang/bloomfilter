package bloomfilter

import (
	"github.com/RoaringBitmap/roaring/roaring64"
)

// roaring type cotains library roaring bitmap
type roaring struct {
	b *roaring64.Bitmap
}

// NewRoaring constructor for Roaring
func NewRoaring(cfg Config) *roaring {
	return &roaring{roaring64.NewBitmap()}
}

// Add element to roaring
func (r *roaring) Add(x uint64) {
	r.b.Add(x)
}

// Check element in roaring
func (r *roaring) Check(x uint64) bool {
	return r.b.Contains(x)
}

// Count returns items num
func (r *roaring) Count() uint64 {
	return r.b.GetCardinality()
}

// Optimize compress the runs of consecutive values found in the bitmap
func (r *roaring) Optimize() {
	r.b.RunOptimize()
}
