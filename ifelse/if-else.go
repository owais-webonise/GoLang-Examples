package ifelse

type IfElseExample struct {
	N int
}

type IfElseDao interface {
	CalculateIfElse(int) string
}

func (daoImpl *IfElseExample) CalculateIfElse() string {
	if daoImpl.N % 2 == 0 {
		return "N is even"
	} else {
		return "N is odd"
	}
}
