package main

import (
	"fmt"
	"sync"
	"time"
)

func syncWait(f, s string) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func(arg string) {
		defer wg.Done()
		time.Sleep(time.Millisecond)
		fmt.Println(arg)
	}(s)

	fmt.Println(f)
	wg.Wait()
}
