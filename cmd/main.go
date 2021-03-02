package main

import (
	"flag"
	"fmt"
	"visual1"
)

var (
	screenW = flag.Int("width", 540, "window width")
	screenH = flag.Int("height", 600, "window height")
)

func main() {
	fmt.Println("Audio sequencer beta")
	flag.Parse()

	g := visual1.New(*screenW, *screenH)
	err := g.Run()
	if err != nil {
		panic(fmt.Errorf("engine init error: %w", err))
	}
}
