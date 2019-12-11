package models

// 分页数据通用结构
type PageMsg struct {
	Count    int64       `json:"count"`
	Previous interface{} `json:"previous"`
	Next     interface{} `json:"next"`
	Results  interface{} `json:"results"`
}

// RespMsg : http响应数据的通用结构
type RespMsg struct {
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

type UserOperation struct {
	Uid               int64  `json:"uid"`
	RemoteAddr        string `json:"remoteAddr"`
	RemoteUser        string `json:"-"`
	TimeLocal         string `json:"timeLocal"`
	HttpMethod        string `json:"httpMethod"`
	HttpUrl           string `json:"httpUrl"`
	HttpVersion       string `json:"-"`
	Status            string `json:"status"`
	BodyBytesSent     string `json:"bodyBytesSent"`
	HttpReferer       string `json:"httpReferer"`
	HttpUserAgent     string `json:"httpUserAgent"`
	HttpXForwardedFor string `json:"-"`
	HttpToken         string `json:"-"`
	ResType           string `json:"resType"`
	ResId             string `json:"resId"`
}

type VisitorCount struct {
	VisType   string
	ResType   string
	ResId     string
	TimeType  string
	TimeLocal string
	Click     int64
}
