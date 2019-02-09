package main

import (
	"fmt"

	"github.com/makyo/gotui/ansi"
)

func main() {
	fmt.Printf("%s", ansi.CapabilityCheck.String())
}
