package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	tmp := -1.0
	z := x
	for i := 0; i < 10; i++ {
		tmp = z
		z -= (z*z - x) / (2*z)
		if z == tmp {
			fmt.Println(i)
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(1))
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(5))
}
