package main

import (
	"fmt"
	"logger-backup/logger"
	"net/http"
	"sync"
	"time"
)

// This code includes an instance of a logger that runs on a separate goroutine.
// This logger has a backup function that runs every 10 seconds and writes the
// log messages to a file in the "./logs" directory with the name "log".
// The logger has three methods, Info, Warning and Error, which write messages with
// different severity levels to the log files.
//
// You can add new log levels in the "recorders" file of the logger package,
// as well as customize the data you want to save through the "logger" model.
//
// The code also includes an instance of sync.RWMutex, which is a read/write
// synchronization structure used to protect concurrent access to shared resources.
// In this case, the mutex is passed to the logger constructor and is used to protect
// access to the log file while messages are being written.

func main() {

	var mtx = &sync.RWMutex{}
	const _time = 15 * time.Second // Specifies the waiting time for each backup
	const _path = "./logs"         // Specifies the path to the directory where the records will be saved.
	const _name = "log"            // Specifies the file name without extension

	loggerSystem := logger.New(_path, _name, _time, mtx) // New logger is created
	go loggerSystem.Backup()                             // Activate backups

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		name := r.URL.Query().Get("name")

		loggerSystem.Info(name + "logged in")
		loggerSystem.Warning(name + "logged in without password")
		loggerSystem.Error("Failed application")

		_, _ = fmt.Fprintf(w, "Hello %s", name)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
