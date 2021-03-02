package visual1

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math"
)

type CircleAnimator1 struct {
	pt     image.Point
	radius float64

	accel float64

	curXDelta float64
	curYDelta float64

	targetX, targetY float64
}

func NewCircleAnimator1(basePos image.Point, radius float64) *CircleAnimator1 {
	return &CircleAnimator1{
		pt:     basePos,
		radius: radius,
		accel:  radius / 60 * 40,
	}
}

func (s *CircleAnimator1) NextFrame() {
	// fixme: rewrite

	s.curXDelta = math.Sin()

}

func (s *CircleAnimator1) Apply(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(s.curXDelta, s.curYDelta)
}
