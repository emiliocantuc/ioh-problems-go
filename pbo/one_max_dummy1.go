package pbo

// Id: 4
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/one_max_dummy1.hpp
type OneMaxDummy1Problem struct {
	state ProblemState
	info  []int
}

func (p *OneMaxDummy1Problem) Init(dim int) {
	p.info = Dummy(dim, 0.5, 10000)
}

func (p *OneMaxDummy1Problem) Eval(x []bool) float64 {

	result := 0.0
	for _, i := range p.info {
		result += BoolAsFloat(x[i])
	}

	// Update state
	p.state.nEvals++

	return result
}

func (p *OneMaxDummy1Problem) State() ProblemState {
	return p.state
}

func (p *OneMaxDummy1Problem) Reset() {
	p.state.Reset()
}
