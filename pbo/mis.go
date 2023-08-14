package pbo

// Id: 22
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/mis.hpp
type MISProblem struct {
	state ProblemState
}

func (p *MISProblem) Init(dim int) {

}

func isEdge(i, j, problemSize int) int {
	if i != problemSize/2 && j == i+1 {
		return 1
	}
	if i <= problemSize/2-1 && j == i+problemSize/2+1 {
		return 1
	}
	if i <= problemSize/2 && i >= 2 && j == i+problemSize/2-1 {
		return 1
	}
	return 0
}

func (p *MISProblem) Eval(x []bool) float64 {

	numOfOnes := 0
	sumEdgesInTheSet := 0
	numOfVariablesEven := len(x)
	onesArray := make([]int, numOfVariablesEven+1)

	if numOfVariablesEven%2 != 0 {
		numOfVariablesEven--
	}

	for index := 0; index < numOfVariablesEven; index++ {
		if x[index] {
			onesArray[numOfOnes] = index
			numOfOnes++
		}
	}

	for i := 0; i < numOfOnes; i++ {
		for j := i + 1; j < numOfOnes; j++ {
			if isEdge(onesArray[i]+1, onesArray[j]+1, numOfVariablesEven) == 1 {
				sumEdgesInTheSet++
			}
		}
	}
	result := float64(numOfOnes) - float64(numOfVariablesEven)*float64(sumEdgesInTheSet)

	// Update state
	p.state.nEvals++

	return result
}

func (p *MISProblem) State() ProblemState {
	return p.state
}

func (p *MISProblem) Reset() {
	p.state.Reset()
}
