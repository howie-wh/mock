package model

type BasicParam struct {
	Project string `json:"project"`
	Env     string `json:"env"`
	Region  string `json:"region"`
	PFB     string `json:"pfb"`
	MainAPI string `json:"main_api"`
	SubAPI  string `json:"sub_api"`
}

// MKAgentRequest ...
type MKAgentRequest struct {
	BasicParam
	SrcRequest string `json:"src_request"`
}

// MKAgentResponse ...
type MKAgentResponse struct {
	Data     string `json:"data"`
	Error    uint32 `json:"error"`
	ErrorMsg string `json:"error_msg"`
}
