package logger

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	ext        = ".json"
	flagCreate = os.O_CREATE | os.O_RDWR
	flagOpen   = os.O_RDWR
)

type dataLogger struct {
	Id      string     `json:"id,omitempty"`
	Level   string     `json:"level,omitempty"`
	Type    string     `json:"type,omitempty"`
	Date    string     `json:"date,omitempty"`
	Details detailsLog `json:"details"`
}

type detailsLog struct {
	Message   string    `json:"message,omitempty"`
	NetReport netReport `json:"netReport"`
}

type netReport struct {
	From   string        `json:"from,omitempty"`
	Method string        `json:"method,omitempty"`
	Url    string        `json:"url,omitempty"`
	Status int           `json:"status,omitempty"`
	Time   time.Duration `json:"time,omitempty"`
}

type Logger struct {
	path       string
	name       string
	pathOrigin string
	backupTime time.Duration
	data       []dataLogger
}

type ManagerLogger interface {
	Backup()
	Error()
	Info()
	Network200()
	Network404()
	Network500()
	NetworkStatus()
	Warning()
}

func New(path, name string, time time.Duration) *Logger {

	_, err := openLogFile(parserPath(path, name), flagCreate)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return &Logger{
		path:       parserPath(path, name),
		name:       name,
		pathOrigin: path,
		backupTime: time,
		data:       []dataLogger{},
	}
}

func (l *Logger) Error(message, _type string) {

	logData := dataLogger{
		Id:    uuid.NewString(),
		Level: "Error",
		Type:  _type,
		Date:  time.Now().Format("2006-01-02 15-04-05"),
		Details: detailsLog{
			Message: message,
		},
	}

	l.data = append(l.data, logData)
	encodeData(l.data, l.path)
}

func (l *Logger) Info(message, _type string) {

	logData := dataLogger{
		Id:    uuid.NewString(),
		Level: "info",
		Type:  _type,
		Date:  time.Now().Format("2006-01-02 15-04-05"),
		Details: detailsLog{
			Message: message,
		},
	}

	l.data = append(l.data, logData)
	encodeData(l.data, l.path)
}

func (l *Logger) Warning(message, _type string) {

	logData := dataLogger{
		Id:    uuid.NewString(),
		Level: "Warning",
		Type:  _type,
		Date:  time.Now().Format("2006-01-02 15-04-05"),
		Details: detailsLog{
			Message: message,
		},
	}

	l.data = append(l.data, logData)
	encodeData(l.data, l.path)
}

func (l *Logger) Network200(r *http.Request, t time.Duration, message string) {
	logData := dataLogger{
		Id:    uuid.NewString(),
		Level: "info",
		Type:  "Respone",
		Date:  time.Now().Format("2006-01-02 15-04-05"),
		Details: detailsLog{
			Message: message,
			NetReport: netReport{
				From:   r.RemoteAddr,
				Method: r.Method,
				Url:    r.RequestURI,
				Status: 200,
				Time:   t,
			},
		},
	}

	l.data = append(l.data, logData)
	encodeData(l.data, l.path)
}

func (l *Logger) Network404(r *http.Request, t time.Duration, message string) {
	logData := dataLogger{
		Id:    uuid.NewString(),
		Level: "Warning",
		Type:  "Respone",
		Date:  time.Now().Format("2006-01-02 15-04-05"),
		Details: detailsLog{
			Message: message,
			NetReport: netReport{
				From:   r.RemoteAddr,
				Method: r.Method,
				Url:    r.RequestURI,
				Status: 404,
				Time:   t,
			},
		},
	}

	l.data = append(l.data, logData)
	encodeData(l.data, l.path)
}

func (l *Logger) Network500(r *http.Request, t time.Duration, message string) {
	logData := dataLogger{
		Id:    uuid.NewString(),
		Level: "Error",
		Type:  "Respone",
		Date:  time.Now().Format("2006-01-02 15-04-05"),
		Details: detailsLog{
			Message: message,
			NetReport: netReport{
				From:   r.RemoteAddr,
				Method: r.Method,
				Url:    r.RequestURI,
				Status: 500,
				Time:   t,
			},
		},
	}

	l.data = append(l.data, logData)
	encodeData(l.data, l.path)
}

func (l *Logger) NetworkStatus(r *http.Request, t time.Duration, status int, level, _type, message string) {
	logData := dataLogger{
		Id:    uuid.NewString(),
		Level: level,
		Type:  _type,
		Date:  time.Now().Format("2006-01-02 15-04-05"),
		Details: detailsLog{
			Message: message,
			NetReport: netReport{
				From:   r.RemoteAddr,
				Method: r.Method,
				Url:    r.RequestURI,
				Status: status,
				Time:   t,
			},
		},
	}

	l.data = append(l.data, logData)
	encodeData(l.data, l.path)
}

func (l *Logger) Backup() {
	for {

		now := time.Now().Format("2006-01-02_15-04-05")
		nameLog := now + "-" + l.name
		_, err := openLogFile(parserPath(l.pathOrigin, nameLog), flagCreate)

		if err != nil {
			log.Println("Failed backup>", err.Error())
			return
		}

		l.path = parserPath(l.pathOrigin, nameLog)

		time.Sleep(l.backupTime)
	}
}

func encodeData(data []dataLogger, path string) {
	file, err := openLogFile(path, flagOpen)

	if err != nil {
		log.Println(err.Error())
		return
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		log.Println(err.Error())
		return
	}

	err = file.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func openLogFile(path string, flag int) (file *os.File, err error) {
	file, err = os.OpenFile(path, flag, 0766)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func parserPath(path, name string) string {
	if path == "./" {
		return path + name + ext
	}
	return path + "/" + name + ext
}
