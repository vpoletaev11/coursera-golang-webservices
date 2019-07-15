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
		err := fmt.Errorf("Func work incorrect")
		fmt.Println(err.Error())
	}
	fmt.Println("TestSingleHash1 DONE")
}

//SingleHashTest2 check efficiency of func SingleHash with multiple values on input
func TestSingleHash2(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})
	inputData := []int{0, 1, 2}

	SingleHash(in, out)
	for _, fibNum := range inputData {
		in <- fibNum
	}

	for outVal := range out {
		fmt.Println(outVal)
	}
}
