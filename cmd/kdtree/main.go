package main

import (
	"fmt"
	"github.com/kouheiszk/kdtree"
	"github.com/kouheiszk/kdtree/internal/naive"
	"github.com/kouheiszk/kdtree/pkg/point"
	"math/rand"
	"time"
)

func main() {
	pointsCount := 10000
	queryCount := 2000

	points := randomPoints(pointsCount)
	ps := naive.NewPointSet(points)
	tree := kdtree.NewKDTree(points)
	queries := randomPoints(queryCount)

	t0 := time.Now()
	for _, q := range queries {
		ps.Nearest(q)
	}
	t1 := time.Now()
	for _, q := range queries {
		tree.Nearest(q)
	}
	t2 := time.Now()

	fmt.Println("naive point set query duration:", t1.Sub(t0))
	fmt.Println("kdtree query duration:", t2.Sub(t1))

	for _, q := range queries {
		expected := ps.Nearest(q)
		actual := tree.Nearest(q)
		if !expected.Equals(actual) {
			fmt.Println("invalid implementation!!")
		}
	}
}

func randomPoints(n int) []point.Point {
	points := make([]point.Point, n)
	for i := 0; i < n; i++ {
		points[i] = random2DPoint()
	}
	return points
}

func random2DPoint() point.Point {
	return point.PointBase{
		Vec: []float64{rand.Float64(), rand.Float64()},
	}
}

//func random3DPoint() point.Point {
//	return point.PointBase{
//		Vec: []float64{rand.Float64(), rand.Float64(), rand.Float64()},
//	}
//}

//func random9DPoint() point.Point {
//	return point.PointBase{
//		Vec: []float64{
//			rand.Float64(),
//			rand.Float64(),
//			rand.Float64(),
//			rand.Float64(),
//			rand.Float64(),
//			rand.Float64(),
//			rand.Float64(),
//			rand.Float64(),
//			rand.Float64(),
//		},
//	}
//}
