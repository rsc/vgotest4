package main // import "github.com/rsc/vgotest4"

import (
	"fmt"

	"github.com/myitcv/vgo_example_compat"
	"github.com/myitcv/vgo_example_compat/sub"
)

func main() {
	fmt.Printf("X=%d, Y=%d\n", vgo_example_compat.X, sub.Y)
}
