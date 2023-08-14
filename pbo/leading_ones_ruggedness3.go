package pbo

// Id: 17
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/leading_ones_ruggedness3.hpp
type LeadingOnesRuggedness3Problem struct {
	state ProblemState
	info_ []float64
}

func (p *LeadingOnesRuggedness3Problem) Init(dim int) {
	p.info_ = Ruggedness3(dim)
}

func (p *LeadingOnesRuggedness3Problem) Eval(x []bool) float64 {

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

	return p.info_[int(float64(result)+0.5)]
}

func (p *LeadingOnesRuggedness3Problem) State() ProblemState {
	return p.state
}

func (p *LeadingOnesRuggedness3Problem) Reset() {
	p.state.Reset()
}
