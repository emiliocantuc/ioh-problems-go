package pbo

// Id: 10
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/one_max_ruggedness3.hpp
type OneMaxRuggedness3Problem struct {
	state ProblemState
	info_ []float64
}

func (p *OneMaxRuggedness3Problem) Init(dim int) {
	p.info_ = Ruggedness3(dim)
}

func (p *OneMaxRuggedness3Problem) Eval(x []bool) float64 {

	result := 0.0
	for _, xi := range x {
		result += BoolAsFloat(xi)
	}
	// Update state
	p.state.nEvals++

	return p.info_[int(result+0.5)]
}

func (p *OneMaxRuggedness3Problem) State() ProblemState {
	return p.state
}

func (p *OneMaxRuggedness3Problem) Reset() {
	p.state.Reset()
}
