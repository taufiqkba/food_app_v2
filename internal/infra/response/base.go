package response

type Response struct {
	HttpStatus     int    `json:"-"`
	StatusCode     string `json:"status_code"`
	Success        bool   `json:"success"`
	Message        string `json:"message"`
	Payload        any    `json:"payload,omitempty"`
	Error          any    `json:"error,omitempty"`
	AdditionalInfo string `json:"additional_info,omitempty"`
}

type OptionResponse func(*Response) *Response

func WithPayload(payload any) OptionResponse {
	return func(r *Response) *Response {
		r.Payload = payload
		return r
	}
}

func WithStatusCode(code string) OptionResponse {
	return func(r *Response) *Response {
		r.StatusCode = code
		return r
	}
}
