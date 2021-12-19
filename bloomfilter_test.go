package bloomfilter

import (
	"fmt"
	"testing"
)

var cfg = Config{
	N:        100000000,      // capacity
	P:        0.00001,        // false probability
	HashName: HASHER_OPTIMAL, // hash functions
}

func TestBloomfilter(t *testing.T) {
	bf := New(cfg, NewRoaring(cfg))
	bf.Add([]byte("www.google.com"))
	bf.Add([]byte("twitter.com"))
	bf.Add([]byte("github.com"))

	if ok := bf.Check([]byte("twitter.com")); !ok {
		t.Errorf("[twitter.com] should be exists")
	}
	if ok := bf.Check([]byte("facebook.com")); ok {
		t.Errorf("[facebook.com] should be not exists")
	}
}

func BenchmarkBloomfilterRoaring_add(b *testing.B) {
	bf := New(cfg, NewRoaring(cfg))
	for i := 0; i < b.N; i++ {
		bf.Add([]byte(fmt.Sprintf("https://www.google.com/%d", i)))
	}
}

func BenchmarkBloomfilterBitSet_add(b *testing.B) {
	bf := New(cfg, NewBitSet(cfg))
	for i := 0; i < b.N; i++ {
		bf.Add([]byte(fmt.Sprintf("https://www.google.com/%d", i)))
	}
}

func BenchmarkBloomfilterRoaring_check(b *testing.B) {
	bf := New(cfg, NewRoaring(cfg))
	for i := 0; i < b.N; i++ {
		bf.Check([]byte(fmt.Sprintf("https://www.google.com/%d", i)))
	}
}

func BenchmarkBloomfilterBitSet_check(b *testing.B) {
	bf := New(cfg, NewBitSet(cfg))
	for i := 0; i < b.N; i++ {
		bf.Check([]byte(fmt.Sprintf("https://www.google.com/%d", i)))
	}
}
