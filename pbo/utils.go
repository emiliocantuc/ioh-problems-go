package pbo

import (
	"math"
	"math/rand"
	"sort"
)

// https://github.com/IOHprofiler/IOHexperimenter/blob/master/include/ioh/problem/utils.hpp

// Dummy returns a set of indices selected randomly from the range [0, nVariables].
func Dummy(nVariables int, selectRate float64, seed int64) []int {

	selectNum := int(math.Floor(float64(nVariables) * selectRate))

	r := rand.New(rand.NewSource(seed))

	randomNumbers := make([]float64, selectNum)
	for i := 0; i < selectNum; i++ {
		randomNumbers[i] = r.Float64()
	}

	position := make([]int, nVariables)
	for i := 0; i < nVariables; i++ {
		position[i] = i
	}

	for i := 0; i < selectNum; i++ {
		randomIndex := int(math.Floor(randomNumbers[i] * float64(nVariables)))
		position[i], position[randomIndex] = position[randomIndex], position[i]
	}

	sort.Ints(position[:selectNum])
	return position[:selectNum]
}

// Neutrality simulates the neutrality transformation on input vector x with parameter mu.
func Neutrality(x []bool, mu int) []bool {
	nVariables := len(x)
	n := int(math.Floor(float64(nVariables) / float64(mu)))

	newVariables := make([]bool, 0, n)
	cumSum := 0.0

	for i := 0; i < nVariables; i++ {
		if x[i] {
			cumSum++
		}
		if (i+1)%mu == 0 && i != 0 {
			newVariables = append(newVariables, cumSum >= float64(mu)/2.0)
			cumSum = 0.0
		}
	}
	return newVariables
}

// Epistasis simulates the epistasis transformation on input vector variables with parameter v.
func Epistasis(variables []bool, v int) []bool {
	numberofVariables := len(variables)
	var epistasisResult bool
	newVariables := make([]bool, 0, numberofVariables)
	h := 0
	for h+v-1 < numberofVariables {
		i := 0
		for i < v {
			epistasisResult = false
			for j := 0; j < v; j++ {
				if v-j-1 != (v-i-1-1)%4 {
					if !epistasisResult {
						epistasisResult = variables[j+h]
					} else {
						epistasisResult = epistasisResult != variables[j+h]
					}
				}
			}
			newVariables = append(newVariables, epistasisResult)
			i++
		}
		h += v
	}
	if numberofVariables-h > 0 {
		v = numberofVariables - h
		for i := 0; i < v; i++ {
			epistasisResult = false
			for j := 0; j < v; j++ {
				if v-j-1 != (v-i-1-1)%4 {
					if !epistasisResult {
						epistasisResult = variables[h+j]
					} else {
						epistasisResult = epistasisResult != variables[h+j]
					}
				}
			}
			newVariables = append(newVariables, epistasisResult)
		}
	}
	return newVariables
}

func Ruggedness1(y float64, number_of_variables int) float64 {
	var ruggedness_y float64
	s := float64(number_of_variables)
	if y == s {
		ruggedness_y = math.Ceil(y/2.0) + 1.0
	} else if y < s && number_of_variables%2 == 0 {
		ruggedness_y = math.Floor(y/2.0) + 1.0
	} else if y < s && number_of_variables%2 != 0 {
		ruggedness_y = math.Ceil(y/2.0) + 1.0
	} else {
		ruggedness_y = y
		if y > s {
			panic("Assertion failed: y should be less than or equal to s")
		}
	}
	return ruggedness_y
}

func Ruggedness2(y float64, number_of_variables int) float64 {
	var ruggedness_y float64
	tempy := int(y + 0.5)
	if tempy == number_of_variables {
		ruggedness_y = y
	} else if tempy < number_of_variables && tempy%2 == 0 && number_of_variables%2 == 0 {
		ruggedness_y = y + 1.0
	} else if tempy < number_of_variables && tempy%2 == 0 && number_of_variables%2 != 0 {
		if y-1.0 > 0 {
			ruggedness_y = y - 1.0
		} else {
			ruggedness_y = 0
		}
	} else if tempy < number_of_variables && tempy%2 != 0 && number_of_variables%2 == 0 {
		if y-1.0 > 0 {
			ruggedness_y = y - 1.0
		} else {
			ruggedness_y = 0
		}
	} else if tempy < number_of_variables && tempy%2 != 0 && number_of_variables%2 != 0 {
		ruggedness_y = y + 1.0
	} else {
		ruggedness_y = y
		if tempy > number_of_variables {
			panic("Assertion failed: tempy should be less than or equal to number_of_variables")
		}
	}
	return ruggedness_y
}

func Ruggedness3(number_of_variables int) []float64 {
	ruggedness_fitness := make([]float64, number_of_variables+1)

	for j := 1; j <= number_of_variables/5; j++ {
		for k := 0; k < 5; k++ {
			ruggedness_fitness[number_of_variables-5*j+k] = float64(number_of_variables - 5*j + (4 - k))
		}
	}
	for k := 0; k < number_of_variables-number_of_variables/5*5; k++ {
		ruggedness_fitness[k] = float64(number_of_variables - number_of_variables/5*5 - 1 - k)
	}
	ruggedness_fitness[number_of_variables] = float64(number_of_variables)
	return ruggedness_fitness
}
