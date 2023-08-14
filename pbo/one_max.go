package pbo

// Id: 1
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/one_max.hpp
type OneMaxProblem struct {
	state ProblemState
}

func (p *OneMaxProblem) Init(dim int) {

}

func (p *OneMaxProblem) Eval(x []bool) float64 {

	result := 0.0
	for _, val := range x {
		result += BoolAsFloat(val)
	}
	// Update state
	p.state.nEvals++

	return result
}

func (p *OneMaxProblem) State() ProblemState {
	return p.state
}

func (p *OneMaxProblem) Reset() {
	p.state.Reset()
}
