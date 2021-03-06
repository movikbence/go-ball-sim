package collision

import (
	"github.com/Aoana/go-ball-sim/assets/images"
	"github.com/Aoana/go-ball-sim/internal/pkg/ball"
	"github.com/Aoana/go-ball-sim/pkg/gfxutil"
	"github.com/Aoana/go-ball-sim/pkg/mathutil"
	"github.com/hajimehoshi/ebiten"
	"math"
)

// Simulation variables
var (
	// Timestep
	dt = 10.0
	// Images
	backgroundImage *ebiten.Image
	SoccerBallImage *ebiten.Image
	// Screen Resolution
	ScreenWidth  = 1600
	ScreenHeight = 900
	// Starting velocity for balls
	MinV0 float64 = -50
	MaxV0 float64 = 50
)

func init() {

	// Pre-Load background image
	backgroundImage, _ = gfxutil.LoadPngSlice(images.ImageSoccerField)
	SoccerBallImage, _ = gfxutil.LoadPngSlice(images.ImageSoccerBall)
}

// StartValues set starting position and velocity for a slice of balls
// Positions spread in a square and velocity is random
func StartValues(nballs int) error {

	s := int(math.Round(1 + math.Sqrt(float64(nballs))))
	cx := make([]float64, 0)
	cy := make([]float64, 0)

	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			cx = append(cx, float64(200+i*20))
			cy = append(cy, float64(200+j*20))
		}
	}

	for i := 0; i < nballs; i++ {
		// Random starting velocity
		vx0, _ := mathutil.RandInRange(MinV0, MaxV0)
		vy0, _ := mathutil.RandInRange(MinV0, MaxV0)
		// Ball constructor
		b, err := ball.New(cx[i], cy[i], vx0, vy0, 0.07, SoccerBallImage)
		if err != nil {
			return err
		}
		err = ball.Add(b)
		if err != nil {
			return err
		}
	}
	return nil
}

// DrawScenery is a helper function to draw background
func DrawScenery(screen *ebiten.Image) {
	// Draw background
	gfxutil.PrintImage(screen, backgroundImage, 0, 0, 4.0, 3.5)
}

// Timestep is a helper function to perform a timestep with position and velocity updates
func Timestep(b *ball.Ball) {

	b.Obj.Position(dt)
}

// Edge checks if ball reached the boundary
func Edge(b *ball.Ball) {

	// Skip check if ball marked as inactive
	if !b.Active {
		return
	}
	// Bounce on boundary
	ball.Boundary(b, 0, float64(ScreenWidth), 0, float64(ScreenHeight), float64(1))
}

// Goal checks if ball reached a goal
func Goal(b *ball.Ball) {

	// Check if reached goal (left and right)
	if b.Obj.X[0] < b.Radius && b.Obj.V[0] < 0 {
		if b.Obj.X[1] < 640 && b.Obj.X[1] > 340 {
			b.Active = false
		}
	}
	if b.Obj.X[0] > float64(ScreenWidth)-b.Radius && b.Obj.V[0] > 0 {
		if b.Obj.X[1] < 630 && b.Obj.X[1] > 330 {
			b.Active = false
		}
	}
}
