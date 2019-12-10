package main

import (
	"fmt"
	"monitor/core/util"
)

func main() {
	timestamp := "1575788400"

	fmt.Println(util.GetTime(timestamp))
}
