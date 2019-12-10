package models

type UserOperation struct {
	Uid               int64
	RemoteAddr        string
	RemoteUser        string
	TimeLocal         string
	HttpMethod        string
	HttpUrl           string
	HttpVersion       string
	Status            string
	BodyBytesSent     string
	HttpReferer       string
	HttpUserAgent     string
	HttpXForwardedFor string
	HttpToken         string
	ResType           string
	ResId             string
}

type VisitorCount struct {
	VisType   string
	ResType   string
	ResId     string
	TimeType  string
	TimeLocal string
	Click     int64
}
