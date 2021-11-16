package init

type node struct {
	a int
	b int
}

func NewInit() node {
	return node{
		a: 1,
		b: 2,
	}
}
