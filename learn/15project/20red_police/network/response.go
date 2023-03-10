package network

type Response struct {
	Data interface{} `json:"data"`
	Err  string      `json:"err"`
}
