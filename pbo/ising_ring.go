package pbo

// Id: 19
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/ising_ring.hpp
type IsingRingProblem struct {
	state ProblemState
}

func (p *IsingRingProblem) Init(dim int) {

}

func moduloIsingRing(x, n int) int {
	return (x%n + n) % n
}

func (p *IsingRingProblem) Eval(x []bool) float64 {

	result := 0.0
	n := len(x)
	for i := 0; i < n; i++ {
		neighbors := BoolAsFloat(x[moduloIsingRing(i-1, n)])
		result += BoolAsFloat(x[i])*neighbors + (1.0-BoolAsFloat(x[i]))*(1.0-neighbors)
	}

	// Update state
	p.state.nEvals++

	return result
}

func (p *IsingRingProblem) State() ProblemState {
	return p.state
}

func (p *IsingRingProblem) Reset() {
	p.state.Reset()
}
