package main

import (
	"fmt"
	"time"
)

var (
	colorSuccess = "\033[32m"
	// colorError = "\033[31m"
)

func main() {
	start := time.Now()

	fmt.Println(colorSuccess, "Successfully init app in", time.Since(start))
}
