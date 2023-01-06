package logger

type LogData struct {
	Level   string `json:"level,omitempty"`
	Date    string `json:"date,omitempty"`
	Message string `json:"message,omitempty"`
}
