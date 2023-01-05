package main

import (
	"fmt"
	"logger-backup/logger"
	"net/http"
	"time"
)

func main() {
	var ServerLogger = logger.New("./logs", "rest", 15*time.Second)

	go ServerLogger.Backup()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ServerLogger.Info("Message", "database")
		ServerLogger.Error("Message", "database")
		ServerLogger.Warning("Message", "system")
		ServerLogger.Network500(r, 1, "Message")
		ServerLogger.Network404(r, 1, "Message")
		ServerLogger.Network200(r, 1, "Message")
		ServerLogger.NetworkStatus(r, 1, 204, "info", "response", "message")
		fmt.Fprintf(w, "Hola")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
