package main

import (
	"gateway/tools/rate"
	"time"
)

func main() {
	bucket := rate.NewTokenBucket(6, 2)

	for a := 1; a < 10000;a++ {
		go func() {
			err := bucket.Take()
			if err != nil {
				println(err.Error())
			} else {
				println("s")
			}
		}()
		if a % 10 == 0 {
			time.Sleep(3 * time.Second)
		}
	}
}
