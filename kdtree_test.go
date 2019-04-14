package kdtree

import (
	"fmt"
	"github.com/kouheiszk/kdtree/pkg/point"
)

func ExampleKDTree_Dim() {
	tree := NewKDTree([]point.Point{
		&point.PointBase{Vec: []float64{1.1, 2.2}},
		&point.PointBase{Vec: []float64{3.3, 4.4}},
		&point.PointBase{Vec: []float64{-2.9, 4.2}},
	})

	fmt.Println(tree.Dim())

	// Output:
	// 2
}

func ExampleKDTree_Nearest() {
	tree := NewKDTree([]point.Point{
		&point.PointBase{Vec: []float64{1.1, 2.2}},
		&point.PointBase{Vec: []float64{3.3, 4.4}},
		&point.PointBase{Vec: []float64{-2.9, 4.2}},
	})

	point := &point.PointBase{Vec: []float64{3.0, 4.0}}
	nearest := tree.Nearest(point)
	fmt.Println(nearest)

	// Output:
	// &{<nil> [3.3 4.4]}
}
