package pbo

// Id: 5
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/one_max_dummy2.hpp
type OneMaxDummy2Problem struct {
	state ProblemState
	info  []int
}

func (p *OneMaxDummy2Problem) Init(dim int) {
	p.info = Dummy(dim, 0.9, 10000)
}

func (p *OneMaxDummy2Problem) Eval(x []bool) float64 {

	result := 0.0
	for _, i := range p.info {
		result += BoolAsFloat(x[i])
	}

	// Update state
	p.state.nEvals++

	return result
}

func (p *OneMaxDummy2Problem) State() ProblemState {
	return p.state
}

func (p *OneMaxDummy2Problem) Reset() {
	p.state.Reset()
}
