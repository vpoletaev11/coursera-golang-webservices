package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

func main() {
	in := make(chan interface{})
	out := make(chan interface{})

	go SingleHash(in, out)
	in <- 1
	fmt.Println(<-out)
}

// SingleHash test
func SingleHash(in, out chan interface{}) {
	for val := range in {
		dataStr := fmt.Sprintf("%s", val)
		half1 := make(chan string)
		half2 := make(chan string)
		go func(dataStr string) {
			half1 <- DataSignerCrc32(dataStr)
		}(dataStr)
		go func(dataStr string) {
			half2 <- DataSignerCrc32(DataSignerMd5(dataStr))
		}(dataStr)
		result := <-half1 + "~" + <-half2
		out <- result
	}
	close(out)
}

// MultiHash test
func MultiHash(in, out chan interface{}) {
	dataStr := fmt.Sprintf("%v", <-in)
	var hashTable [6]string
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(i int, dataStr string, wg *sync.WaitGroup) {
			defer wg.Done()
			hash := DataSignerCrc32(string(i) + dataStr)
			mu.Lock()
			hashTable[i] = hash
			mu.Unlock()
		}(i, dataStr, wg)
	}
	result := ""
	wg.Wait()
	for _, val := range hashTable {
		result += val
	}
	out <- result
}

// CombineResults test
func CombineResults(in, out chan interface{}) {
	hashes := make([]string, 0)
	for val := range in {
		hashes = append(hashes, fmt.Sprintf("%s", val))
	}
	sort.Strings(hashes)
	result := strings.Join(hashes, "_")
	out <- result
}

func ExecutePipeline(...job) {

}
