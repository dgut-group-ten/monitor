package main

import (
	"fmt"
	"monitor/core/db"
)

func main() {
	fmt.Println(db.CountUserOperationDB(25))
	fmt.Println(db.GetUserHistoryDB(25, 2, 20))
}
