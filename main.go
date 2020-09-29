package main

import "fmt"

func main() {
	var e float64
	var nFactorial int
	var limite int
	e = 0
	fmt.Scanln(&limite)

	for n := 0; n < limite; n++ {
		nFactorial = n
		if n > 1 {
			for i := n; i > 1; i-- {
				nFactorial = nFactorial * (i - 1)
			}
			e = e + (1 / float64(nFactorial))
		} else {
			e = e + (1 / 1)
		}
	}
	fmt.Println(e)
}
