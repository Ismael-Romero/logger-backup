package logger

import "time"

func (l *Logger) Info(m string) {
	content := LogData{
		Level:   "Info",
		Date:    time.Now().Format("2006-01-02 15-04-05"),
		Message: m,
	}
	l.encoder(content)
}

func (l *Logger) Error(m string) {
	content := LogData{
		Level:   "Error",
		Date:    time.Now().Format("2006-01-02 15-04-05"),
		Message: m,
	}
	l.encoder(content)
}

func (l *Logger) Warning(m string) {
	content := LogData{
		Level:   "Warning",
		Date:    time.Now().Format("2006-01-02 15-04-05"),
		Message: m,
	}
	l.encoder(content)
}
