package value

type Example struct {
    Num1 int
    Num2 int
}

type NumberDao interface {
    CalculateSum(int, int) int
}

func (daoImpl *Example) CalculateSum() int {
    return daoImpl.Num1 + daoImpl.Num2
}
