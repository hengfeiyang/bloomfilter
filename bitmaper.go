package bloomfilter

// Bitmaper bitmap interface
type Bitmaper interface {
	Add(x uint64)
	Check(x uint64) bool
	Count() uint64
	Optimize()
}
