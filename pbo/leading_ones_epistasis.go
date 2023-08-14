package pbo

// Id: 14
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/leading_ones_epistasis.hpp
type LeadingOnesEpistasisProblem struct {
	state ProblemState
}

func (p *LeadingOnesEpistasisProblem) Init(dim int) {
}

func (p *LeadingOnesEpistasisProblem) Eval(x []bool) float64 {

	newVariables := Epistasis(x, 4)
	result := 0.0
	for i, value := range newVariables {
		if value {
			result = float64(i) + 1.0
		} else {
			break
		}
	}
	// Update state
	p.state.nEvals++

	return float64(result)
}

func (p *LeadingOnesEpistasisProblem) State() ProblemState {
	return p.state
}

func (p *LeadingOnesEpistasisProblem) Reset() {
	p.state.Reset()
}
