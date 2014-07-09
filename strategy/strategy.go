package strategy

type Calc interface {
	Execute (int, int)
}

type CalcAdd struct {
}

func (CalcAdd) Execute(x int, y int) int {
	return x+y
}

type CalcMult struct {
}

func (CalcMult) Execute(x int, y int) int {
	return x*y
}
