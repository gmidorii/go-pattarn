package main

import (
	"fmt"
	"sync"
)

func Semaphore(num int, s ...string) {
	sem := make(chan int, num)

	var wg sync.WaitGroup
	for _, v := range s {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			sem <- 1
			fmt.Println(v)
			<-sem
		}(v)
	}
	wg.Wait()
}
