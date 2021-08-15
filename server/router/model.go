package router

import "encoding/json"



type RequestSessionBody struct {
	SessionID string `json:"session_id"`
}

type UpdateSessionBody struct {
	SessionID string `json:"session_id"`
	Tree      map[string]string
}

type CompileSessionBody struct {
	SessionID string `json:"session_id"`
}

type DownloadCompiledSessionBody struct {
	SessionID string `json:"session_id"`
}




type Response struct {
	Status string `json:"status"`
	Result interface{} `json:"result"`
}

func (r *Response) Bytes() []byte {
	res, _ := json.Marshal(&r)
	return res
}

func SuccessfulResponseFrom(result interface{}) *Response {
	return &Response {
		Status: "ok",
		Result: result,
	}
}

func FailedResponseFrom(result interface{}) *Response {
	return &Response {
		Status: "error",
		Result: result,
	}
}

