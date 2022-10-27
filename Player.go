package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CharacterState int

const (
	Idle CharacterState = iota
	Walk
	Jump
)

type Player struct {
	scale float64
	po    *PhysicsObject
	image *ebiten.Image
}

func newPlayer(texPath string) (*Player, error) {
	p := &Player{}

	image, _, err := ebitenutil.NewImageFromFile(texPath)
	if err != nil {
		return nil, err
	}
	p.image = image
	p.scale = 1

	return p, nil
}

func (p *Player) Init(pos *Vec2) {
	p.po.pos = pos
	p.po.vel.x = 0
	p.po.vel.y = 0
}

func (p *Player) Update() {

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.po.pos.x, p.po.pos.y)
	op.GeoM.Scale(p.scale, p.scale)
	screen.DrawImage(p.image, op)
}
