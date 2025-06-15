package main

import (
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func randomtask() {
	defer wg.Done()
	randomSeconds := rand.Intn(5) + 1
	time.Sleep(time.Second * time.Duration(randomSeconds))
}

func scheduler(fns []func()) {

	for _, fn := range fns {
		wg.Add(1)
		fn()
	}
	wg.Wait()
}
