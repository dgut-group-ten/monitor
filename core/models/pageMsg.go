package models

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

// 分页数据通用结构
type PageMsg struct {
	Count    int64       `json:"count"`
	Previous interface{} `json:"previous"`
	Next     interface{} `json:"next"`
	Results  interface{} `json:"results"`
}

func (pmsg *PageMsg) GenPageNum(r *http.Request) (p, ps int64) {

	// 看一下有没有p,ps这个参数
	p, err := strconv.ParseInt(r.Form.Get("p"), 10, 64)
	if err != nil {
		p = 1
		r.RequestURI += "&p=1"
	}
	ps, err = strconv.ParseInt(r.Form.Get("ps"), 10, 64)
	if err != nil {
		ps = 10
	}

	re, _ := regexp.Compile("p=[0-9]*")
	// 上一页
	if p > 1 {
		page := "p=" + fmt.Sprintf("%d", p-1)
		r.RequestURI = re.ReplaceAllString(r.RequestURI, page)
		pmsg.Previous = "http://" + r.Host + r.RequestURI
	}

	// 下一页
	if p < pmsg.Count/ps+1 {
		page := "p=" + fmt.Sprintf("%d", p+1)
		r.RequestURI = re.ReplaceAllString(r.RequestURI, page)
		pmsg.Next = "http://" + r.Host + r.RequestURI
	}

	return p, ps
}
