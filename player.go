package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	*tl.Entity
	prevX int
	prevY int
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		player.prevX, player.prevY = player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			if player.prevX < 30 {
				player.SetPosition(player.prevX+1, player.prevY)
			}
		case tl.KeyArrowLeft:
			if player.prevX > 0 {
				player.SetPosition(player.prevX-1, player.prevY)
			}
		}
	}
}

func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*Star); ok {
		player.SetPosition(10, 1)
	}
}
