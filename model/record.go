package model

// RecordAgentRequest ...
type RecordAgentRequest struct {
	BasicParam
	RequestID  string `json:"request_id"`
	OriginReq  string `json:"origin_req"`
	OriginResp string `json:"origin_resp"`
}

// RecordAgentResponse ...
type RecordAgentResponse struct {
	Error    uint32 `json:"error"`
	ErrorMsg string `json:"error_msg"`
}
