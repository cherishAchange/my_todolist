package middleware

type BaseResponse struct {
	StatusCode int32  `json:"statusCode"`
	StatusMsg  string `json:"statusMsg,omitempty"`
}

func GetRes(code int32, msg string) BaseResponse {
	baseResponse := BaseResponse{StatusCode: code, StatusMsg: msg}
	return baseResponse
}
