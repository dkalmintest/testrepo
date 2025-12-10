package main

import (
	transformer "gh/transform/transformer"
	"fmt"
)

func main() {
	println("Start of transform")
	
	transformResult := transformer.Transform()

	if transformResult != true {
			fmt.Printf("Expected true, got '%v'", transformResult)
		} else {
			println("We transformed!")
		}

	println("End of transform")
}
