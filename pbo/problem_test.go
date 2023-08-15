package pbo

import (
	"fmt"
	"math"
	"testing"
)

// https://github.com/IOHprofiler/IOHexperimenter/blob/master/tests/cpp/problem/test_pbo_problem.cpp

func testEval(p Problem) float64 {
	x := []bool{true, true, false, true, false, false, false, true, true}
	return p.Eval(x)
}

// Tests that all problems increment their eval count properly
func TestNEvals(t *testing.T) {
	for i := 0; i <= 25; i++ {

		p := GetProblem(i, 9, 0)

		// Should start at 0
		if got := p.State().nEvals; got != 0 {
			t.Errorf("Problem %d's nEvals is %d, not 0 before evaling", i, got)
		}

		// Evaluate a solution
		testEval(p)

		// After it should be 1
		if got := p.State().nEvals; got != 1 {
			t.Errorf("Problem %d's nEvals is %d, not 1 after evaling", i, got)
		}

	}
}

// Tests that all problems implement seed instances correctly
func TestSeeds(t *testing.T) {

	for i := 0; i <= 25; i++ {

		p1 := GetProblem(i, 9, 0)
		p2 := GetProblem(i, 9, 0)

		// Should get same result
		if testEval(p1) != testEval(p2) {
			t.Errorf("Problem %d does not implement seed correctly.", i)
		}

	}
}

func TestProblems(t *testing.T) {

	tests := []struct {
		id    int
		score float64
	}{
		{0, 0.0},
		{1, 5.0},
		{2, 2.0},
		{3, 24.0},
		{6, 2.0},
		{7, 6.0},
		{8, 4.0},
		{9, 6.0},
		{10, 7.0},
		{13, 1.0},
		{14, 0.0},
		{15, 2.0},
		{16, 1.0},
		{17, 1.0},
		{18, 2.5312},
		{19, 5.0},
		{20, 6.0},
		{21, 9.0},
		{22, -4.0},
		{23, -16.0},
	}

	eps := 0.0001

	for _, test := range tests {
		testName := fmt.Sprintf("%d", test.id)
		t.Run(testName, func(t *testing.T) {

			// Test be get expected result from prob on specific instance (x).
			p := GetProblem(test.id, 9, 0)
			if got := testEval(p); math.Abs(got-test.score) > eps {
				t.Errorf("%s got %f wanted %f", testName, got, test.score)
			}
			// nEvals must always be 1
			if got := p.State().nEvals; got != 1 {
				t.Errorf("%s nEvals got %d wanted %d", testName, got, 1)
			}
		})
	}

}

func TestAproxProblems(t *testing.T) {

	tests := []struct {
		id    int
		score float64
	}{

		// {4, -22.7},
		// {5, -20.18},
		// {11, -22.7},
		// {12, -25.2},
		// {24, -29.1},
		{25, -32.0},
	}

	eps := 0.5
	n := 10 //1000000 //+0

	for _, test := range tests {

		testName := fmt.Sprintf("%d", test.id)

		t.Run(testName, func(t *testing.T) {

			avg := 0.0
			for i := 0; i < n; i++ {
				p := GetProblem(test.id, 9, i)
				fmt.Println(testEval(p))
				avg += testEval(p)
			}
			fmt.Println(avg)
			avg /= float64(n)

			fmt.Println(testName, avg)

			if math.Abs(avg-test.score) > eps {
				t.Errorf("%s got %f wanted %f", testName, avg, test.score)
			}
		})
	}

}
