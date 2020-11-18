package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	n, p, q, e, d, tmp int
)

func Equation(a int, b int, x *int, y *int) int {
	if a == 0 {
		*x, *y = 0, 1
		return b
	}
	var x2, y2 int
	d2 := Equation(b%a, a, &x2, &y2)
	*x = y2 - (b/a)*x2
	*y = x2
	return d2
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func IsPrime(x int) bool {
	if x == 2 {
		return true
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
func GetRandomPrime(x, y int) (int, int) {
	a := []int{}
	for i := x; i <= y; i++ {
		if IsPrime(i) {
			a = append(a, i)
		}
	}
	i, j := rand.Intn(len(a)), rand.Intn(len(a))
	for i == j {
		j = rand.Intn(len(a))
	}
	return a[i], a[j]
}
func pow(x, y int) int {
	if y == 0 {
		return 1
	}
	if y%2 == 1 {
		return ((pow(x, y-1) % n) * x) % n
	} else {
		cur := (pow(x, y/2)) % n
		return cur * cur % n
	}
}
func f(x int) int {
	return pow(x, 2)
}
func phi(x int) int {
	result := x
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			for x%i == 0 {
				x /= i
			}
			result -= result / i
		}
	}
	if x > 1 {
		result -= result / x
	}
	return result
}
func main() {
	rand.Seed(time.Now().UnixNano())
	p, q = GetRandomPrime(9000, 10000)
	p, q = 9241, 9689
	n = p * q
	for i := 2; i < (p-1)*(q-1); i++ {
		e = i
		if gcd(e, (p-1)*(q-1)) == 1 {
			Equation(e, (p-1)*(q-1), &d, &tmp)
			if d > 0 {
				break
			}
		}
	}
	x := rand.Intn((1 << 31))
	r := rand.Intn((1 << 31))
	for gcd(r, n) != 1 {
		r = rand.Intn((1 << 31))
	}
	y := f(x) * pow(r, e) % n
	y = pow(y, d)
	y = y * pow(r, phi(n)-1) % n
	fmt.Println(pow(y, e) == f(x))
}
