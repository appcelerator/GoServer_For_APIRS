package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var path = "/ctlog/requests.log"

// Create a log entry in /ctlog/requests.log
func writeLogFile(requestlog RequestLog) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	check(err)
	defer file.Close()

	logstring, err := json.Marshal(requestlog)
	check(err)
	// write some text line-by-line to file
	_, err = file.Write(logstring)
	check(err)
	_, err = file.WriteString("\n")
	check(err)

	// save changes
	err = file.Sync()
	check(err)
}

type Req struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type Res struct {
	StatusCode int `json:"statusCode"`
}

type RequestLog struct {
	Time         time.Time `json:"time"`
	ResponseTime int       `json:"response_time"`
	Req          Req       `json:"req"`
	Res          Res       `json:"res"`
}

func main() {

	// arrowPing.json is used by Axway API Runtime Services to determine health of an application. 
	// It is not recommended to log this request since Axway API Runtime services makes frequent calls to this endpoint
	http.HandleFunc("/arrowPing.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{success: true}")
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi from goserver")
		headers := map[string]string{
			"user-agent":      "curl/7.29.0",
			"accept":          "*/*",
			"host":            "goserver.cloud.appctest.com", 
			"x-forwarded-for": "54.214.34.215", 	// Provide client ip here
		}
		req := Req{Method: "GET", Url: "/hi", Headers: headers}
		res := Res{StatusCode: 200}
		accesslog := RequestLog{Time: time.Now(), ResponseTime: 2, Req: req, Res: res}
		writeLogFile(accesslog)

		fmt.Println("==> Hi")
	})

	log.Println("Listen to port 80...")

	log.Fatal(http.ListenAndServe(":80", nil))

}
