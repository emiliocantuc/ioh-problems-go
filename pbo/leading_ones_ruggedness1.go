package pbo

// Id: 15
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/leading_ones_ruggedness1.hpp
type LeadingOnesRuggedness1Problem struct {
	state ProblemState
}

func (p *LeadingOnesRuggedness1Problem) Init(dim int) {
}

func (p *LeadingOnesRuggedness1Problem) Eval(x []bool) float64 {

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

	return Ruggedness1(float64(result), len(x))
}

func (p *LeadingOnesRuggedness1Problem) State() ProblemState {
	return p.state
}

func (p *LeadingOnesRuggedness1Problem) Reset() {
	p.state.Reset()
}
