package handler

import (
	"fmt"
	"monitor/core/db"
	"monitor/core/models"
	"monitor/core/util"
	"net/http"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		resp := util.RespMsg{Msg: "Hello World"}
		_, _ = w.Write(resp.JSONBytes())
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed) // 其他操作不允许
	}
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		if err := r.ParseForm(); err != nil {
			fmt.Println(w, "ParseForm() err: "+err.Error())
			return
		}
		uid, _ := strconv.ParseInt(r.Form.Get("uid"), 10, 64)
		p, _ := strconv.ParseInt(r.Form.Get("p"), 10, 64)
		ps, _ := strconv.ParseInt(r.Form.Get("ps"), 10, 64)

		count, _ := db.CountUserOperationDB(uid)
		uoList, err := db.GetUserHistoryDB(uid, p, ps)
		if err != nil {
			fmt.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		resp := util.RespMsg{
			Msg: "用户历史行为",
			Data: struct {
				Count  int64                  `json:"count"`
				P      int64                  `json:"p"`
				PS     int64                  `json:"ps"`
				UOList []models.UserOperation `json:"uoList"`
			}{
				Count:  count,
				P:      p,
				PS:     ps,
				UOList: uoList,
			},
		}
		_, _ = w.Write(resp.JSONBytes())
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed) // 其他操作不允许
	}
}
