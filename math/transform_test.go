package math

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslation(t *testing.T) {
	actual := Trans(Vec3{5, -3, 2}).MulVec(Point4(-3, 4, 5))
	expected := Point4(2, 1, 7)
	assert.Equal(t, expected, actual)
}

func TestInverseTranslation(t *testing.T) {
	actual := Trans(Vec3{5, -3, 2}).Inv().MulVec(Point4(-3, 4, 5))
	expected := Point4(-8, 7, 3)
	assert.Equal(t, expected, actual)
}

func TestTranslateVectorDoesNotChange(t *testing.T) {
	actual := Trans(Vec3{5, -3, 2}).MulVec(Vector4(-3, 4, 5))
	expected := Vector4(-3, 4, 5)
	assert.Equal(t, expected, actual)
}

func TestScaling(t *testing.T) {
	actual := Scale(Vec3{2, 3, 4}).MulVec(Point4(-4, 6, 8))
	expected := Point4(-8, 18, 32)
	assert.Equal(t, expected, actual)
}

func TestScalingVector(t *testing.T) {
	actual := Scale(Vec3{2, 3, 4}).MulVec(Vector4(-4, 6, 8))
	expected := Vector4(-8, 18, 32)
	assert.Equal(t, expected, actual)
}

func TestInverseScaling(t *testing.T) {
	actual := Scale(Vec3{2, 3, 4}).Inv().MulVec(Vector4(-4, 6, 8))
	expected := Vector4(-2, 2, 2)
	assert.Equal(t, expected, actual)
}

func TestScaleNegativeIsReflection(t *testing.T) {
	actual := Scale(Vec3{-1, 1, 1}).MulVec(Point4(2, 3, 4))
	expected := Point4(-2, 3, 4)
	assert.Equal(t, expected, actual)
}

func TestRotationAroundX(t *testing.T) {
	p := Point4(0, 1, 0)
	halfQuarter := RotX(math.Pi/4.0)
	fullQuarter := RotX(math.Pi/2.0)
	expectedHalf := Point4(0, math.Sqrt2/2.0, math.Sqrt2/2.0)
	expectedFull := Point4(0, 0, 1)
	assert.True(t, expectedHalf.ApproxEqual(halfQuarter.MulVec(p)))
	assert.True(t, expectedFull.ApproxEqual(fullQuarter.MulVec(p)))
}

func TestRotationAroundY(t *testing.T) {
	p := Point4(0, 0, 1)
	halfQuarter := RotY(math.Pi/4.0)
	fullQuarter := RotY(math.Pi/2.0)
	expectedHalf := Point4(math.Sqrt2/2.0, 0, math.Sqrt2/2.0)
	expectedFull := Point4(1, 0, 0)
	assert.True(t, expectedHalf.ApproxEqual(halfQuarter.MulVec(p)))
	assert.True(t, expectedFull.ApproxEqual(fullQuarter.MulVec(p)))
}

func TestRotationAroundZ(t *testing.T) {
	p := Point4(0, 1, 0)
	halfQuarter := RotZ(math.Pi/4.0)
	fullQuarter := RotZ(math.Pi/2.0)
	expectedHalf := Point4(-math.Sqrt2/2.0, math.Sqrt2/2.0, 0)
	expectedFull := Point4(-1, 0, 0)
	assert.True(t, expectedHalf.ApproxEqual(halfQuarter.MulVec(p)))
	assert.True(t, expectedFull.ApproxEqual(fullQuarter.MulVec(p)))
}

func TestShearMoveXInProportionToY(t *testing.T) {
	type testData struct {tf Mat4; res Vec4}
	p := Point4(2, 3, 4)
	cases := []testData{
		{Shear(1, 0, 0, 0, 0, 0), Point4(5, 3, 4)},
		{Shear(0, 1, 0, 0, 0, 0), Point4(6, 3, 4)},
		{Shear(0, 0, 1, 0, 0, 0), Point4(2, 5, 4)},
		{Shear(0, 0, 0, 1, 0, 0), Point4(2, 7, 4)},
		{Shear(0, 0, 0, 0, 1, 0), Point4(2, 3, 6)},
		{Shear(0, 0, 0, 0, 0, 1), Point4(2, 3, 7)},
	}
	for _, c := range cases {
		assert.Equal(t, c.res, c.tf.MulVec(p))
	}
}

func TestTransformationsAreAppliedSequentially(t *testing.T) {
	a := RotX(math.Pi/2.0)
	b := Scale(Vec3{5, 5, 5})
	c := Trans(Vec3{10, 5, 7})
	p := Point4(1, 0, 1)
	p2 := a.MulVec(p)
	p3 := b.MulVec(p2)
	p4 := c.MulVec(p3)

	assert := assert.New(t)
	assert.True(p2.ApproxEqual(Point4(1, -1, 0)))
	assert.True(p3.ApproxEqual(Point4(5, -5, 0)))
	assert.True(p4.ApproxEqual(Point4(15, 0, 7)))
}

func TestChainedTransformationsAppiedReverseOrder(t *testing.T) {
	tf := Trans(Vec3{10, 5, 7}).Mul(Scale(Vec3{5, 5, 5})).Mul(RotX(math.Pi/2.0))
	assert.True(t, tf.MulVec(Point4(1, 0, 1)).ApproxEqual(Point4(15, 0, 7)))
}
