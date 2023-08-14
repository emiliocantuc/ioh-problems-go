package pbo

// Id: 16
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/leading_ones_ruggedness2.hpp
type LeadingOnesRuggedness2Problem struct {
	state ProblemState
}

func (p *LeadingOnesRuggedness2Problem) Init(dim int) {
}

func (p *LeadingOnesRuggedness2Problem) Eval(x []bool) float64 {

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

	return Ruggedness2(float64(result), len(x))
}

func (p *LeadingOnesRuggedness2Problem) State() ProblemState {
	return p.state
}

func (p *LeadingOnesRuggedness2Problem) Reset() {
	p.state.Reset()
}
