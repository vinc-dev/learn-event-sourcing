package transport

type Message struct {
	Str string `json:"message"`
}

// Success represents base response structure if a request is success
type Success struct {
	Result   interface{}       `json:"data"`
	Metadata interface{}       `json:"_metadata,omitempty"`
	Header   map[string]string `json:"-"`
}

func OK() *Success {
	return &Success{
		Result: Message{
			Str: "OK",
		},
	}
}

type errorResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Debug   *errorDebug `json:"_debug,omitempty"`
}

type errorDebug struct {
	Trace   string `json:"trace,omitempty"`
	Message string `json:"err_message,omitempty"`
}
