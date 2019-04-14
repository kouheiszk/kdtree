package point

import "math"

type Point interface {
	Dim() int
	GetValue(dim int) float64
	Distance(point Point) float64
	Equals(point Point) bool
}

type PointBase struct {
	Point
	Vec []float64
}

func (base PointBase) Dim() int {
	return len(base.Vec)
}

func (base PointBase) GetValue(dim int) float64 {
	return base.Vec[dim]
}

func (base PointBase) Distance(point Point) float64 {
	dim := min(base.Dim(), point.Dim())
	sum := 0.0
	for i := 0; i < dim; i++ {
		sum += math.Pow(base.GetValue(i)-point.GetValue(i), 2)
	}
	return math.Sqrt(sum)
}

func (base PointBase) Equals(point Point) bool {
	dim := min(base.Dim(), point.Dim())
	result := true
	for i := 0; i < dim; i++ {
		result = result && base.GetValue(i) == point.GetValue(i)
	}
	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
