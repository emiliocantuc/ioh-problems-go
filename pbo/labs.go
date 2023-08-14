package pbo

// Id: 18
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/labs.hpp
type LABSProblem struct {
	state ProblemState
}

func (p *LABSProblem) Init(dim int) {

}

func correlation(x []bool, n, k int) float64 {
	result := 0.0
	for i := 0; i < n-k; i++ {
		result += (BoolAsFloat(x[i])*2.0 - 1.0) * (BoolAsFloat(x[i+k])*2.0 - 1.0)
	}
	return result
}

func (p *LABSProblem) Eval(x []bool) float64 {

	result := 0.0
	n := len(x)
	for k := 1; k < n; k++ {
		cor := correlation(x, n, k)
		result += (cor * cor)
	}
	result = (float64(n) * float64(n)) / (2.0 * result)

	// Update state
	p.state.nEvals++

	return result
}

func (p *LABSProblem) State() ProblemState {
	return p.state
}

func (p *LABSProblem) Reset() {
	p.state.Reset()
}
