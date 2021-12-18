package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/devopsfaith/bloomfilter"
	bf "github.com/devopsfaith/bloomfilter/bloomfilter"
)

func main() {
	logger := log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds)
	cfg := bloomfilter.Config{
		N:        10000000000,
		P:        0.00001,
		HashName: bloomfilter.HASHER_OPTIMAL,
	}
	client := bf.New(cfg)

	hash := bloomfilter.OptimalHashFactory(17)
	logger.Println(hash[0]([]byte("www.baidu.com")))
	logger.Println(hash[0]([]byte("www.baidu.com")))

	urls := []string{
		"http://www.baidu.com",
		"http://www.google.com",
	}
	for i := 0; i < 1000000; i++ {
		if i%100000 == 0 {
			logger.Println(i, client.GetSizeInBytes())
		}
		for _, url := range urls {
			client.Add([]byte(fmt.Sprintf("%d:%s", i, url)))
		}
	}

	logger.Println(client.GetSizeInBytes())
	logger.Println(client.Check([]byte("100:http://www.baidu.com")))
	logger.Println(client.Check([]byte("http://www.fei.com")))
	logger.Println(client.Check([]byte("100000:http://www.google.com")))

	time.Sleep(time.Minute)
}
