package forLoop

type ForExample struct {
	I int
}

type ForDao interface {
	CalculateFor(int) int
}

func (daoImpl *ForExample) CalculateFor() int {
	var sum int
	for n := daoImpl.I; n <=3 ; n++ {
		sum = sum + n
	}
	return sum;
}
