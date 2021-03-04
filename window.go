package visual1

import (
	"image"
	"sync"
)

type Box struct {
	mu     sync.RWMutex
	pos    image.Point
	bounds image.Rectangle
}

func (s *Box) Pos() image.Point {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.pos
}

func (s *Box) SetPos(pos image.Point) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.pos = pos
}

func (s *Box) Rect() image.Rectangle {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.bounds.Add(s.pos)
}

func (s *Box) SetBounds(b image.Rectangle) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.bounds = b
}

func (s *Box) Bounds() image.Rectangle {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.bounds
}

func NewBox(pos image.Point, w, h int) *Box {
	return &Box{pos: pos, bounds: image.Rect(0, 0, w, h)}
}
