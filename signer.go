package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

func main() {
}

// SingleHash calculate value crc32(data)+"~"+crc32(md5(data)), data is what came to the input.
func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for val := range in {
		dataStr := fmt.Sprintf("%v", val)
		md5 := DataSignerMd5(dataStr)
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			half1 := make(chan string)
			half2 := make(chan string)
			go func(dataStr string) {
				half1 <- DataSignerCrc32(dataStr)
			}(dataStr)
			go func(dataStr string) {
				half2 <- DataSignerCrc32(md5)
			}(dataStr)
			result := <-half1 + "~" + <-half2
			out <- result
		}(wg)
	}
	wg.Wait()
}

// MultiHash calculate value crc32(th+data)), where th=0..5 (i.e 6 hashes for every input value).
// After concatenate hashes in the order of calculation (0..5)
func MultiHash(in, out chan interface{}) {
	wgroup := &sync.WaitGroup{}
	for val := range in {
		dataStr := fmt.Sprintf("%v", val)
		wgroup.Add(1)
		go func(wgroup *sync.WaitGroup) {
			defer wgroup.Done()
			var hashTable [6]string
			wg := &sync.WaitGroup{}
			mu := &sync.Mutex{}
			for i := 0; i < 6; i++ {
				wg.Add(1)
				go func(i int, dataStr string, wg *sync.WaitGroup) {
					defer wg.Done()
					hash := DataSignerCrc32(fmt.Sprintf("%v", i) + dataStr)
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
		}(wgroup)
	}
	wgroup.Wait()
}

// CombineResults sort all results, concatenate them by "_".
func CombineResults(in, out chan interface{}) {
	hashes := make([]string, 0)
	for val := range in {
		hashes = append(hashes, fmt.Sprintf("%v", val))
	}
	sort.Strings(hashes)
	result := strings.Join(hashes, "_")
	out <- result
}

// ExecutePipeline provides pipelining workers
func ExecutePipeline(jobs ...job) {
	in := make(chan interface{})
	out := make(chan interface{})
	wg := sync.WaitGroup{}
	for _, job := range jobs {
		wg.Add(1)
		go func(in, out chan interface{}) {
			defer wg.Done()
			job(in, out)
			close(out)
		}(in, out)
		time.Sleep(1 * time.Millisecond)
		in = out
		out = make(chan interface{})
	}
	wg.Wait()
}
