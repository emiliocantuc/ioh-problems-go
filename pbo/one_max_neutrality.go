package pbo

// Id: 6
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/one_max_neutrality.hpp
type OneMaxNeutralityProblem struct {
	state ProblemState
}

func (p *OneMaxNeutralityProblem) Init(dim int) {
}

func (p *OneMaxNeutralityProblem) Eval(x []bool) float64 {

	newVariables := Neutrality(x, 3)
	result := 0.0
	for _, value := range newVariables {
		result += BoolAsFloat(value)
	}

	// Update state
	p.state.nEvals++

	return result
}

func (p *OneMaxNeutralityProblem) State() ProblemState {
	return p.state
}

func (p *OneMaxNeutralityProblem) Reset() {
	p.state.Reset()
}
