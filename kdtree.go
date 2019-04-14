package kdtree

import (
	"github.com/kouheiszk/kdtree/pkg/point"
)

func NewKDTree(points []point.Point) *KDTree {
	if len(points) == 0 {
		return nil
	}

	tree := &KDTree{
		dim: points[0].Dim(),
	}
	for _, point := range points {
		tree.root = insert(tree.root, point, 0)
	}

	return tree
}

type KDTree struct {
	root *node
	dim  int
}

func (t *KDTree) Dim() int {
	return t.dim
}

func (t *KDTree) Nearest(goal point.Point) point.Point {
	return nearest(t.root, goal, t.root.point)
}

type node struct {
	point point.Point
	axis  int
	left  *node
	right *node
}

func insert(n *node, p point.Point, depth int) *node {
	axis := depth % p.Dim()

	if n == nil {
		return &node{
			point: p,
			axis:  axis,
		}
	}

	if p.Equals(n.point) {
		return n
	}

	// 新規に挿入するデータが小さい場合は左の枝に追加
	cmp := comparePoints(n.point, p, axis)
	if cmp < 0 {
		n.left = insert(n.left, p, depth+1)
	} else {
		n.right = insert(n.right, p, depth+1)
	}
	return n
}

func nearest(n *node, goal point.Point, best point.Point) point.Point {
	if n == nil {
		return best
	}

	minDistance := goal.Distance(best)
	distance := goal.Distance(n.point)
	if distance < minDistance {
		best = n.point
	}

	cmp := comparePoints(n.point, goal, n.axis)
	var goodSide *node
	var badSide *node
	if cmp < 0 {
		goodSide = n.left
		badSide = n.right
	} else {
		goodSide = n.right
		badSide = n.left
	}

	best = nearest(goodSide, goal, best)
	if goal.Distance(possibleBest(n, goal)) <= minDistance {
		best = nearest(badSide, goal, best)
	}

	return best
}

func comparePoints(p1 point.Point, p2 point.Point, axis int) float64 {
	return p1.GetValue(axis) - p2.GetValue(axis)
}

func possibleBest(n *node, goal point.Point) point.Point {
	vec := make([]float64, goal.Dim())
	for i := 0; i < goal.Dim(); i++ {
		if i == n.axis {
			vec[i] = n.point.GetValue(i)
		} else {
			vec[i] = goal.GetValue(i)
		}
	}

	return &point.PointBase{
		Vec: vec,
	}
}
