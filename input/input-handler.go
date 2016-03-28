package input

type (
	InputHandler interface {
		WaitForInput() int
	}
)
