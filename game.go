package visual1

import (
	"context"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/exp/rand"
	"image"
	"image/color"
	"log"
	"os"
	"time"
)

var (
	logInfo = log.New(os.Stdout, "INFO ", log.LstdFlags)
	logErr  = log.New(os.Stdout, "ERROR ", log.LstdFlags|log.Lshortfile)
)

const maxBoxes = 100

type Game struct {
	window *Box

	gameLoopCtx    context.Context
	gameLoopCancel context.CancelFunc

	boxes []*EBox1

	tmpImg *ebiten.Image
}

// EBox1 = entity box 1
type EBox1 struct {
	Pt *image.Point

	Anim *Animator1
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

	return nil
}

// Draw draws game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	for i := range g.boxes {
		b := g.boxes[i]
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(b.Pt.X), float64(b.Pt.Y))

		b.Anim.Apply(op)

		screen.DrawImage(g.tmpImg, op)
	}
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
	t := time.NewTicker(1 * time.Second / 20)

	for {
		select {

		case <-ctx.Done():
			logInfo.Println("game loop stopped")
			return

		case <-t.C:

			for i := range g.boxes {
				b := g.boxes[i]
				b.Anim.NextFrame()
			}
		}
	}
}

func New(w, h int) *Game {
	g := &Game{
		window: NewBox(image.Pt(0, 0), w, h),
		tmpImg: ebiten.NewImage(16, 16),
	}

	g.tmpImg.Fill(color.White)

	g.boxes = make([]*EBox1, 0, maxBoxes)
	for i := 0; i < maxBoxes; i++ {
		x := rand.Intn(g.window.Bounds().Max.X - 16)
		y := rand.Intn(g.window.Bounds().Max.Y - 16)

		b := &EBox1{
			Pt: &image.Point{
				X: x,
				Y: y,
			},
		}

		b.Anim = &Animator1{
			pt:                *b.Pt,
			curFrameAnimDir:   normal,
			curFrameEntityPos: image.Point{},
		}

		g.boxes = append(g.boxes, b)
	}

	return g
}
