package main

import(
	"testing"
)

type TestCase struct {
	dice	[]int
	board 	[]int
	expect	int
}


func Test(t *testing.T) {
	testCases := []TestCase{
		{dice:[]int{2,1,5,1,5,4}, board:[]int{0,0,3,0,0,0,0,-2,0,0,0}, expect:10},
	}

	for _, tc := range(testCases) {
		got := SnakesAndLadders(tc.board, tc.dice)
		want := tc.expect
		if (got!=want) {
			t.Errorf("\ndice: %v\nboard: %v\ngot: %v\nwant: %v", tc.dice, tc.board, got, want)
	
		}
	}

}