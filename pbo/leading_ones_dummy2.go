package pbo

// Id: 12
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/leading_ones_dummy2.hpp
type LeadingOnesDummy2Problem struct {
	state ProblemState
	info  []int
}

func (p *LeadingOnesDummy2Problem) Init(dim int) {
	p.info = Dummy(dim, 0.9, 10000)
}

func (p *LeadingOnesDummy2Problem) Eval(x []bool) float64 {

	result := 0.0
	for i := 0; i < len(p.info); i++ {
		if x[p.info[i]] {
			result = float64(i) + 1.0
		} else {
			break
		}
	}
	// Update state
	p.state.nEvals++

	return float64(result)
}

func (p *LeadingOnesDummy2Problem) State() ProblemState {
	return p.state
}

func (p *LeadingOnesDummy2Problem) Reset() {
	p.state.Reset()
}
