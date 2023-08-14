package pbo

import (
	"fmt"
)

type Problem interface {
	Init(dim int)
	Eval(x []bool) float64
	State() ProblemState
	Reset()
}

func GetProblem(id, dim, instance int) Problem {

	var p Problem
	switch id {
	case 1:
		p = &OneMaxProblem{}
	case 2:
		p = &LeadingOnesProblem{}
	case 3:
		p = &LinearProblem{}
	case 4:
		p = &OneMaxDummy1Problem{}
	case 5:
		p = &OneMaxDummy2Problem{}
	case 6:
		p = &OneMaxNeutralityProblem{}
	case 7:
		p = &OneMaxEpistasisProblem{}
	case 8:
		p = &OneMaxRuggedness1Problem{}
	case 9:
		p = &OneMaxRuggedness2Problem{}
	case 10:
		p = &OneMaxRuggedness3Problem{}
	case 11:
		p = &LeadingOnesDummy1Problem{}
	case 12:
		p = &LeadingOnesDummy2Problem{}
	case 13:
		p = &LeadingOnesNeutralityProblem{}
	case 14:
		p = &LeadingOnesEpistasisProblem{}
	case 15:
		p = &LeadingOnesRuggedness1Problem{}
	case 16:
		p = &LeadingOnesRuggedness2Problem{}
	case 17:
		p = &LeadingOnesRuggedness3Problem{}
	case 18:
		p = &LABSProblem{}
	case 19:
		p = &IsingRingProblem{}
	case 20:
		p = &IsingTorusProblem{}
	case 21:
		p = &IsingTriangularProblem{}
	case 22:
		p = &MISProblem{}
	case 23:
		p = &NQueensProblem{}
	case 24:
		p = &ConcatenatedTrapProblem{}
	case 25:
		p = &NKLandscapesProblem{}
	default:
		panic("Id not registered to any problem.")
	}

	p.Init(dim)
	return p
}

type ProblemState struct {
	nEvals int
}

func (p *ProblemState) Reset() {
	p.nEvals = 0
}

// UTILS
func BoolAsFloat(x bool) float64 {
	if x {
		return 1.0
	}
	return 0.0
}

func BoolAsInt(x bool) int {
	if x {
		return 1
	}
	return 0
}

func main() {

	x := []bool{true, true, false, true, false, false, false, true, true}
	p := GetProblem(4, 9, 0)
	fmt.Println(p.Eval(x))
	p = GetProblem(4, 9, 0)
	fmt.Println(p.Eval(x))

}
