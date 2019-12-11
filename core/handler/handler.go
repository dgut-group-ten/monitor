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
		resp := models.RespMsg{Msg: "Hello World"}
		_, _ = w.Write(util.JSONBytes(resp))
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

		// 上一页
		var pre interface{}
		if p > 1 {
			pre = "http://" + r.Host + "/uo/history/?p=" + fmt.Sprintf("%d", p-1)
		} else {
			pre = nil
		}
		// 下一页
		var next interface{}
		if p < count/ps+1 {
			next = "http://" + r.Host + "/uo/history/?p=" + fmt.Sprintf("%d", p+1)
		} else {
			next = nil
		}

		w.WriteHeader(http.StatusOK)
		resp := models.PageMsg{
			Count:    count,
			Previous: pre,
			Next:     next,
			Results:  uoList,
		}

		_, _ = w.Write(util.JSONBytes(resp))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed) // 其他操作不允许
	}
}
