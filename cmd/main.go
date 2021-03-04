package main

import (
	"flag"
	"fmt"
	"visual1"
)

var (
	screenW = flag.Int("width", 600, "window width")
	screenH = flag.Int("height", 600, "window height")
)

func main() {
	fmt.Println("Visual experiment 1 beta")
	flag.Parse()

	g := visual1.New(*screenW, *screenH)
	err := g.Run()
	if err != nil {
		panic(fmt.Errorf("engine init error: %w", err))
	}
}
