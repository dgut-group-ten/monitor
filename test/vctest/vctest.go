package main

import (
	"fmt"
	"monitor/core/db"
)

func main() {
	fmt.Println(db.GetVisitorCount("pv", "song", "410715051", "hour", 1, 50))
}
