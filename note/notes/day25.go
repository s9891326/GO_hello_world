package note

import "fmt"

// 5.2 接口
type textMes struct {
	Text string
	Type string
}

type imgMes struct {
	Img  string
	Type string
}

func (tm *textMes) setText() {
	tm.Text = "hello"
}

func (tm *textMes) setType() {
	tm.Type = "text"
}

func (im *imgMes) setImg() {
	im.Img = "image"
}

func (im *imgMes) setType() {
	im.Type = "im"
}

type Mes interface {
	setType()
}

func SendMes(m Mes) {
	m.setType()
	switch mptr := m.(type) {
	case *textMes:
		mptr.setText()
	case *imgMes:
		mptr.setImg()
	}
	fmt.Println("m=", m)
}

func main25() {
	tm := textMes{}
	SendMes(&tm)
	fmt.Println("tm=", tm)
	im := imgMes{}
	SendMes(&im)

	var n1 int = 1
	n1Interface := interface{}(n1)
	n2, ok := n1Interface.(int)
	if ok {
		fmt.Println("n2=", n2)
	} else {
		fmt.Println("error")
	}
}
