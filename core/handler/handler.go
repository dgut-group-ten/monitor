package handler

import (
	"fmt"
	"monitor/core/db"
	"monitor/core/models"
	"monitor/core/util"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := util.RespMsg{Msg: "Hello World"}
	_, _ = w.Write(resp.JSONBytes())
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	count, _ := db.CountUserOperationDB(25)
	uoList, err := db.GetUserHistoryDB(25, 2, 20)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	resp := util.RespMsg{
		Msg: "用户历史行为",
		Data: struct {
			Count  int64                  `json:"count"`
			UOList []models.UserOperation `json:"uoList"`
		}{
			Count:  count,
			UOList: uoList,
		},
	}
	_, _ = w.Write(resp.JSONBytes())
}
