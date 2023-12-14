package math

import (
	"math"
)

type (
	Vec4 [4]float64
)

const (
	EPSILON = 0.0001
)

func Point4(x, y, z float64) Vec4 {
	return Vec4{x, y, z, 1.0}
}

func Vector4(x, y, z float64) Vec4 {
	return Vec4{x, y, z, 0.0}
}

func Vec4With(v float64) Vec4 {
	return Vec4{v, v, v, v}
}

func (lhs Vec4) Add(rhs Vec4) Vec4 {
	return Vec4{lhs[0]+rhs[0], lhs[1]+rhs[1], lhs[2]+rhs[2], lhs[3]+rhs[3]}
}

func (lhs Vec4) Sub(rhs Vec4) Vec4 {
	return Vec4{lhs[0]-rhs[0], lhs[1]-rhs[1], lhs[2]-rhs[2], lhs[3]-rhs[3]}
}

func (v Vec4) Mul(x float64) Vec4 {
	return Vec4{v[0]*x, v[1]*x, v[2]*x, v[3]*x}
}

func (v Vec4) Div(x float64) Vec4 {
	return Vec4{v[0]/x, v[1]/x, v[2]/x, v[3]/x}
}

func (v Vec4) Negate() Vec4 {
	return v.Mul(-1)
}

func (v Vec4) Magnitude() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3])
}

func (v Vec4) Normalize() Vec4 {
	return v.Div(v.Magnitude())
}

func (lhs Vec4) Approx(rhs Vec4) bool {
	for i := range lhs {
		if !eqApprox(lhs[i], rhs[i]) {
			return false
		}
	}
	return true
}

func eqApprox(a, b float64) bool {
	if math.Abs(a - b) < EPSILON {
		return true
	}
	return false
}
