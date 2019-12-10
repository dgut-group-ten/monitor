package main

import (
	"fmt"
	"monitor/core/db"
)

func main() {
	fmt.Println(db.CountUserOperationDB(28))
}
