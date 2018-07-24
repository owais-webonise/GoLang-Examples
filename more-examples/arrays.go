package main

import "fmt"

func main() {
	var a[5] int
	fmt.Println("emp", a)

	var twoD [2][2]int 
	for i:=0; i < 2; i++ {
		for j:=0; j < 2; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)
}
