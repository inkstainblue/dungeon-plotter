package input

type (
	InputHandler interface {
		WaitForInput() (code int, label string)
	}
)
