package main
import "fmt"
import "github.com/owais-webonise/value"
import "github.com/owais-webonise/forLoop"
import "github.com/owais-webonise/ifelse"

func main() {
    a := &value.Example{Num1: 10 , Num2: 20}
    b := a.CalculateSum()
    c := &forLoop.ForExample{I: 1}
    d := c.CalculateFor()
    i := &ifelse.IfElseExample{N: 2}
    e := i.CalculateIfElse()
    fmt.Println("Value",b)
    fmt.Println("ForSum",d)
    fmt.Println("If-Else",e)
}
