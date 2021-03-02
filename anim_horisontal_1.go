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

	curFrameAnimDir   animDirType
	curFrameEntityPos image.Point
}

func (s *Animator1) NextFrame() {
	const moveDeltaX = 10

	if math.Abs(float64(s.curFrameEntityPos.X)) >= moveDeltaX {
		switch s.curFrameAnimDir {
		case normal:
			s.curFrameAnimDir = reverse
		case reverse:
			s.curFrameAnimDir = normal
		}
	}

	switch s.curFrameAnimDir {
	case normal:
		s.curFrameEntityPos.X++
	case reverse:
		s.curFrameEntityPos.X--
	}
}

func (s *Animator1) Apply(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(float64(s.curFrameEntityPos.X), float64(s.curFrameEntityPos.Y))
}
