package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

// сюда писать код
func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for val := range in {
		dataStr := fmt.Sprintf("%s", val)
		half1 := make(chan string)
		half2 := make(chan string)
		wg.Add(1)
		go func(dataStr string, ch chan string, wg *sync.WaitGroup) {
			defer wg.Done()
			half1 <- DataSignerCrc32(dataStr)
		}(dataStr, half1, wg)
		wg.Add(1)
		go func(dataStr string, ch chan string, wg *sync.WaitGroup) {
			defer wg.Done()
			half2 <- DataSignerCrc32(DataSignerMd5(dataStr))
		}(dataStr, half2, wg)
		wg.Wait()
		result := <-half1 + "~" + <-half2
		out <- result
	}
	close(out)
}

func main() {
	in := make(chan interface{})
	out := make(chan interface{}, 10)

	array := []string{"data", "biba", "lol"}
	go SingleHash(in, out)
	in <- array
	//////////
	//for val := range out {
	//	fmt.Println(val)
	// }
	// go MultiHash(in, out)
	// in <- "2918445923~1798600672"
	// fmt.Println(<-out)
	// go CombineResults(in, out)
	// in <- array
	// time.Sleep(3 * time.Second)
	for i := range out {
		fmt.Println(i)
	}

}

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

func CombineResults(in, out chan interface{}) {
	hashes := make([]string, 0)
	for val := range in {
		hashes = append(hashes, fmt.Sprintf("%s", val))
	}
	sort.Strings(hashes)
	result := strings.Join(hashes, "_")
	out <- result
}
