package collision

import (
	"github.com/Aoana/ball-sim-go/internal/pkg/ball"
	"github.com/Aoana/ball-sim-go/pkg/gfxutil"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/hajimehoshi/ebiten"
)

// Simulation variables
var (
	// Mathematical values
	dt = 10.0
	// Images
	backgroundImage *ebiten.Image
	SoccerBallImage *ebiten.Image
	// Screen Resolution
	ScreenWidth  = 1600
	ScreenHeight = 900
	// Starting values for balls
	//	X0            = float64(ScreenWidth) / 2
	//	Y0            = float64(ScreenHeight) / 10
	MinV0 float64 = -50
	MaxV0 float64 = 50
)

func init() {

	// Load background image
	backgroundImage, _ = gfxutil.LoadPng("./assets/soccerfield.png")
	SoccerBallImage, _ = gfxutil.LoadPng("./assets/soccerball.png")

}

// DrawScenery is a helper function to draw background
func DrawScenery(screen *ebiten.Image) {
	// Draw background
	gfxutil.PrintImage(screen, backgroundImage, 0, 0, 4.0, 3.5)
}

// Timestep is a helper function to perform a timestep with position and velocity updates
func Timestep(o *objects.Object) {
	o.Position(dt)
}

// OutOfBound is a helper function to set the right boundary
func OutOfBound(b *ball.Ball) {
	ball.Boundary(b, float64(0), float64(ScreenWidth-40), float64(0), float64(ScreenHeight-40), float64(1))
}
