package main

import "fmt"

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
