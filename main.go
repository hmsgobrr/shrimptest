package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 450
	worldWidth   = 800
	worldHeight  = 450
)

// func loadPNG(filePath string) *image.Image {
// 	imgFile, err := os.Open(filePath)
// 	if err != nil {
// 		log.Println("Cannot read PNG file:", err)
// 	}
// 	defer func() {
// 		if err := imgFile.Close(); err != nil {
// 			log.Println("Cannot close PNG file:", err)
// 		}
// 	}()

// 	img, err := png.Decode(imgFile)
// 	if err != nil {
// 		log.Println("Cannot decode PNG file:", err)
// 	}

// 	return &img
// }

type Game struct {
	player *Player
}

func (g *Game) Init() error {
	g.player.Init(&Vec2{0, 0})

	return nil
}

func (g *Game) Update() error {
	g.player.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{30, 30, 30, 255})
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return worldWidth, worldHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("ShrimpTest")
	game := &Game{}
	newPlayer, err := newPlayer("assets/player.png")
	if err != nil {
		panic(err)
	}
	game.player = newPlayer
	if err := game.Init(); err != nil {
		panic(err)
	}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
