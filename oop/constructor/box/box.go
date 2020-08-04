package box

import "fmt"

func NewBox(w, h int) *Box {
	return &Box{
		Width: w,
		Hight: h,
	}
}

type Box struct {
	Width int
	Hight int
}

func (b *Box) String() string {
	return fmt.Sprintf("box : width %v , hight %v", b.Width, b.Hight)
}
