package main

import (
	"fmt"
)

// maxX and maxY are the maximum allowed coordinates
const (
	maxX = 1000
	maxY = 600
)

// Item is an item in the game
type Item struct {
	X, Y int
}

type Player struct {
	Name string
	Item
	Keys []Key // slice of Key
}

func (i Item) String() string {
	return fmt.Sprintf("[%d,%d]", i.X, i.Y)
}

// i is the receiver
// if you want to mutate, use pointer receiver
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

// NewItem creates a new item
func NewItem(x, y int) (*Item, error) {

	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("Invalid coordinates (%d,%d)", x, y)
	}

	i := Item{
		X: x,
		Y: y,
	}

	// The go compiler does 'escape analysis' and allochttps://www.youtube.com/watch?v=q4CjJK6c8cgates the object on the heap
	return &i, nil
}

func main() {
	var i1 Item
	fmt.Println(i1) // {0 0}
	fmt.Printf("i1: %#v\n", i1)

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v\n", i2)

	i3 := Item{
		X: 10,
		Y: 20,
	}

	fmt.Printf("i3: %#v\n", i3)

	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(1001, 20))

	i3.Move(100, 200)
	fmt.Printf("i3 move: %#v\n", i3)

	p1 := Player{
		Name: "Player 1",
		Item: Item{500, 300},
	}
	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1.X: %#v\n", p1.X)
	fmt.Printf("p1.Item.X: %#v\n", p1.Item.X)

	ms := []mover{
		&i1, &p1, &i2}
	moveAll(ms, 10, 20)

	for _, v := range ms {
		fmt.Printf("%#v\n", v)
	}

	b := NewNumber[int]("")
	fmt.Println(b)

	k := Jade
	fmt.Println(k)
	fmt.Println(Key(17))

	// time.Time import json.Marshaler interface
	// json.NewEncoder(os.Stdout).Encode(time.Now())

	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
	p1.FoundKey(Key(17))
	fmt.Println(p1.Keys)
}

type mover interface {
	Move(dx, dy int)
}

func moveAll(ms []mover, dx, dy int) {
	for _, m := range ms {
		m.Move(dx, dy)
	}
}

func NewNumber[T int | float64](kind string) T {
	if kind == "int" {
		return 0
	}
	return 0.0
}

type Key byte

// Go's version of enums
const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey // internal - not exported
)

func (k Key) String() string {
	switch k {
	case Jade:
		return "Jade"
	case Copper:
		return "Copper"
	case Crystal:
		return "Crystal"
		// default:
		// 	return "Unknown"
	}
	return fmt.Sprintf("<Key %d>", k)
}

/* Exercise
- Add a "Keys" field to the Player struct which is a slice of Key
- Add a "FoundKey (k Key) error" method to Player which add ke to key if it's not already there
- Err if k is not one of the known keys
*/

func (p *Player) FoundKey(k Key) error {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("Unknown key %d", k)
	}
	for _, v := range p.Keys {
		if v == k {
			return nil
		}
	}
	p.Keys = append(p.Keys, k)
	return nil
	// if k < Jade || k >= invalidKey {
	// 	return fmt.Errorf("invalid key: %#v", k)
	// }

	// // if !containsKey(p.Keys, k) {
	// if !slices.Contains(p.Keys, k) {// slices.Contains is a generic function -- golang.org/x/exp/slices
	// 	p.Keys = append(p.Keys, k)
	// }

	// return nil
}

func containsKey(keys []Key, k Key) bool {
	for _, k2 := range keys {
		if k2 == k {
			return true
		}
	}
	return false
}
