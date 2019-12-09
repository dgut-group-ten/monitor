package handler

import (
	"monitor/core/util"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := util.RespMsg{Msg: "Hello World"}
	_, _ = w.Write(resp.JSONBytes())
}
