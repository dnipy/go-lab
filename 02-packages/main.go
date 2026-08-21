package main

import (
	"fmt"

	"github.com/dnipy/go-lab/02-packages/greeting"
)

func main() {
	hello := greeting.Hello()
	goodbye := greeting.Goodbye()

	fmt.Println(hello)
	fmt.Println(goodbye)
}
