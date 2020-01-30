package game

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	PlayerSnake Snake
	Fruits PointGroup
}

func (g *Game) Init(){
	g.PlayerSnake = NewSnake()
	g.Fruits = NewPointGroup()
	g.SpawnFruit()
}

func (g *Game) SpawnFruit() {
	for g.Fruits.Size() < MinFruit {
		x := rl.GetRandomValue(0, GameWidth - 1)
		y := rl.GetRandomValue(0, GameHeight - 1)
		f := Point{x, y}
		
		if !f.Equals(g.PlayerSnake.Head) && !g.PlayerSnake.Body.Contains(f) {
			g.Fruits.Add(f)
		}
	}
}

func (g Game) GameOver() bool {
	return g.PlayerSnake.Dead()
}

func (g *Game) Logic() {
	g.PlayerSnake.Act(g.Fruits)
	g.SpawnFruit()
}

func (g Game) Draw() {
	rl.ClearBackground(rl.Black)
	
	// grid
	for i := int32(0); i < GameWidth; i++ {
		rl.DrawLineV(
			rl.NewVector2(float32(Unit * i), 0),
			rl.NewVector2(float32(Unit * i), float32(Unit * GameHeight)),
			rl.DarkGray,
		)
	}
	for i := int32(0); i < GameHeight; i++ {
		rl.DrawLineV(
			rl.NewVector2(0, float32(Unit * i)),
			rl.NewVector2(float32(Unit * GameWidth), float32(Unit * i)),
			rl.DarkGray,
		)
	}

	// fruits
	for e := g.Fruits.Points.Front(); e != nil; e = e.Next() {
		p := e.Value.(Point)
		rl.DrawRectangleV(
			rl.NewVector2(float32(p.X * Unit), float32(p.Y * Unit)),
			rl.NewVector2(float32(Unit), float32(Unit)),
			rl.Red,
		)
	}

	// snake body
	for e := g.PlayerSnake.Body.Points.Front(); e != nil; e = e.Next() {
		p := e.Value.(Point)
		rl.DrawRectangleV(
			rl.NewVector2(float32(p.X * Unit), float32(p.Y * Unit)),
			rl.NewVector2(float32(Unit), float32(Unit)),
			rl.Green,
		)
	}

	// snake head
	s := g.PlayerSnake.Head
	rl.DrawRectangleV(
		rl.NewVector2(float32(s.X * Unit), float32(s.Y * Unit)),
		rl.NewVector2(float32(Unit), float32(Unit)),
		rl.DarkGreen,
	)

	
}

