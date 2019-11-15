package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
	}

	/* colors["red"] = "#ff0000"
	colors["white"] = "#fff"

	delete(colors, "white") */

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
