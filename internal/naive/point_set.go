package naive

import "github.com/kouheiszk/kdtree/pkg/point"

func NewPointSet(points []point.Point) *PointSet {
	return &PointSet{
		points: points,
	}
}

type PointSet struct {
	points []point.Point
}

func (ps *PointSet) Nearest(goal point.Point) point.Point {
	var nearest point.Point
	minDistance := 0.0

	for _, p := range ps.points {
		distance := goal.Distance(p)
		if nearest == nil || distance < minDistance {
			nearest = p
			minDistance = distance
		}
	}

	return nearest
}
