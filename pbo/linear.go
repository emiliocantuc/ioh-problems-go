package pbo

// Id: 3
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/linear.hpp
type LinearProblem struct {
	state ProblemState
}

func (p *LinearProblem) Init(dim int) {

}

func (p *LinearProblem) Eval(x []bool) float64 {

	result := 0.0
	for i := 0; i < len(x); i++ {
		result += BoolAsFloat(x[i]) * (float64(i) + 1.0)
	}
	// Update state
	p.state.nEvals++

	return result
}

func (p *LinearProblem) State() ProblemState {
	return p.state
}

func (p *LinearProblem) Reset() {
	p.state.Reset()
}
