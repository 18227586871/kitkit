package demo

type PingReq struct {
	Ping string `json:"ping"`
}

type PingResp struct {
	Pong string `json:"pong"`
}
