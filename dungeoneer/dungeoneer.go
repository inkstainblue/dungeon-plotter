package dungeoneer

import (
	"sync"

	"github.com/inkstainblue/dungeon-plotter/canvas"
	"github.com/inkstainblue/dungeon-plotter/controller"
	"github.com/inkstainblue/dungeon-plotter/input"
)

type (
	Dungeoneer struct {
		ctrl controller.Controller

		moveLock sync.Mutex
		pos      canvas.Point
	}

	DungeoneerOptions struct{}
)

const (
	maxGridWidth  = 100
	maxGridHeight = 100
)

// New creates a new dungeoneer using the given options.
func New(opts DungeoneerOptions) (d Dungeoneer) {
	// FIXME: Use flags in the DungeoneerOptions to choose what to initialize.
	sc := canvas.NewScreenCanvas(maxGridWidth, maxGridHeight)
	k := input.NewKeyboard()

	d.ctrl = controller.New([]canvas.Canvas{&sc}, []input.InputManager{&k})

	d.pos = canvas.Pt(50, 50)

	return
}

// Start makes the dungeoneer start processing input.
func (d *Dungeoneer) Start() {
	// FIXME: Remove this temporary world outline.
	d.ctrl.DrawWall(canvas.Pt(0, 0), canvas.Pt(0, 99))
	d.ctrl.DrawWall(canvas.Pt(0, 0), canvas.Pt(99, 0))
	d.ctrl.DrawWall(canvas.Pt(99, 0), canvas.Pt(99, 99))
	d.ctrl.DrawWall(canvas.Pt(0, 99), canvas.Pt(99, 99))

	// Mark the starting position.
	d.ctrl.DrawPath(d.pos, d.pos)

	d.ctrl.BindInput(111, d.moveUp)    // Up
	d.ctrl.BindInput(116, d.moveDown)  // Down
	d.ctrl.BindInput(113, d.moveLeft)  // Left
	d.ctrl.BindInput(114, d.moveRight) // Right
}

// WaitForQuit blocks until the dungeoneer has exited.
func (d *Dungeoneer) WaitForQuit() {
	d.ctrl.WaitForQuit()
}

func (d *Dungeoneer) move(delta canvas.Point) {
	d.moveLock.Lock()
	defer d.moveLock.Unlock()

	// FIXME: Don't allow values outside of the grid.
	// FIXME: Don't allow values that intersect with the walls.
	dest := d.pos.Add(delta)

	d.ctrl.DrawPath(d.pos, dest)

	d.pos = dest
}

func (d *Dungeoneer) moveUp() {
	d.move(canvas.Pt(0, -1))
}

func (d *Dungeoneer) moveDown() {
	d.move(canvas.Pt(0, 1))
}

func (d *Dungeoneer) moveLeft() {
	d.move(canvas.Pt(-1, 0))
}

func (d *Dungeoneer) moveRight() {
	d.move(canvas.Pt(1, 0))
}
