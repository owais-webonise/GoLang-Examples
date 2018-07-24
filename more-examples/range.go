package main
import "fmt"
func main() {
	a := map[string]string{"a": "hello", "b": "world"}
	for b,c := range a {
		fmt.Printf("%s -> %s\n", b,c)
	}
}
