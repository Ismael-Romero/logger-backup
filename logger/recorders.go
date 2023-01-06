package logger

import (
	"net/http"
	"time"
)

func (l *Logger) Info(m string) {
	content := logData{
		Type:  "System",
		Level: "info",
		Date:  time.Now().Format("2006-01-02 15:04:05"),
		Body: defaultBody{
			Message: m,
		},
	}
	encoder(l, content)
}

func (l *Logger) Error(m string) {
	content := logData{
		Type:  "System",
		Level: "Error",
		Date:  time.Now().Format("2006-01-02 15:04:05"),
		Body: defaultBody{
			Message: m,
		},
	}
	encoder(l, content)
}

func (l *Logger) Warning(m string) {
	content := logData{
		Type:  "System",
		Level: "Warning",
		Date:  time.Now().Format("2006-01-02 15:04:05"),
		Body: defaultBody{
			Message: m,
		},
	}
	encoder(l, content)
}

func (l *Logger) Response200(r *http.Request, m string) {
	content := logData{
		Type:  "Response",
		Level: "info",
		Date:  time.Now().Format("2006-01-02 15:04:05"),
		Body: detailsRequest{
			Message:  m,
			From:     r.RemoteAddr,
			Protocol: r.Proto,
			Method:   r.Method,
			URL:      r.RequestURI,
			Status:   200,
		},
	}
	encoder(l, content)
}
