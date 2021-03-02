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
	pt image.Point

	//kind animKindType

	curFrameAnimDir animDirType
	curXDelta       float64
}

func (s *Animator1) NextFrame() {
	const moveDeltaX = 10

	var accel float64 = float64(moveDeltaX) / 60 * 40

	if math.Abs(s.curXDelta) >= moveDeltaX {
		switch s.curFrameAnimDir {
		case normal:
			s.curFrameAnimDir = reverse
		case reverse:
			s.curFrameAnimDir = normal
		}
	}

	switch s.curFrameAnimDir {
	case normal:
		s.curXDelta += accel
	case reverse:
		s.curXDelta -= accel
	}
}

func (s *Animator1) Apply(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(s.curXDelta, 0)
}
