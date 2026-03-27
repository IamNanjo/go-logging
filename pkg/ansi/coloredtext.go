package ansi

type ColoredText struct {
	Color Color
	Text  string
}

func (ct *ColoredText) String() string {
	if ct.Color == None {
		return ct.Text
	}
	return string(ct.Color) + ct.Text + string(Reset)
}
