package pbo

// Id: 13
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/leading_ones_neutrality.hpp
type LeadingOnesNeutralityProblem struct {
	state ProblemState
}

func (p *LeadingOnesNeutralityProblem) Init(dim int) {
}

func (p *LeadingOnesNeutralityProblem) Eval(x []bool) float64 {

	newVariables := Neutrality(x, 3)
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

func (p *LeadingOnesNeutralityProblem) State() ProblemState {
	return p.state
}

func (p *LeadingOnesNeutralityProblem) Reset() {
	p.state.Reset()
}
