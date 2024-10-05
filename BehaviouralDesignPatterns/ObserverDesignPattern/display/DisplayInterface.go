package display

type IDisplay interface {
	GetName() string
	Update(temp int)
}
