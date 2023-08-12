package main

import "fmt"

const (
	maxX = 1000
	maxY = 500
)

// Item is an item in the game
type Item struct {
	X int
	Y int
}

// Player represents a user of an Item
type Player struct {
	Name string
	Item // embeds Item
}

type mover interface {
	// Move(int, int)
	Move(x, y int)
}

// func NewItem(x, y int) Item {
// func NewItem(x, y int) *Item {
// func NewItem(x, y int) (Item, error) {
func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}

	// The Go compiler does "escape analysis" and will allocation
	// i on the heap.
	// If curious about what's being moved to the heap: go build -gcflags=-m
	return &i, nil
}

// i is called "the receiver"
// similar to "this" in JS and "self" in Python
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

// Rule of Thumb: generally, methods should accept interfaces and return types

func main() {
	var i1 Item
	fmt.Println(i1)
	fmt.Printf("i1: %#v\n", i1)

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v\n", i2)

	i3 := Item{
		Y: 10,
		X: 20,
	}
	fmt.Printf("i3: %#v\n", i3)

	fmt.Println(NewItem(35, 89))
	fmt.Println(NewItem(35, -89))

	i3.Move(5, 10)
	fmt.Printf("i3 (move): %#v\n", i3)

	p1 := Player{
		Name: "James",
		Item: Item{500, 400},
	}
	fmt.Printf("p1: %#v\n", p1)
	// Player can access X directly because the embedded
	// Item's fields are "lifted up" to the parent's level.
	// If the parent has the same fields then it defaults
	// to the parent's fields.
	fmt.Printf("p1.X: %#v\n", p1.X)

	p1.Move(20, 25)
	fmt.Printf("p1 (move): %#v\n", p1)

	ms := []mover{
		&i1,
		&p1,
		&i2,
	}
	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Println(m)
	}
}
