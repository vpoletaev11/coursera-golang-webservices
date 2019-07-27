package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

// TestSingleHash1 check efficiency of func SingleHash with one value on input
func TestSingleHash1(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})

	go SingleHash(in, out)
	in <- 1
	if <-out != "1140956898~3176729503" {
		t.Error("Func work incorrect")
	}
}

//TestSingleHash2 check efficiency of func SingleHash with multiple values on input
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

// TestMultiHash1 check efficiency of func MultiHash with one value on input
func TestMultiHash1(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})

	go MultiHash(in, out)
	in <- "1140956898~3176729503"
	if <-out != "300712648027528722082641220176231982452840200132644172606480" {
		t.Error("Func work incorrect")
	}
}

//TestMultiHash2 check efficiency of func MultiHash with multiple values on input
func TestMultiHash2(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})
	inputValues := [3]string{"1562029987~3666559038", "1140956898~3176729503", "1865207073~94904396"}

	go func() {
		for _, val := range inputValues {
			in <- val
		}
		close(in)
	}()
	go MultiHash(in, out)
	outVals := make([]string, 0)
	for outVal := range out {
		outVals = append(outVals, fmt.Sprintf("%s", outVal))
	}
	sort.Strings(outVals)
	received := strings.Join(outVals, "")
	expected := "15310779131042636383244187637341090521153638267521892201734260766864023533796803040678944271931414433520554563504613984300712648027528722082641220176231982452840200132644172606480"
	if received != expected {
		t.Error("Func work incorrect")
	}
}
