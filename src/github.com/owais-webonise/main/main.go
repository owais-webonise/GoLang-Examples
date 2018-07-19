package main
import "fmt"
import "github.com/owais-webonise/value"
import "github.com/owais-webonise/forLoop"

func main() {
    a := &value.Example{Num1: 10 , Num2: 20}
    b := a.CalculateSum()
    c := &forLoop.ForExample{I: 1}
    d := c.CalculateFor()
    fmt.Println("Value",b)
    fmt.Println("ForSum",d)
}
