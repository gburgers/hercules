package math

type Numbers struct {
	Age    int
	Number int
}

func Times(c Numbers) int {
	return c.Age * c.Number
}
