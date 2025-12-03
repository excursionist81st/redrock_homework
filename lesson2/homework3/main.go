package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.RWMutex
	count int
}

func Increment(c *Counter, incre *sync.WaitGroup) {
	defer incre.Done()
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}
func Incre10(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	var incre sync.WaitGroup
	for i := 0; i < 10; i++ {
		incre.Add(1)
		go Increment(c, &incre)
	}
	incre.Wait()
}
func Value(c *Counter) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}
func main() {
	var c Counter
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go Incre10(&c, &wg)
	}
	wg.Wait()
	end := Value(&c)
	fmt.Printf("最终计数：%d", end)
}
