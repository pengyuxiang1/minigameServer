package models

type RspComm struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Info string `json:"info"`
}
