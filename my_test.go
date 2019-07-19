package main

import (
	"fmt"
	"testing"
)

// SingleHashTest1 check efficiency of func SingleHash with one value on input
func TestSingleHash1(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})

	go SingleHash(in, out)
	in <- 1
	if <-out != "1140956898~3176729503" {
		t.Error("Func work incorrect")
	}
}

//SingleHashTest2 check efficiency of func SingleHash with multiple values on input
func TestSingleHash2(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})
	inputData := []int{0, 1, 2}

	go SingleHash(in, out)
	go func() {
		for _, fibNum := range inputData {
			in <- fibNum
		}
	}()
	for outVal := range out {
		fmt.Println(outVal)
	}
}


func TestMultiHash1(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})

	go MultiHash(in, out)
	in <- "1140956898~3176729503"
	fmt.Println(<-out)
}
