package main
import "fmt"
func main() {
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	_, prs := n["foo"]
    fmt.Println("prs:", prs)
}	
