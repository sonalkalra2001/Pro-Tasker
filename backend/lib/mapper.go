package lib

type ConnDet struct {
	Tx       bool `json:"tx"`
	Writable bool `json:"writable"`
}

type ApiResp struct {
	Status   string      `json:"status"`
	Data     interface{} `json:"data"`
	DebugMsg string      `json:"debug_msg"`
}
