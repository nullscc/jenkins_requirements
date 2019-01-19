package models

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Info interface{} `json:"info"`
}

type Test struct {
	Test string `json:"test"`
}
