package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int // s is a slice of int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 = %#v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	s3 := s2[1:4]

	fmt.Printf("s3 = %#v, len=%d, cap=%d\n", s3, len(s3), cap(s3))

	// fmt.Println("s2[1] = ", s2[:100])  panic
	s3 = append(s3, 100)
	fmt.Printf("s3 = %#v, len=%d, cap=%d\n", s3, len(s3), cap(s3))

	fmt.Printf("s3 (append) = %#v\n", s3)
	fmt.Printf("s2 (append) = %#v\n", s2)

	fmt.Printf("s2 = %#v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3 = %#v, len=%d, cap=%d\n", s3, len(s3), cap(s3))

	// s4 := []int{1, 2, 3, 4, 5, 6, 7}
	var s4 []int
	// s4 := make([]int, 0, 100) // Single allocation
	for i := 0; i < 100; i++ {
		s4 = appendInt(s4, i)
	}

	fmt.Printf("s4 = %#v, len=%d, cap=%d\n", s4, len(s4), cap(s4))

	fmt.Println(concat([]string{"a", "b", "c"}, []string{"d", "e", "f"})) // [a b c d e f]

	vs := []float64{2, 1, 3}
	fmt.Println(median(vs)) // 2
	vs = []float64{2, 1, 3, 4}
	fmt.Println(median(vs)) // 2.5

	fmt.Println(vs)

	fmt.Println(median(nil))

	// fmt.Println(reflect.TypeOf(2))
}

func appendInt(s []int, v int) []int {
	i := len(s)
	// if len(s)+1 > cap(s) {
	// 	newSlice := make([]int, len(s)*2)
	// 	copy(newSlice, s)
	// 	s = newSlice
	// }
	if len(s) < cap(s) { // enough space in underlying array
		s = s[:len(s)+1]
	} else { // need to re-alocate and copy
		fmt.Printf("reallocate: %d->%d\n", len(s), len(s)*2+1)
		s2 := make([]int, len(s)*2+1)
		copy(s2, s)
		s = s2[:len(s)+1]
	}
	s[i] = v
	return s
}

func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}
	// copy in order to not to change values
	nums := make([]float64, len(values))
	copy(nums, values)
	sort.Float64s(nums)
	i := len(nums) / 2
	// if len(nums)&1 == 1 {
	if len(nums)%2 == 1 {
		return nums[i], nil
	}

	v := (nums[i-1] + nums[i]) / 2
	return v, nil
}

func concat(s1, s2 []string) []string {
	s3 := make([]string, len(s1)+len(s2))
	copy(s3, s1)
	copy(s3[len(s1):], s2)
	return s3
}
