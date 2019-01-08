package unittests

type Math struct {
}

func (m Math) Square(i int) int {
	return i * i
}

func (m Math) BrokenSquare(i int) int {
	return (i * i) + 1
}
