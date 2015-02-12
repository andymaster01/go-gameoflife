package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {

	z := float64(3)
	last := 0.0

	for z-last > 0.00000001 {
		//for v := 0; v < 10; v++ {
		last = z

		z = z - (((z * z) - x) / (2 * z))
		fmt.Println(z - last)
		fmt.Println(z)
	}

	return z

}

func main() {
	x := Sqrt(2)
	fmt.Println("---")
	fmt.Println(x)
}
