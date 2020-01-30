package game

import "github.com/gen2brain/raylib-go/raylib"

type Snake struct {
	Head Point
	Body PointGroup
	Dir int32
}

func NewSnake() Snake {
	s := Snake{}
	x := rl.GetRandomValue(5, GameWidth - 4)
	y := rl.GetRandomValue(5, GameHeight - 4)
	s.Head = Point{x, y}
	s.Body = NewPointGroup()
	s.Dir = rl.GetRandomValue(0, 3)
	return s
}

func (s *Snake) Move() {
	s.Body.Add(s.Head)
	x, y := s.Head.X, s.Head.Y

	switch s.Dir {

	case DirUp:
	y--
	if y < 0 {
		y = GameHeight - 1
	}

	case DirDown: 
	y++
	if y > GameHeight - 1 {
		y = 0
	}

	case DirLeft:
	x--
	if x < 0 {
		x = GameWidth - 1
	}

	case DirRight:
	x++
	if x > GameWidth - 1 {
		x = 0
	}

	}

	s.Head = Point{x, y}
}


func (s *Snake) PopTail() {
	s.Body.Pop()
}

func (s Snake) Eat(fruits PointGroup) bool {
	return fruits.Remove(s.Head) 
}

func (s *Snake) Turn() {
	if rl.IsKeyPressed(rl.KeyRight) {
		if s.Dir != DirLeft {
			s.Dir = DirRight
		}
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		if s.Dir != DirRight {
			s.Dir = DirLeft
		}
	} else if rl.IsKeyPressed(rl.KeyUp) {
		if s.Dir != DirDown {
			s.Dir = DirUp
		}
	} else if rl.IsKeyPressed(rl.KeyDown) {
		if s.Dir != DirUp {
			s.Dir = DirDown
		}
	} 
}

func (s *Snake) Act(fruits PointGroup) {
	s.Turn()
	s.Move()
	if !s.Eat(fruits){
		s.PopTail()
	}
}


func (s Snake) Dead() bool {
	return s.Body.Contains(s.Head)
}
