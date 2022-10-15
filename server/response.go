package main

type Resposne struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Result  interface{} `json:"result"`
}
type Error struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Error   interface{} `json:"error"`
}
