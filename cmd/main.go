package main

import "github.com/gen2brain/beeep"

func main() {
	err := beeep.Notify("PopTart", "You've for a message", "./cmd/poptart.png")

	if err != nil {
		panic(err)
	}
}
