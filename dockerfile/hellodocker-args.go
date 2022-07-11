package main

import (
	"fmt"
	"os"
)

func main() {
	name := "world"
	args := os.Args
	if len(args) > 1 {
		name = args[1]
	}
	fmt.Printf("Hello %s!\n", name)
}
