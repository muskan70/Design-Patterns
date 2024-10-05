package display

import "fmt"

type TVDisplay struct {
	name string
}

func NewTVDisplay() IDisplay {
	return &TVDisplay{name: "TV"}
}

func (tv *TVDisplay) GetName() string {
	return tv.name
}

func (tv *TVDisplay) Update(temp int) {
	fmt.Println(tv.name, " Updated Temperature : ", temp)
}
