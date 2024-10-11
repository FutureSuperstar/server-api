package request

type ZstdRequest struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}

type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ZstdResponse struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}
