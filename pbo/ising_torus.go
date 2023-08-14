package pbo

import (
	"math"
)

// Id: 20
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/ising_torus.hpp
type IsingTorusProblem struct {
	state ProblemState
}

func (p *IsingTorusProblem) Init(dim int) {

}

func moduloIsingTorus(x, n int) int {
	return (x%n + n) % n
}

func (p *IsingTorusProblem) Eval(x []bool) float64 {

	result := 0.0
	var neighbors [2]float64
	doubleN := float64(len(x))
	latticeSize := int(math.Sqrt(doubleN))

	if math.Floor(math.Sqrt(doubleN)) != math.Sqrt(doubleN) {

		panic("Number of parameters in the Ising square problem must be a square number")
	}

	for i := 0; i < latticeSize; i++ {
		for j := 0; j < latticeSize; j++ {
			neighbors[0] = BoolAsFloat(x[moduloIsingTorus(i+1, latticeSize)*latticeSize+j])
			neighbors[1] = BoolAsFloat(x[latticeSize*i+moduloIsingTorus(j+1, latticeSize)])
			for _, neighbor := range neighbors {
				result += BoolAsFloat(x[latticeSize*i+j])*(neighbor) +
					(1.0-BoolAsFloat(x[latticeSize*i+j]))*(1.0-(neighbor))
			}
		}
	}

	// Update state
	p.state.nEvals++

	return result
}

func (p *IsingTorusProblem) State() ProblemState {
	return p.state
}

func (p *IsingTorusProblem) Reset() {
	p.state.Reset()
}
