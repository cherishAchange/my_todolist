package middleware

type BaseResponse struct {
	StatusCode int32  `json:"statusCode"`
	StatusMsg  string `json:"statusMsg,omitempty"`
}

var baseResponse *BaseResponse

func GetRes(code int32, msg string) *BaseResponse {
	baseResponse.StatusCode = code
	baseResponse.StatusMsg = msg
	return baseResponse
}
