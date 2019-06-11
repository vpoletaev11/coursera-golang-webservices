package main

import (
	"fmt"
	"sync"
)

// сюда писать код
func SingleHash(in, out chan interface{}) {
	dataStr := fmt.Sprintf("%s", <-in)
	half1 := make(chan string)
	half2 := make(chan string)
	go func(dataStr string, ch chan string) {
		half1 <- DataSignerCrc32(dataStr)
	}(dataStr, half1)
	go func(dataStr string, ch chan string) {
		half2 <- DataSignerCrc32(DataSignerMd5(dataStr))
	}(dataStr, half2)
	result := <-half1 + "~" + <-half2
	out <- result
}

func main() {
	in := make(chan interface{})
	out := make(chan interface{})

	// go SingleHash(in, out)
	// in <- "data"
	// fmt.Println(<-out)

	go MultiHash(in, out)
	in <- "2918445923~1798600672"
	fmt.Println(<-out)
}

func MultiHash(in, out chan interface{}) {
	dataStr := fmt.Sprintf("%v", <-in)
	var hashTable [6]string
	mu := &sync.Mutex{}
	for i := 0; i < 6; i++ {
		go func(i int, dataStr string) {
			mu.Lock()
			hashTable[i] = DataSignerCrc32(string(i) + dataStr)
			mu.Unlock()
		}(i, dataStr)
	}
	result := ""
	for _, val := range hashTable {
		mu.Lock()
		result += val
		mu.Unlock()
	}
	out <- result
}
