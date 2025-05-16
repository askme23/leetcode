package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidString(t *testing.T) {
	testCases := []struct {
		str    string
		expect bool
	}{
		{str: "(*))", expect: true},
		{str: "(*", expect: true},
		{str: "**************************************************))))))))))))))))))))))))))))))))))))))))))))))))))", expect: true},
		{str: "(*(()*))", expect: true},
		{str: "********", expect: true},
	}

	for _, tc := range testCases {
		t.Run(tc.str, func(t *testing.T) {
			res := checkValidString(tc.str)
			assert.Equal(t, tc.expect, res)
		})
	}
}

func checkValidString(s string) bool {
	n := len(s)
	memo := make([][]int, n)
	for i := range memo {
		row := make([]int, n)
		for j := 0; j < n; j++ {
			row[j] = -1
		}
		memo[i] = row
	}

	return isValidString(0, 0, s, memo)
}

func isValidString(index int, openCount int, s string, memo [][]int) bool {
	if index == len(s) {
		return openCount == 0
	}

	if memo[index][openCount] != -1 {
		// fmt.Println("memo", s[:index+1], "|", index, "|", openCount, "|", memo[index][openCount])
		return memo[index][openCount] == 1
	}

	var isValid bool

	if s[index] == '*' {
		isValid = isValid || isValidString(index+1, openCount+1, s, memo)

		if openCount > 0 {
			isValid = isValid || isValidString(index+1, openCount-1, s, memo)
		}

		isValid = isValid || isValidString(index+1, openCount, s, memo)
	} else {
		if s[index] == '(' {
			isValid = isValid || isValidString(index+1, openCount+1, s, memo)
		} else if openCount > 0 {
			isValid = isValid || isValidString(index+1, openCount-1, s, memo)
		}
	}

	// fmt.Println(s[:index+1], "|", index, "|", openCount, "|", isValid)
	memo[index][openCount] = 0
	if isValid {
		memo[index][openCount] = 1
	}

	return isValid
}
