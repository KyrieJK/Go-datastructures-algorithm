package Karastuba

import (
	"math"
	"math/big"
)

func k_multiply(a, b *big.Int) *big.Int {
	if a.Cmp(big.NewInt(10)) < 1 || b.Cmp(big.NewInt(10)) < 1 {
		return mul(a, b)
	}

	pivot := _pivot(a, b)
	leftA, rightA := _split(a, uint(pivot))
	leftB, rightB := _split(b, uint(pivot))

	z0 := k_multiply(leftA, leftB)
	z1 := k_multiply(rightA, rightB)
	z2 := k_multiply(add(leftA, rightA), add(leftB, rightB))
	z2 = sub(z2, add(z0, z1))

	temp0 := mul(z0, big.NewInt(int64(math.Pow(10.0, 2.0*float64(pivot)))))
	temp1 := mul(z2, big.NewInt(int64(math.Pow(10.0, float64(pivot)))))
	temp2 := add(temp0, temp1)

	return add(temp2, z1)
}

func sub(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, b)
}

func add(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}

func _pivot(a, b *big.Int) int {
	len_a := len(a.String())
	len_b := len(b.String())

	if len_a > len_b {
		return len_a / 2
	} else {
		return len_b / 2
	}
}

func _split(a *big.Int, pivot uint) (left, right *big.Int) {
	denominator := big.NewInt(int64(math.Pow(10.0, float64(pivot))))
	left = big.NewInt(0).Div(a, denominator)
	right = big.NewInt(0).Mod(a, denominator)

	return left, right
}

func mul(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, b)
}
