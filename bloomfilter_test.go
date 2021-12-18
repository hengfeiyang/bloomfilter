package bloomfilter

import (
	"testing"
)

func TestBloomfilter(t *testing.T) {
	cfg := Config{
		N:        1000000,        // capacity
		P:        0.00001,        // false probability
		HashName: HASHER_OPTIMAL, // hash functions
	}
	bf := New(cfg)
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
