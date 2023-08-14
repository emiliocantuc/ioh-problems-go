package pbo

// Id: 8
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/one_max_ruggedness1.hpp
type OneMaxRuggedness1Problem struct {
	state ProblemState
}

func (p *OneMaxRuggedness1Problem) Init(dim int) {
}

func (p *OneMaxRuggedness1Problem) Eval(x []bool) float64 {

	result := 0.0
	for i := 0; i < len(x); i++ {
		if x[i] {
			result += 1.0
		}
	}
	// Update state
	p.state.nEvals++

	return Ruggedness1(result, len(x))
}

func (p *OneMaxRuggedness1Problem) State() ProblemState {
	return p.state
}

func (p *OneMaxRuggedness1Problem) Reset() {
	p.state.Reset()
}
