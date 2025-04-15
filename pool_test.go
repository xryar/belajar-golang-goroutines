package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Arya")
	pool.Put("Rizki")
	pool.Put("Andaru")

	for i := 0; i < 10; i++ {
		data := pool.Get()
		fmt.Println(data)
		pool.Put(data)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}
