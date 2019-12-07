package handler

import (
	util2 "monitor/core/util"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := util2.RespMsg{Msg: "Hello World"}
	util2.Logerr(w.Write(resp.JSONBytes()))
}
