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
		//获取参数
		if err := r.ParseForm(); err != nil {
			fmt.Println(w, "ParseForm() err: "+err.Error())
			return
		}
		uid, err := strconv.ParseInt(r.Form.Get("uid"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := models.RespMsg{Msg: "请提供用户ID"}
			_, _ = w.Write(util.JSONBytes(resp))
			return
		}

		// 拿出总数并计算页码
		count, _ := db.CountUserOperationDB(uid)
		resp := models.PageMsg{Count: count}
		p, ps := resp.GenPageNum(r)

		// 拿出要返回的数据
		uoList, err := db.GetUserHistoryDB(uid, p, ps)
		if err != nil {
			fmt.Println(err)
		}
		resp.Results = uoList

		// 返回结果
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(util.JSONBytes(resp))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed) // 其他操作不允许
	}
}

func VisitCountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		//获取参数
		if err := r.ParseForm(); err != nil {
			fmt.Println(w, "ParseForm() err: "+err.Error())
			return
		}

		visType := r.Form.Get("visType")
		resType := r.Form.Get("resType")
		resId := r.Form.Get("resId")
		timeType := r.Form.Get("timeType")

		// 拿出总数并计算页码
		count, _ := db.CountVisitorCountDB(visType, resType, resId, timeType)
		resp := models.PageMsg{Count: count}
		p, ps := resp.GenPageNum(r)

		// 拿出要返回的数据
		vcList, err := db.GetVisitorCount(visType, resType, resId, timeType, p, ps)
		if err != nil {
			fmt.Println(err)
		}
		resp.Results = vcList

		// 返回结果
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(util.JSONBytes(resp))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed) // 其他操作不允许
	}
}
