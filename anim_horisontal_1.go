package visual1

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math"
)

type animDirType uint8

const (
	normal = animDirType(iota + 1)
	reverse
)

type Animator1 struct {
	// anim base point
	pt              image.Point
	curFrameAnimDir animDirType
	moveXDelta      float64

	curXDelta float64
	accel     float64
}

func NewAnimator1(basePos image.Point, moveXDelta float64) *Animator1 {
	return &Animator1{
		pt:              basePos,
		curFrameAnimDir: normal,
		moveXDelta:      moveXDelta,
		accel:           moveXDelta / 60 * 40,
	}
}

func (s *Animator1) NextFrame() {

	if math.Abs(s.curXDelta) >= s.moveXDelta {
		switch s.curFrameAnimDir {
		case normal:
			s.curFrameAnimDir = reverse
		case reverse:
			s.curFrameAnimDir = normal
		}
	}

	switch s.curFrameAnimDir {
	case normal:
		s.curXDelta += s.accel
	case reverse:
		s.curXDelta -= s.accel
	}
}

func (s *Animator1) Apply(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(s.curXDelta, 0)
}
