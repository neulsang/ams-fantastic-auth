package beautiprint

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

// Logo prints beautiful ASCII art from text.
func Logo(text string) {
	fmt.Println(figure.NewFigure(text, "doom", true))
}
