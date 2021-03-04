package visual1

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math"
	"sync/atomic"
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
	boost          int32
}

func (s *CircleAnimator1) SetBoost(v int32) {
	atomic.StoreInt32(&s.boost, v)
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

	if s.boost > 0 {
		s.boost -= 10
	}
}

func (s *CircleAnimator1) calcCirclePathPos(angle, radius float64) (x, y float64) {
	x = math.Cos(angle*math.Pi/180) * radius
	y = -math.Sin(angle*math.Pi/180) * radius
	return x, y
}

func (s *CircleAnimator1) Apply(op *ebiten.DrawImageOptions) {
	if s.boost > 0 {
		op.GeoM.Translate(-8, -8)
		op.GeoM.Scale(float64(s.boost)/100+1, float64(s.boost)/100+1)
		op.GeoM.Translate(8, 8)
	}

	op.GeoM.Translate(s.x, s.y)
}
