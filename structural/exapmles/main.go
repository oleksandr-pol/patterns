package main

import (
	"fmt"

	"../../structural"
)

func Sqrt(x float64) float64 {
	z := x / 2
	curr := z

	for {
		curr -= (z*z - x) / (2 * z)
		fmt.Println(z)

		if curr == z {
			return z
		}

		z = curr
	}

	return z
}

func main() {
	// proxy example
	fmt.Println("---proxy example---")
	proxy := structural.NewLocalHostProxy()
	res, _ := proxy.Get("localhost/user")
	fmt.Println(res)

	// decorator example
	fmt.Println("---decorator example---")
	fn := structural.LogDecorate(Sqrt)
	fn(4)
}
