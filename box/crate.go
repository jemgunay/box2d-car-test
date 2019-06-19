package box

import (
	"github.com/ByteArena/box2d"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Crate struct {
	bodyDef *box2d.B2BodyDef
	body    *box2d.B2Body

	width, length float64
}

func NewCrate(w *box2d.B2World, x, y, width, length float64) *Crate {
	// create rigid body definition
	bodyDef := box2d.NewB2BodyDef()
	bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
	bodyDef.Position = box2d.MakeB2Vec2(x/box2dScale, y/box2dScale)
	bodyDef.Angle = 0

	// create fixture shape
	shape := box2d.NewB2PolygonShape()
	shape.SetAsBox(width/2.0/box2dScale, length/2.0/box2dScale)

	// create fixture
	fixDef := box2d.MakeB2FixtureDef()
	fixDef.Shape = shape
	fixDef.Density = 1.0
	fixDef.Friction = 0.3
	fixDef.Restitution = 0.4

	// create body
	body := w.CreateBody(bodyDef)
	body.CreateFixtureFromDef(&fixDef)

	return &Crate{
		bodyDef: bodyDef,
		body:    body,

		width:  width,
		length: length,
	}
}

func (w *Crate) Draw(win *pixelgl.Window) {
	pos := box2dToPixel(w.body.GetPosition())
	rot := w.body.GetAngle()

	x := (pos.X * box2dScale) - w.width/2.0
	y := (pos.Y * box2dScale) - w.length/2.0
	width := w.width
	length := w.length

	carBodySprite := imdraw.New(nil)
	carBodySprite.Color = pixel.RGB(0.8, 0.8, 0.8)
	carBodySprite.Push(
		pixel.V(x, y),
		pixel.V(x, y+length),
		pixel.V(x+width, y+length),
		pixel.V(x+width, y),
	)

	carBodySprite.SetMatrix(pixel.IM.Rotated(pixel.V(x+(width/2.0), y+(length/2.0)), rot))

	carBodySprite.Polygon(0)
	carBodySprite.Draw(win)
}