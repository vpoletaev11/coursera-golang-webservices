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

	go func() {
		for _, fibNum := range inputData {
			in <- fibNum
		}
	}()
	go SingleHash(in, out)
	received := ""
	expected := "1562029987~36665590381140956898~31767295031865207073~94904396"
	for outVal := range out {
		strOutVal := fmt.Sprintf("%s", outVal)
		received += strOutVal
	}
	if received != expected {
		t.Error("Func work incorrect")
	}
}

func TestMultiHash1(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})

	go MultiHash(in, out)
	in <- "1140956898~3176729503"
	if <-out != "300712648027528722082641220176231982452840200132644172606480" {
		t.Error("Func work incorrect")
	}
}

func TestMultiHash2(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})
	inputValues := [3]string{"1562029987~3666559038", "1140956898~3176729503", "1865207073~94904396"}

	go MultiHash(in, out)
	go func() {
		for _, val := range inputValues {
			in <- val
		}
		close(in)
	}()
	for outVal := range out {
		fmt.Println(outVal)
	}
}
