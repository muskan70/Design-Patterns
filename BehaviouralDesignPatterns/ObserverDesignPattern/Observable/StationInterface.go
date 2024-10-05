package observable

import (
	"observerPattern/display"
)

type IStation interface {
	AddDisplay(d display.IDisplay)
	RemoveDisplay(displayName string)
	GetUpdates() int
	SetUpdates(update int)
	NotifyDisplay()
}
