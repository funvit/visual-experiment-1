package visual1

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math"
)

const circleAnimatorMaxAngle = 360.0

// Animates game entity moving it using closed circle path.
type CircleAnimator1 struct {
	pt     image.Point
	radius float64

	curAngle       float64 // [0, 360)
	speed          float64
	x, y           float64
	deltaX, deltaY float64
}

func NewCircleAnimator1(basePos image.Point, radius float64, animDurFrames int) *CircleAnimator1 {
	return &CircleAnimator1{
		pt:     basePos,
		radius: radius,
		speed:  circleAnimatorMaxAngle / float64(animDurFrames),
	}
}

func (s *CircleAnimator1) NextFrame() {
	if s.curAngle == 0 {
		s.deltaX, s.deltaY = s.calcCirclePathPos(s.curAngle, s.radius)
	}

	if s.curAngle >= 360 {
		s.curAngle = s.speed
	} else {
		s.curAngle += s.speed
	}

	s.x, s.y = s.calcCirclePathPos(s.curAngle, s.radius)
	s.x -= s.deltaX
	s.y -= s.deltaY
}

func (s *CircleAnimator1) calcCirclePathPos(angle, radius float64) (x, y float64) {
	x = math.Cos(s.curAngle*math.Pi/180) * s.radius
	y = -math.Sin(s.curAngle*math.Pi/180) * s.radius
	return x, y
}

func (s *CircleAnimator1) Apply(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(s.x, s.y)
}
