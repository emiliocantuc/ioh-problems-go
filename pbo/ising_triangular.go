package pbo

import "math"

// Id: 21
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/ising_triangular.hpp
type IsingTriangularProblem struct {
	state ProblemState
}

func (p *IsingTriangularProblem) Init(dim int) {

}

func moduloIsingTriangular(x, n int) int {
	return (x%n + n) % n
}

func (p *IsingTriangularProblem) Eval(x []bool) float64 {

	result := 0.0
	var neighbors [3]float64
	latticeSize := int(math.Sqrt(float64(len(x))))

	for i := 0; i < latticeSize; i++ {
		for j := 0; j < latticeSize; j++ {
			neighbors[0] = BoolAsFloat(x[moduloIsingTriangular(i+1, latticeSize)*latticeSize+j])
			neighbors[1] = BoolAsFloat(x[i*latticeSize+moduloIsingTriangular(j+1, latticeSize)])
			neighbors[2] = BoolAsFloat(x[moduloIsingTriangular(i+1, latticeSize)*latticeSize+
				moduloIsingTriangular(j+1, latticeSize)])

			for _, neighbor := range neighbors {
				if x[i*latticeSize+j] {
					result += float64(neighbor)
				} else {
					result += 1.0 - float64(neighbor)
				}
			}
		}
	}
	// Update state
	p.state.nEvals++

	return result
}

func (p *IsingTriangularProblem) State() ProblemState {
	return p.state
}

func (p *IsingTriangularProblem) Reset() {
	p.state.Reset()
}
