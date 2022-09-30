package model

type BaseResponseSchema struct {
	Success bool        `json:"success"`
	TraceId string      `json:"traceId"`
	Error   interface{} `json:"error"`
	Results interface{} `json:"results"`
}

type BaseErrorResponseSchema struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

type BaseErrorValidationSchema struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
