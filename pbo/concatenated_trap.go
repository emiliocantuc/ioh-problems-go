package pbo

// Id: 24
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/concatenated_trap.hpp
type ConcatenatedTrapProblem struct {
	state ProblemState
}

func (p *ConcatenatedTrapProblem) Init(dim int) {

}

func (p *ConcatenatedTrapProblem) Eval(x []bool) float64 {

	k := 5
	result := 0.0
	var block_result float64
	m := len(x) / k

	for i := 1; i <= m; i++ {
		block_result = 0.0
		for j := i*k - k; j < i*k; j++ {
			block_result += BoolAsFloat(x[j])
		}
		if block_result == float64(k) {
			result += 1
		} else {
			result += (float64(k) - 1.0 - block_result) / float64(k)
		}
	}

	remain_k := len(x) - m*k
	if remain_k != 0 {
		block_result = 0.0
		for j := m * (k - 1); j < len(x); j++ {
			block_result += BoolAsFloat(x[j])
		}
		if block_result == float64(remain_k) {
			result += 1
		} else {
			result += (float64(remain_k) - 1.0 - block_result) / float64(remain_k)
		}
	}
	// Update state
	p.state.nEvals++

	return result
}

func (p *ConcatenatedTrapProblem) State() ProblemState {
	return p.state
}

func (p *ConcatenatedTrapProblem) Reset() {
	p.state.Reset()
}
