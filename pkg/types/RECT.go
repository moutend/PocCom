package types

type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

func (r RECT) IsZero() bool {
	return r.Left == 0 && r.Top == 0 && r.Right == 0 && r.Bottom == 0
}
