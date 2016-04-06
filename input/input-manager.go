package input

type (
	InputManager interface {
		WaitForInput() (code int, label string)
	}
)
