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
	in <- 0
	if <-out != "4108050209~502633748" {
		t.Error("Func work incorrect")
	}
}

// TestSingleHash2 check efficiency of func SingleHash with multiple values on input
func TestSingleHash2(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})
	inputData := []int{0, 1, 2}

	go func() {
		for _, fibNum := range inputData {
			in <- fibNum
		}
		close(in)
	}()
	go func() {
		SingleHash(in, out)
		close(out)
	}()
	received := ""
	expected := "4108050209~5026337482212294583~709660146450215437~1933333237"
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
	in <- "4108050209~502633748"
	if <-out != "29568666068035183841425683795340791879727309630931025356555" {
		t.Error("Func work incorrect")
	}
}

// TestMultiHash2 check efficiency of func MultiHash with multiple values on input
func TestMultiHash2(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})
	inputValues := [3]string{"4108050209~502633748", "2212294583~709660146", "1450215437~1933333237"}

	go func() {
		for _, val := range inputValues {
			in <- val
		}
		close(in)
	}()
	go func() {
		MultiHash(in, out)
		close(out)
	}()
	outVals := make([]string, 0)
	for outVal := range out {
		outVals = append(outVals, fmt.Sprintf("%s", outVal))
	}
	sort.Strings(outVals)
	received := strings.Join(outVals, "")
	expected := "295686660680351838414256837953407918797273096309310253565554124676206380750763036832117583433119022284321470231929404624958044192186797981418233587017209679042592862002427381542"
	if received != expected {
		t.Error("Func work incorrect")
	}
}

// TestCombineResults check efficiency of func CombineResults
func TestCombineResults(t *testing.T) {
	in := make(chan interface{})
	out := make(chan interface{})
	inputData := []string{
		"4958044192186797981418233587017209679042592862002427381542",
		"412467620638075076303683211758343311902228432147023192940462",
		"29568666068035183841425683795340791879727309630931025356555"}

	go func() {
		for _, val := range inputData {
			in <- val
		}
		close(in)
	}()
	go CombineResults(in, out)
	expected := "29568666068035183841425683795340791879727309630931025356555_412467620638075076303683211758343311902228432147023192940462_4958044192186797981418233587017209679042592862002427381542"
	if fmt.Sprintf("%s", <-out) != expected {
		t.Error("Func work incorrect")
	}
}
