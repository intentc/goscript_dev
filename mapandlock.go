package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var lock sync.Mutex
var rwlock sync.RWMutex

func testmap() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[1] = 10
	a[2] = 10
	a[3] = 10
	a[4] = 10
	a[5] = 10
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[1] = rand.Intn(100)
			lock.Unlock()

		}(a)

		lock.Lock()
		fmt.Println(a)
		lock.Unlock()

		time.Sleep(time.Second)
	}
}

func main() {
	testmap()
	fmt.Printf("pass")

}
