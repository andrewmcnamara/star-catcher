package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Star struct {
	*tl.Entity
	game  *tl.Game
	speed int
	frame int
	slide int
}

func NewStar(game *tl.Game) *Star {
	star := Star{
		game:   game,
		slide:  1,
		Entity: tl.NewEntity(1, 1, 1, 1),
	}

	star.Entity.SetCell(0, 0, &tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorYellow, Ch: starChar})
	return &star
}

func (star *Star) Draw(screen *tl.Screen) {
	star.speed = int(star.game.Screen().TimeDelta() * float64(30000000/1))

	if star.frame > star.speed {
		star.slide = star.slide * -1
		prevX, prevY := star.Position()
		if prevY == 10 {
			screen.Level().RemoveEntity(star)

		} else {
			star.Entity.SetPosition(prevX+star.slide, prevY+1)
			star.frame = 0
		}
	}
	star.frame++

	star.Entity.Draw(screen)
}

func (star *Star) Tick(event tl.Event) {
}
