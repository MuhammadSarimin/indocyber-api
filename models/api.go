package models

type Response struct {
	ResponseCode    string      `json:"response_code"`
	ResponseMessage string      `json:"response_message"`
	ResponseData    interface{} `json:"response_data,omitempty"`
}
