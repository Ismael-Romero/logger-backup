# ⚙️Logger Backup

> ### version: 1.0.0


## ℹ️ How to use
Clone the repository to a directory of your choice and make a copy of the Logger folder to your project:
```
git clone https://github.com/Ismael-Romero/logger-backup.git
```

#### 💡 Create a new logger
Once the folder has been copied to your project directory, you will be able to create an instance of the logger through its constructor.
To make the instance, it is necessary to pass as arguments to the constructor; the **path** where the log file will be located, the **name of the file** without extension, the **time to wait** for backup copies and a **mutex**.

Example:
```go
	var mtx = &sync.RWMutex{}
	const _time = 15 * time.Second // Specifies the waiting time for each backup
	const _path = "./logs"         // Specifies the path to the directory where the records will be saved.
	const _name = "log"            // Specifies the file name without extension

	loggerSystem := logger.New(_path, _name, _time, mtx) // New logger is created
```

> Nota: Mutex is short for mutual exclusion. Mutexes keep track of which thread has access to a variable at any given time.

The effect of this code is immediate, so once instantiated you will be able to make use of the available methods of the logger.

#### 📝 Methods
By default, the logger backup package has 4 methods that will allow us to save messages according to the severity level, these are:
* **Info()**: records a message with the severity level of type Informational.
* **Error()**: records a message with the severity level of type Error.
* **Warning()**: records a message with the severity level of type Warning.
* **Response200()**: records a message with the severity level of information and the data of the request made by the client.

#### ⚙️ Enable backups
To enable backups, just call the Backup() method of the logger and prefix it with the word **go** so that it is executed in a new goroutine.
```go

	go loggerSystem.Backup()
```

#### Full example: 
```go

package main

import (
	"fmt"
	"logger-backup/logger"
	"net/http"
	"sync"
	"time"
)

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
		loggerSystem.Response200(r, "Welcome")

		_, _ = fmt.Fprintf(w, "Hello %s", name)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

```
## 🏷️ Default model
By default, the logger includes a data model that records basic information, you can modify the content according to the needs of your program and make their respective configurations in the "Recorders" file.

```go

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

```

#### Recorder example
```go
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

```
> Note: Do not forget to call the enconder() function since this is the one in charge of writing the data in the generated json file.

## 👀 Preview
### Backup
<img src="https://github.com/Ismael-Romero/logger-backup/blob/master/doc/backups.gif" />

### Output content
```json

[
    {
        "type": "System",
        "level": "info",
        "date": "2023-01-06 16-44-07",
        "body": {
            "message": "Daniellogged in"
        }
    },
    {
        "type": "System",
        "level": "Warning",
        "date": "2023-01-06 16-44-07",
        "body": {
            "message": "Daniellogged in without password"
        }
    },
    {
        "type": "System",
        "level": "Error",
        "date": "2023-01-06 16-44-07",
        "body": {
            "message": "Failed application"
        }
    },
    {
        "type": "System",
        "level": "info",
        "date": "2023-01-06 16-44-07",
        "body": {
            "from": "[::1]:61211",
            "protocol": "HTTP/1.1",
            "method": "GET",
            "URL": "/?name=Daniel",
            "status": 200
        }
    },
    ...
```
