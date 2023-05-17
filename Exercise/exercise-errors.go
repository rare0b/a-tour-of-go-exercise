package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	v := float64(e)
	return fmt.Sprint("cannot Sqrt negative number: ", v)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	
	tmp := -1.0
	z := x
	for i := 0; i < 10; i++ {
		tmp = z
		z -= (z*z - x) / (2*z)
		if z == tmp {
			fmt.Println(i)
			return z, nil
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
