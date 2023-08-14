package pbo

// Id: 7
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/one_max_epistasis.hpp
type OneMaxEpistasisProblem struct {
	state ProblemState
}

func (p *OneMaxEpistasisProblem) Init(dim int) {
}

func (p *OneMaxEpistasisProblem) Eval(x []bool) float64 {

	newVariables := Epistasis(x, 4)
	result := 0.0
	for _, value := range newVariables {
		result += BoolAsFloat(value)
	}

	// Update state
	p.state.nEvals++

	return result
}

func (p *OneMaxEpistasisProblem) State() ProblemState {
	return p.state
}

func (p *OneMaxEpistasisProblem) Reset() {
	p.state.Reset()
}
