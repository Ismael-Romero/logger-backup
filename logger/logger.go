package logger

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

const (
	ext        = ".json"
	flagCreate = os.O_CREATE | os.O_RDWR
	flagOpen   = os.O_RDWR
)

var payload []interface{}

type Logger struct {
	name       string
	path       string
	pathOrigin string
	backupTime time.Duration
	data       []interface{}
	mutex      *sync.RWMutex
}

func New(path, name string, time time.Duration, mtx *sync.RWMutex) *Logger {
	_, err := openLogFile(parserPath(path, name), flagCreate)

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return &Logger{
		name:       name,
		path:       parserPath(path, name),
		pathOrigin: path,
		backupTime: time,
		data:       payload,
		mutex:      mtx,
	}
}

func (l *Logger) Backup() {
	for {
		l.mutex.Lock()
		now := time.Now().Format("2006-01-02_15-04-05")
		nameLog := now + "-" + l.name
		_, err := openLogFile(parserPath(l.pathOrigin, nameLog), flagCreate)

		if err != nil {
			log.Println("Failed backup>", err.Error())
			return
		}

		l.path = parserPath(l.pathOrigin, nameLog)
		l.mutex.Unlock()
		time.Sleep(l.backupTime)
	}
}

func encoder(l *Logger, content interface{}) {
	l.mutex.Lock()
	l.data = append(l.data, content)
	file, err := openLogFile(l.path, flagOpen)

	if err != nil {
		log.Println(err.Error())
		return
	}

	_encoder := json.NewEncoder(file)
	err = _encoder.Encode(l.data)

	if err != nil {
		log.Println(err.Error())
		return
	}
	err = file.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}
	l.mutex.Unlock()
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
