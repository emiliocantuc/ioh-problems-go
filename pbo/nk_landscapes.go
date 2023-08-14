package pbo

import (
	"math"
)

// Id: 25
// Taken from: https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/pbo/nk_landscapes.hpp
type NKLandscapesProblem struct {
	state ProblemState
	f_    [][]float64
	e_    [][]int
}

func (p *NKLandscapesProblem) Init(dim int) {
	p.f_ = make([][]float64, 0)
	p.e_ = make([][]int, 0)
	p.setNK(dim, 1)
}

func (p *NKLandscapesProblem) setNK(n, k int) {

	if k > n {
		panic("k <= n")
	}

	for i := 0; i < n; i++ {
		randVec := make([]float64, k)
		for j := 0; j < k; j++ {
			randVec[j] = float64(j+1) * float64(i+1)
		}

		sampledNumber := make([]int, 0)
		population := make([]int, n)
		for j := 0; j < n; j++ {
			population[j] = j
		}

		for i1 := n - 1; i1 > 0; i1-- {
			randPos := int(math.Floor(randVec[n-1-i1] * float64(i1+1)))
			temp := population[i1]
			population[i1] = population[randPos]
			population[randPos] = temp
			sampledNumber = append(sampledNumber, population[i1])
			if n-i1-1 == k-1 {
				break
			}
		}
		if n == k {
			sampledNumber = append(sampledNumber, population[0])
		}
		p.e_ = append(p.e_, sampledNumber)
	}
	for i := 0; i < n; i++ {
		f_i := make([]float64, int(math.Pow(2, float64(k+1))))
		for j := 0; j < int(math.Pow(2, float64(k+1))); j++ {
			f_i[j] = float64(j+1) * float64(k) * float64(i+1) * 2
		}
		p.f_ = append(p.f_, f_i)
	}
}

func (p *NKLandscapesProblem) Eval(x []bool) float64 {
	k := 1
	result := 0.0
	for i := 0; i < len(x); i++ {
		index := BoolAsInt(x[i])
		for j := 0; j < k; j++ {
			index = index + int(math.Pow(2, float64(j+1))*BoolAsFloat(x[p.e_[i][j]]))
		}
		result += p.f_[i][index]
	}

	result = -result / float64(len(x))

	// Update state
	p.state.nEvals++

	return result
}

func (p *NKLandscapesProblem) State() ProblemState {
	return p.state
}

func (p *NKLandscapesProblem) Reset() {
	p.state.Reset()
}
