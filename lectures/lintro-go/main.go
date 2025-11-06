package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	// Slide 3
	fmt.Println("Hello, world")

	// Slide 5
	fmt.Println("go" + "lang")
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Slide 7
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// Slide 12
	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	// Slide 13
	fmt.Println(s)

	const n_c = 500_000_000

	const d_c = 3e20 / n_c
	fmt.Println(d_c)

	fmt.Println(int64(d_c))
	fmt.Println(math.Sin(n_c))

	// Slide 16
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Slide 17
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Slide 18
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// Slide 19
	if 8%4 == 0 {
		fmt.Println("8 is divisable by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// Slide 20
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Slide 21
}
