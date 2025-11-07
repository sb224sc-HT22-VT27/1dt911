package main

import (
	"errors"
	"fmt"
	"maps"
	"math"
	"regexp"
	"slices"
	"time"
)

const s string = "constant"

type person struct {
	name string
	age  int
}

type rect struct {
	width, height float64
}

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type base struct {
	num int
}

type container struct {
	base
	str string
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

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
	val3 := 20
	fanon := func() {
		fmt.Println("val3 = ", val3)
		val3 = 30
	}
	fanon()
	fmt.Println("val3 = ", val3)

	// Slide 59
	add := func(a, b int) int {
		return a + b
	}
	runme(add)

	// Slide 61
	nextInt := intSeq(0)
	for range 3 {
		fmt.Println(nextInt())
	}

	newInts := intSeq(10)
	fmt.Println(newInts())

	// Slide 62
	fmt.Println(fact(8) == 40_320)

	// Slide 63
	ival := 1
	fmt.Println("ival = ", ival)
	setToZero(ival)
	fmt.Println("ival = ", ival)

	// Slide 64
	iptr := 1
	fmt.Println("ival = ", iptr)
	setToZeroPtr(&iptr)
	fmt.Println("ival = ", iptr)

	// Slide 67
	iptr = 1
	fmt.Println("Value = ", iptr, ", Pointer = ", &iptr)

	// Slide 68
	const srune = "佛瑞德"
	fmt.Println("Len: ", len(srune))

	// Slide 69
	for i := range len(srune) {
		fmt.Printf("%x ", srune[i])
	}
	fmt.Println()

	// Slide 70
	for _, rv := range srune {
		fmt.Printf("%#U\n", rv)
	}
	fmt.Println()

	// Slide 71
	_ = isT('A') != isT('T')

	// Slide 72
	p1 := person{"Alice", 10}
	fmt.Println(p1)
	p2 := person{name: "Bob"}
	p2.age = 20
	fmt.Println(p2)

	// Slide 73
	var p3 *person
	p3 = newPersonPtr("Carol", 30)
	fmt.Println(p3.name)

	// Slide 75
	dog := struct {
		name   string
		isGood bool
	}{"Rex", true}

	if dog.isGood {
		fmt.Println(dog.name)
	}

	// Slide 77
	r := rect{width: 10.0, height: 5.0}
	fmt.Println("Area: ", r.area())
	fmt.Println("Perimiter: ", r.perim())

	// Slide 79
	measure(&r)

	// Slide 80
	c0 := circle{5}
	measure(c0)

	// Slide 82
	co := container{base: base{1}, str: "my string"}
	fmt.Println(co.base.num)
	fmt.Println(co.num)
	fmt.Println(co.descrbe())

	// Slide 83
	var m0 = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys: ", MapKeys(m0))
	_ = MapKeys(m0)

	// Slide 87
	lst := List[int]{}
	lst.Push(10)
	lst.Push(11)
	lst.Push(12)
	fmt.Println("List: ", lst.GetAll())

	// Slide 89
	if r, e := f1(42); e != nil {
		fmt.Println("f1 failed: ", e)
	} else {
		fmt.Println("f1 worked: ", r)
	}

	// Slide 91
	// if _, e := f1(42); e != nil {
	// 	panic(e)
	// }

	// Slide 93
	defer fmt.Println("Done")
	myfunc()

	// Slide 94
	myfunc0()
	fmt.Println("Survived function call")

	// Slide 97
	regex := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(regex.MatchString("peach"))
	fmt.Println(regex.FindAllString("peach punch pinch", 2))

	// Slide 98
	qs := regexp.MustCompile(`"(.+?)"`)
	fmt.Println(qs.MatchString("peach"))
	fmt.Println(qs.FindAllString("peach punch pinch", 2))

	// Slide 99
	fmt.Println(qs.FindStringSubmatch("A quoted string: \"dog\"")[1])
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

func runme(f func(int, int) int) {
	fmt.Println(f(1, 2))
}

func intSeq(start int) func() int {
	return func() int {
		start++
		return start
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func setToZero(ival int) {
	ival = 0
}

func setToZeroPtr(iptr *int) {
	*iptr = 0
}

func isT(r rune) bool {
	return r == 'T'
}

func newPersonPtr(name string, age int) *person {
	p := person{name: name, age: age}
	return &p
}

func newPerson(name string, age int) person {
	// Slide 74
	return person{name: name, age: age}
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perim: ", g.perim())
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (b base) descrbe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("42 does not work")
	}
	return arg + 3, nil
}

func myfunc() {
	defer fmt.Println("Exiting")
	fmt.Println("Entering")
	fmt.Println("Executing")
}

func myfunc0() int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered error: ", r)
		}
	}()

	a := 0
	return 1 / a
}
