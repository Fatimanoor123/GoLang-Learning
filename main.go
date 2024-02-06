package main

import (
	"fmt"

	"math"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("1+1=", 1+1)
	fmt.Println("Hi, fatima how are you?")
	fmt.Println("7.0/3.0=", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)

	// GO with variables
	var a = "initial"
	fmt.Println(a)
	var b, c int = 3, 4
	fmt.Println(b, c)

	//Go with constant
	const number = 45
	fmt.Println(math.Sin(number))

	//for loop
	i := 0
	for i <= 4 {
		i++

	}
	for i := 1; i <= 4; i++ {
		fmt.Println("?")
	}
	fmt.Println(i)

	for n := 0; n <= 10; n++ {
		if n%2 == 0 {
			fmt.Println(n, " is the even number")
			continue
		}
		fmt.Println(n, "Number is not even")
	}
	// Swtich statement
	j := 2
	switch j {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")

	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its week day")

	default:
		fmt.Println(time.Now().Weekday(), "It's a working day")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("I am int")
		case bool:
			fmt.Println("I am boolean")
		default:
			fmt.Println("I don't know")

		}

	}
	whatAmI(23)
	whatAmI(false)
	whatAmI(-123)
	// arrays in GoLang

	var arr [4]int
	arr[3] = 34
	fmt.Println("Elements in arrays: ", arr[0:4])
	// 2D arrays
	var array [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			array[i][j] = i * j
		}
	}
	fmt.Println(array)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	// slicing in arrays

	var s []string
	s = make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("before appending: ", s)
	s = append(s, "f")
	fmt.Println("After appending: ", s)
	l := s[2:5]
	fmt.Println("After slicing: ", l)

	m := make([]string, len(s))
	copy(m, s)
	fmt.Println("After copying the slice: ", m[0:2])
	
}
