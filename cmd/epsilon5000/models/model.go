package models

type ResponseError struct {
	Status string
	Description string
}

type ServicesCount struct {
	Ip   string `json:"ip"`
	Spd  string `json:"spd"`
	Iptv string `json:"iptv"`
	Sip  string `json:"sip"`
	Addr string `json:"addr"`
}

type Ip struct {
	Ip []string `json:"ip"`
}

type Id struct {
	Id []string `json:"id"`
}

type СameraIncidents struct {
	Start string `json:"start"`
	Stop string `json:"stop"`
	Host string `json:"host"`
	Name string `json:"name"`
	Duration string `json:"duration"`
}


