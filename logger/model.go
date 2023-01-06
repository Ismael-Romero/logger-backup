package logger

type logData struct {
	Type  string      `json:"type,omitempty"`
	Level string      `json:"level,omitempty"`
	Date  string      `json:"date,omitempty"`
	Body  interface{} `json:"body,omitempty"`
}

type defaultBody struct {
	Message string `json:"message,omitempty"`
}

type detailsRequest struct {
	Message  string `json:"message,omitempty"`
	From     string `json:"from,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Method   string `json:"method,omitempty"`
	URL      string `json:"URL,omitempty"`
	Status   int    `json:"status,omitempty"`
}
