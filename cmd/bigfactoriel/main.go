package main

import (
	"fmt"
	"math/big"
)

func main() {

	r := big.NewInt(100)
	fmt.Println(factorial(r))

}

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}
