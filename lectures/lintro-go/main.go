package main

import (
	"fmt"
	"maps"
	"math"
	"slices"
	"time"
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
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// Slide 22
	i = 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Slide 23
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend")
	default:
		fmt.Println("It is a weekday")
	}

	// Slide 24
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is after noon")
	}

	// Slide 25
	var arr [5]int
	fmt.Println("emp: ", arr)

	arr[4] = 100
	fmt.Println("set: ", arr)
	fmt.Println("get: ", arr)

	fmt.Println("len: ", len(arr))

	// Slide 27
	barr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", barr)

	// Slide 28
	var twoD [2][3]int

	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

	// Slide 29
	var sarr []string
	fmt.Println("uninit: ", sarr, sarr == nil, len(sarr) == 0)

	sarr = make([]string, 3)
	fmt.Println("emp: ", sarr, " len: ", len(sarr), " cap: ", cap(sarr))

	// Slide 31
	var iarr []int
	fmt.Println("iarr: ", iarr, " len: ", len(iarr), " cap: ", cap(iarr))
	iarr = append(iarr, 10)
	fmt.Println("iarr: ", iarr, " len: ", len(iarr), " cap: ", cap(iarr))

	// Slide 32
	iarr0 := make([]int, 5)
	fmt.Println("iarr0: ", iarr0, " len: ", len(iarr0), " cap: ", cap(iarr0))
	iarr0 = append(iarr0, 10)
	fmt.Println("iarr0: ", iarr0, " len: ", len(iarr0), " cap: ", cap(iarr0))

	// Slide 33
	iarr1 := make([]int, 5, 10)
	fmt.Println("iarr1: ", iarr1, " len: ", len(iarr1), " cap: ", cap(iarr1))
	iarr1 = append(iarr1, 10)
	fmt.Println("iarr1: ", iarr1, " len: ", len(iarr1), " cap: ", cap(iarr1))

	// Slide 34
	carr := make([]string, len(sarr))
	copy(carr, sarr)
	fmt.Println("cpy: ", carr)

	// Slide 35
	sarr0 := []string{"a", "b", "c", "d", "e", "f"}

	l := sarr0[2:5]
	fmt.Println("sl0: ", l)

	l = sarr0[:5]
	fmt.Println("sl1: ", l)

	l = sarr0[2:]
	fmt.Println("sl2: ", l)

	// Slide 37
	tarr := []string{"g", "h", "i"}
	fmt.Println("dcl: ", tarr)

	t0 := []string{"g", "h", "i"}
	if slices.Equal(tarr, t0) {
		fmt.Println("tarr == t0")
	}

	// Slide 38
	twoD0 := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD0[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD0[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD0)

	// Slide 39
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m, " len: ", len(m))

	// Slide 40
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	v2, here := m["k1"]
	fmt.Println("v2: ", v2, " here: ", here)

	// Slide 41
	delete(m, "k2")
	fmt.Println("map: ", m)

	clear(m)
	fmt.Println("map: ", m)

	// Slide 42
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)

	n0 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n0) {
		fmt.Println("n == n0")
	}

	// Slide 43
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum: ", sum)

	// Slide 44
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	// Slide 46
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Slide 47
	for k := range kvs {
		fmt.Println("key: ", k)
	}

	// Slide 48
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	// Slide 49
	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	// Slide 50
	fmt.Println("1 + 2 + 3 = ", plusPlus(1, 2, 3))

	// Slide 51
	val0, val1 := vals()
	_, val2 := vals()
	fmt.Println("val0: ", val0, "val1: ", val1, "val2: ", val2)

	// Slide 54
	_ = fsum(1, 2)
	_ = fsum(1, 2, 3)

	// Slide 56
	nums0 := []int{1, 2, 3, 4}
	_ = fsum(nums0...)

	// Slide 57
	func() { fmt.Println("Anon!") }()

	// Slide 58
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (x, y int) {
	x = 3
	y = 7
	return x, y
}

func fsum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}
