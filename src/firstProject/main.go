package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	f := fib()

	fmt.Println("hallo TINF19B3")
	fmt.Println(add(1, 2))
	fmt.Println(swap("welt", "Bier"))
	fmt.Println(f(), f(), f(), f(), f(), f())

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		i++
	}
	fmt.Println(i)

	fmt.Println(SwitchIt("9142"))

	fmt.Println(stringutil.Reverse("!selpmaxe oG ,olleH"))

}

func add(x, y int) (result int) {
	result = x + y
	return
}

var (
	a, b string
	z    int
	c    complex64
)

func swap(x, y string) (string, string) {
	return y, x
}

func SwitchIt(a string) int {
	switch a {
	case "":
		return 0
	case "abc":
		return 1
	default:
		return 100
	}
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
