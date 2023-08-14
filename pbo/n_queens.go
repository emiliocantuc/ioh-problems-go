package pbo

import (
	"math"
)

// Id: 23
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/n_queens.hpp
type NQueensProblem struct {
	state ProblemState
}

func (p *NQueensProblem) Init(dim int) {
}

func (p *NQueensProblem) Eval(x []bool) float64 {

	var j, i int
	nQueens := int(math.Sqrt(float64(len(x))) + 0.5)
	var numberOfQueensOnBoard int
	var kPenalty, lPenalty, rowsPenalty, columnsPenalty float64
	var index int
	c := float64(nQueens)

	if math.Floor(math.Sqrt(float64(len(x)))) != math.Sqrt(float64(len(x))) {
		panic("Number of parameters in the N Queen problem must be a square number")
	}

	for _, val := range x {
		if val {
			numberOfQueensOnBoard++
		}
	}

	for j = 1; j <= nQueens; j++ {
		var sumColumn float64
		for i = 1; i <= nQueens; i++ {
			index = (i-1)*nQueens + (j-1)%nQueens
			sumColumn += BoolAsFloat(x[index])
		}
		columnsPenalty += math.Max(0.0, -1.0+sumColumn)
	}

	for i = 1; i <= nQueens; i++ {
		var sumRaw float64
		for j = 1; j <= nQueens; j++ {
			index = (i-1)*nQueens + (j-1)%nQueens
			sumRaw += BoolAsFloat(x[index])
		}
		rowsPenalty += math.Max(0.0, -1.0+sumRaw)
	}

	for k := 2 - nQueens; k <= nQueens-2; k++ {
		var sumK float64
		for i = 1; i <= nQueens; i++ {
			if k+i >= 1 && k+i <= nQueens {
				index = (i-1)*nQueens + (k+i-1)%nQueens
				sumK += BoolAsFloat(x[index])
			}
		}
		kPenalty += math.Max(0.0, -1.0+sumK)
	}

	for l := 3; l <= 2*nQueens-1; l++ {
		var sumL float64
		for i = 1; i <= nQueens; i++ {
			if l-i >= 1 && l-i <= nQueens {
				index = (i-1)*nQueens + (l-i-1)%nQueens
				sumL += BoolAsFloat(x[index])
			}
		}
		lPenalty += math.Max(0.0, -1.0+sumL)
	}

	result := float64(numberOfQueensOnBoard) - c*rowsPenalty - c*columnsPenalty - c*kPenalty - c*lPenalty

	// Update state
	p.state.nEvals++

	return result
}

func (p *NQueensProblem) State() ProblemState {
	return p.state
}

func (p *NQueensProblem) Reset() {
	p.state.Reset()
}
