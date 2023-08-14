package pbo

// Id: 2
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/leading_ones.hpp
type LeadingOnesProblem struct {
	state ProblemState
}

func (p *LeadingOnesProblem) Init(dim int) {
}

func (p *LeadingOnesProblem) Eval(x []bool) float64 {

	result := 0
	for i := 0; i < len(x); i++ {
		if x[i] {
			result = i + 1
		} else {
			break
		}
	}
	// Update state
	p.state.nEvals++

	return float64(result)
}

func (p *LeadingOnesProblem) State() ProblemState {
	return p.state
}

func (p *LeadingOnesProblem) Reset() {
	p.state.Reset()
}
