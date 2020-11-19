package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

var lock sync.Mutex
var rwlock sync.RWMutex

func testmap() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[11] = 10
	a[21] = 10
	a[31] = 10
	a[41] = 10
	a[51] = 10

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, v := range keys {
		fmt.Println(v, a[v])
	}

	//	b := make(map[int]int)

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwlock.Lock()
			b[11] = rand.Intn(100)
			rwlock.Unlock()

		}(a)

		rwlock.RLock()
		fmt.Println(a)
		rwlock.RUnlock()

		time.Sleep(time.Second)
	}
}

func main() {
	testmap()
	fmt.Println("pass")

}
