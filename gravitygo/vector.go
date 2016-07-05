//package vector
package gravitygo

import "math"

type Vector struct {
	vector [3]float64
}

func (v *Vector) GetX() float64 {
	return v.vector[0]
}

func (v *Vector) GetY() float64 {
	return v.vector[1]
}

func (v *Vector) GetZ() float64 {
	return v.vector[2]
}

func (v *Vector) Norm() float64 {

	var result float64 = 0

	for i := 0; i < len(v.vector); i++ {
		result += math.Pow(v.vector[i], 2)
	}

	return math.Pow(result, 1.5)
}

func (v *Vector) NormSquared() float64 {

	var result float64 = 0

	for i := 0; i < len(v.vector); i++ {
		result += math.Pow(v.vector[i], 2)
	}

	return result
}

func VectorAdd(LHSVector, RHSVector Vector) Vector {
	var result Vector

	for i := 0; i < len(result.vector); i++ {
		result.vector[i] = LHSVector.vector[i] + RHSVector.vector[i]
	}

	return result
}

func VectorSubstract(LHSVector, RHSVector Vector) Vector {
	var result Vector

	for i := 0; i < len(result.vector); i++ {
		result.vector[i] = LHSVector.vector[i] - RHSVector.vector[i]
	}

	return result
}
func (v *Vector) MakeVector(vector_ [3]float64) {
	v.vector = vector_
}
