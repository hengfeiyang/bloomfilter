# bloomfilter

a bloomfilter implement with roaring bitmap.

Usage:

```
package main

import (
    "fmt"

    "github.com/safeie/bloomfilter"
)

func main() {
    cfg := bloomfilter.Config{
		N:        10000000000,    // capacity
		P:        0.00001,        // false probability
		HashName: HASHER_OPTIMAL, // hash functions
	}
	bf := bloomfilter.New(cfg)
    bf.Add([]byte("www.google.com"))
	bf.Add([]byte("twitter.com"))
	bf.Add([]byte("github.com"))

    fmt.Println(bf.Check([]byte("twitter.com")))
    fmt.Println(bf.Check([]byte("facebook.com")))
}
```
