package main

import tl "github.com/JoelOtter/termloop"

const (
	starChar = '*'
)

// func (player *Player) Size() (int, int) {
// 	return player.Size()
// }

// func (player *Player) Position() (int, int) {
// 	return player.Position()
// }

// func (player *Player) Collide(collision tl.Physical) {
// 	// Check if it's a Rectangle we're colliding with
// 	if _, ok := collision.(*tl.Rectangle); ok {
// 		player.SetPosition(player.prevX, player.prevY)
// 	}
// }
type Game struct {
	status *tl.Text
}

func NewGame(game *tl.Game) {

}

func main() {
	game := tl.NewGame()
	// game.Screen()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorWhite,
	})
	status := tl.NewText(19, 0, "Yippie", tl.ColorWhite, tl.ColorBlack)
	level.AddEntity(status)
	// screenWidth, screenHeight := game.Screen().Size()
	// x, y := player.Position()
	// player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	// level.AddEntity(tl.NewRectangle(0, screenHeight-10, 80, 20, tl.ColorGreen))
	player := Player{Entity: tl.NewEntity(10, 40, 1, 1)}
	// Set the character at position (0, 0) on the entity.
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorWhite, Ch: '🛒'})
	level.AddEntity(&player)
	star := NewStar(game)
	level.AddEntity(star)
	game.Screen().SetLevel(level)
	game.Start()
}
