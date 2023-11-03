// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := &big.Int{}
	a.SetString("1111111111111111111111111111", 10)

	b := &big.Int{}
	b.SetString("23232312421451421412412412421421432423452343214", 10)

	sum := &big.Int{}
	sum = sum.Add(a, b)
	fmt.Println("sum =", sum)

	sub := &big.Int{}
	sub = sub.Sub(a, b)
	fmt.Println("sub =", sub)

	div := &big.Int{}
	div = div.Quo(b, a)
	fmt.Println("div =", div)

	mul := &big.Int{}
	mul = mul.Mul(a, b)
	fmt.Println("mul =", mul)

}
