package display

import "fmt"

type MobileDisplay struct {
	name string
}

func NewMobileDisplay() IDisplay {
	return &MobileDisplay{name: "Mobile"}
}

func (mb *MobileDisplay) GetName() string {
	return mb.name
}

func (mb *MobileDisplay) Update(temp int) {
	fmt.Println(mb.name, " Updated Temperature : ", temp)
}
