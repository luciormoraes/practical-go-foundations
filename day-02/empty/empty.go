package main

import "fmt"

func main() {

	// empty interface

	var i any
	i = 7
	fmt.Println(i)
	i = "hello"
	fmt.Println(i)

	// rule of thumb: don't use any

	s := i.(string) // type assergion
	fmt.Println(s)

	// comma, ok
	n, ok := i.(int) // panic
	if !ok {
		fmt.Println("not an int")
	} else {
		fmt.Println(n)
	}

	// type switch
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("unknown:", v)
	}

	fmt.Println(maxInts([]int{1, 2, 3}))
	fmt.Println(maxFloat64([]float64{1.1, 2.2, 3.3}))
	fmt.Println(max([]int{1, 2, 3}))
	fmt.Println(max([]float64{1.1, 2.2, 3.3}))
}

func maxInts(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}
func maxFloat64(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

type Number interface {
	int | float64
}

func max[T Number](nums []T) T {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}
