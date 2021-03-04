package visual1

import (
	"context"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/exp/rand"
	"image"
	"image/color"
	"log"
	"os"
	"sync/atomic"
	"time"
)

var (
	logInfo = log.New(os.Stdout, "INFO ", log.LstdFlags)
	logErr  = log.New(os.Stdout, "ERROR ", log.LstdFlags|log.Lshortfile)
)

const (
	maxBoxes = 40 //100
	loopFps  = 60
)

type Game struct {
	window *Box

	loopFrame      int
	gameLoopCtx    context.Context
	gameLoopCancel context.CancelFunc

	boxes []*EBox1

	//tmpImg    *ebiten.Image
	debugMode bool

	beatRate  int32
	beatRateC chan int32
}

// EBox1 = entity box 1
type EBox1 struct {
	Pt   *image.Point
	img  *ebiten.Image
	Anim *CircleAnimator1
}

func init() {
	rand.Seed(uint64(time.Now().Nanosecond()))
}

func (g *Game) Layout(int, int) (int, int) {
	return g.window.bounds.Dx(), g.window.bounds.Dy()
}

func (g *Game) Update() error {
	// play animation
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		if g.gameLoopCancel != nil {
			g.gameLoopCancel()
			g.gameLoopCancel = nil
		} else {
			g.gameLoopCtx, g.gameLoopCancel = context.WithCancel(context.Background())
			go g.loop(g.gameLoopCtx)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		if g.gameLoopCancel != nil {
			g.gameLoopCancel()
			g.gameLoopCancel = nil
		}
		g.loopOneFrame()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.debugMode = !g.debugMode
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		if atomic.LoadInt32(&g.beatRate) < 300 {
			atomic.AddInt32(&g.beatRate, 1)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if atomic.LoadInt32(&g.beatRate) > 0 {
			atomic.AddInt32(&g.beatRate, -1)
		}
	}

	return nil
}

// Draw draws game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	for i := range g.boxes {
		b := g.boxes[i]
		op := &ebiten.DrawImageOptions{}
		b.Anim.Apply(op)
		op.GeoM.Translate(float64(b.Pt.X), float64(b.Pt.Y))

		screen.DrawImage(b.img, op)
	}

	// must be last
	g.drawDebug(screen)
}

func (g *Game) Run() error {
	logInfo.Printf("window size is %dx%d",
		g.window.bounds.Dx(), g.window.bounds.Dy())

	ebiten.SetWindowSize(g.window.bounds.Dx(), g.window.bounds.Dy())
	ebiten.SetWindowTitle("Visual experiment No 1")

	return ebiten.RunGame(g)
}

func (g *Game) loop(ctx context.Context) {
	logInfo.Println("game loop started")
	t := time.NewTicker(1 * time.Second / loopFps)

	for {
		select {

		case <-ctx.Done():
			logInfo.Println("game loop stopped")
			return

		case <-t.C:
			g.loopOneFrame()
		}
	}
}

func (g *Game) loopOneFrame() {
	g.loopFrame++
	if g.loopFrame > loopFps {
		g.loopFrame = 1
	}

	var boost int32
	select {
	case v := <-g.beatRateC:
		boost = v
	default:
	}

	for i := range g.boxes {
		b := g.boxes[i]

		if boost > 0 {
			b.Anim.SetBoost(boost)
		}

		b.Anim.NextFrame()
	}
}

func (g *Game) drawDebug(screen *ebiten.Image) {
	if !g.debugMode {
		return
	}
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("Loop frame: %d", g.loopFrame),
		0,
		0,
	)
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("Beat rate: %d", atomic.LoadInt32(&g.beatRate)),
		0,
		14,
	)
}

func New(w, h int) *Game {
	g := &Game{
		window:    NewBox(image.Pt(0, 0), w, h),
		beatRate:  30,
		beatRateC: make(chan int32),
	}

	g.boxes = make([]*EBox1, 0, maxBoxes)
	for i := 0; i < maxBoxes; i++ {
		x := 40 + rand.Intn(g.window.Bounds().Max.X-16-40)
		y := 40 + rand.Intn(g.window.Bounds().Max.Y-16-40)

		b := &EBox1{
			img: ebiten.NewImage(16, 16),
			Pt: &image.Point{
				X: x,
				Y: y,
			},
		}
		b.img.Fill(color.RGBA{
			R: uint8(rand.Int31n(0xff/3*2)) + 0xff/3,
			G: uint8(rand.Int31n(0xff/3*2)) + 0xff/3,
			B: uint8(rand.Int31n(0xff/3*2)) + 0xff/3,
			A: 0xff,
		})
		b.Anim = NewCircleAnimator1(
			*b.Pt,
			20+rand.Float64()+20,
			rand.Intn(loopFps*4)+loopFps,
		)

		g.boxes = append(g.boxes, b)
	}

	go func() {
		for {
			select {
			case <-time.Tick(10_000 * time.Millisecond / time.Duration(g.beatRate)):
				g.beatRateC <- 100
			}
		}
	}()

	return g
}
