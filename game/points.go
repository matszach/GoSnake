package game

import l "container/list"


// ================ POINT ================
type Point struct {
	X, Y int32
}

func (p1 Point) Equals(p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}


// ================ POINT LIST ================
type PointGroup struct {
	Points *l.List
}

func (pg PointGroup) Contains(p Point) bool {
	for e := pg.Points.Front(); e != nil; e = e.Next() {
		if e.Value.(Point).Equals(p) {
			return true
		}
	}
	return false
}

func (pg *PointGroup) Add(p Point) int {
	pg.Points.PushFront(p)
	return pg.Points.Len()
} 

func (pg *PointGroup) Pop() Point {
	p := pg.Points.Back()
	pg.Points.Remove(p)
	return p.Value.(Point)
}

func (pg *PointGroup) Size() int {
	return pg.Points.Len()
}

func (pg *PointGroup) Remove(p Point) bool {
	for e := pg.Points.Front(); e != nil; e = e.Next() {
		if e.Value.(Point).Equals(p) {
			pg.Points.Remove(e)
			return true
		}
	}
	return false
}

func NewPointGroup() PointGroup {
	p := PointGroup{}
	p.Points = l.New()
	return p
}





