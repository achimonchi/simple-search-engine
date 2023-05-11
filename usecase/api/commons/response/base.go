package response

type Response struct {
	Status      int         `json:"status"`
	Error       interface{} `json:"error,omitempty"`
	Message     string      `json:"message"`
	ErrorDetail interface{} `json:"error_detail,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}
