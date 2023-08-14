package pbo

import "math/rand"

// Id: 0
// Custom added
type KnapsackProblem struct {
	state           ProblemState
	maxWeight       float64
	weights, values []float64
}

func (p *KnapsackProblem) Init(dim int) {

	p.maxWeight = float64(dim) / 3
	p.weights = make([]float64, dim)
	p.values = make([]float64, dim)

	r := rand.New(rand.NewSource(0)) // TODO set seed param
	for i := range p.weights {
		p.weights[i] = r.Float64()
		p.values[i] = r.Float64()
	}
}

func (p *KnapsackProblem) Eval(x []bool) float64 {

	totalWeight := 0.0
	totalValue := 0.0
	for i := range x {
		if x[i] {
			totalWeight += p.weights[i]
			totalValue += p.values[i]
		}
	}

	if totalWeight > p.maxWeight {
		totalValue = 0.0
	}

	// Update state
	p.state.nEvals++

	return totalValue
}

func (p *KnapsackProblem) State() ProblemState {
	return p.state
}

func (p *KnapsackProblem) Reset() {
	p.state.Reset()
}
