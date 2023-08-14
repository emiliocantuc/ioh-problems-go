package pbo

// Id: 11
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/leading_ones_dummy1.hpp
type LeadingOnesDummy1Problem struct {
	state ProblemState
	info  []int
}

func (p *LeadingOnesDummy1Problem) Init(dim int) {
	p.info = Dummy(dim, 0.5, 10000)
}

func (p *LeadingOnesDummy1Problem) Eval(x []bool) float64 {

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

func (p *LeadingOnesDummy1Problem) State() ProblemState {
	return p.state
}

func (p *LeadingOnesDummy1Problem) Reset() {
	p.state.Reset()
}
