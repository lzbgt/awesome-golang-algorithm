package Solution

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
)

//	solution func Info
type SolutionFuncType func([]int, int) []int

var SolutionFuncList = []SolutionFuncType{
	twoSum,
}

//	test info struct
type Case struct {
	name   string
	inputs []int
	target int
	expect []int
}

// 	test case
var cases = []Case{
	{
		name:   "TestCase 1",
		inputs: []int{2, 7, 11, 15},
		target: 9,
		expect: []int{2, 7},
	},

	{
		name:   "TestCase 2",
		inputs: []int{10, 26, 30, 31, 47, 60},
		target: 40,
		expect: []int{10, 30},
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.inputs, c.target)
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}