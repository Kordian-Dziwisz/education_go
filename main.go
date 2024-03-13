package main

import "fmt"

func plus(a, b int) (int, string, string) {
	err := "no error"
	msg := "hi"
	return a + b, err, msg
}

func sum(arg ...int) int {
	_sum := 0
	for _, v := range arg {
		_sum += v
	}
	return _sum
}

func main() {
	a := [...][2]int{[...]int{1, 2}, [...]int{3, 4}}
	fmt.Println(a)
	slc := make([]string, 3, 3)
	slc[0] = "hi"
	slc[1] = "hello"
	slc[2] = "bye"
	for _, v := range slc[:] {
		fmt.Println(v)
	}
	fmt.Println(sum(1, 2, 3))
	// fmt.Println(plus(1, 2))
}
