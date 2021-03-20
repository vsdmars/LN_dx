package main

import (
	"dx/cmd"
	"fmt"
)

func main() {
	if err := cmd.Cmd(); err != nil {
		fmt.Printf("dx error: %s\n", err.Error())
	}
}
