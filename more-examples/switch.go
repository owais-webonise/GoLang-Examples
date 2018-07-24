package main
import "fmt"

func main() {
i:= 2
fmt.Print("Number is ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")		
	}

	whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("Bool")
        case int:
            fmt.Println("int")
        default:
            fmt.Printf("Invalid %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")	
}
