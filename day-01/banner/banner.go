package main

import (
	"fmt"
)

func banner(text string, width int) {
	// padding := (width - len(text)) / 2
	// padding := (width - utf8.RuneCountInString(text)) / 2
	// for i := 0; i < padding; i++ {
	// 	fmt.Print(" ")
	// }
	// fmt.Println(text)
	// for i := 0; i < width; i++ {
	// 	fmt.Print("-")
	// }
	// fmt.Println()
	// x, y := 1, "1"
	// fmt.Printf("%v, %v \n", x, y)
	// fmt.Printf("%#v, %#v \n", x, y)

}

func isPalindrome(text string) bool {
	// for i := 0; i < len(text)/2; i++ {
	// 	if text[i] != text[len(text)-1-i] {
	// 		return false
	// 	}
	// }
	// return true
	// for i, v := range text {
	// 	if v != rune(text[len(text)-1-i]) {
	// 		return false
	// 	}
	// }
	// return true
	runes := []rune(text)
	for i := 0; i < len(runes)/2; i++ {
		fmt.Println(i, len(runes)-1-i)
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}

	}
	return true
}

func main() {
	banner("Go", 6)
	banner("G☺", 6)
	s1 := "Go"
	s2 := "G☺"
	fmt.Println(len(s1), len(s2))
	// for i := 0; i < len(s1); i++ {
	// 	fmt.Printf("%x ", s1[i])
	// }

	// for i, v := range s2 {
	// 	fmt.Printf("%d %v \n", i, v)
	// }

	// fmt.Println()
	// for i := 0; i < len(s2); i++ {
	// 	fmt.Printf("%x \n", s2[i])
	// }

	// b := s2[0]
	// fmt.Printf("%c of type %T \n", b, b)

	g := isPalindrome("g")
	gogo := isPalindrome("gogo")
	gog := isPalindrome("gog")
	gogog := isPalindrome("g☺g☺g")

	fmt.Printf("%v, %v, %v %v\n", g, gogo, gog, gogog)
}
