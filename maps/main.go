package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#DJD298",
		"green": "#EE8833",
		"white": "#FFFFFF",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
