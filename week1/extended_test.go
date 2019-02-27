package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ├───project
func TestTabGen01(t *testing.T) {
	//arrange
	input := dir{IsLast: false, PrevDirsLast: []bool{false}}

	//act
	actual := tabGen(input)

	//assert
	assert.Equal(t, "├───", actual)
}

// │	├───file.txt (19b)
func TestTabGen02(t *testing.T) {
	//arrange
	input := dir{IsLast: false, PrevDirsLast: []bool{false, false}}

	//act
	actual := tabGen(input)

	//assert
	assert.Equal(t, "│	├───", actual)
}

// │	└───gopher.png (70372b)
func TestTabGen03(t *testing.T) {
	//arrange
	input := dir{IsLast: true, PrevDirsLast: []bool{false, false}}

	//act
	actual := tabGen(input)

	//assert
	assert.Equal(t, "│	└───", actual)
}
