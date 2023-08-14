package pbo

// Id: 9
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/one_max_ruggedness2.hpp
type OneMaxRuggedness2Problem struct {
	state ProblemState
}

func (p *OneMaxRuggedness2Problem) Init(dim int) {
}

func (p *OneMaxRuggedness2Problem) Eval(x []bool) float64 {

	result := 0.0
	for i := 0; i < len(x); i++ {
		if x[i] {
			result += 1.0
		}
	}
	// Update state
	p.state.nEvals++

	return Ruggedness2(result, len(x))
}

func (p *OneMaxRuggedness2Problem) State() ProblemState {
	return p.state
}

func (p *OneMaxRuggedness2Problem) Reset() {
	p.state.Reset()
}
