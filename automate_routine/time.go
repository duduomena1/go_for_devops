package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func ReadFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Filed to read file: %s", err)
	}
	fmt.Printf("Contents of %s: \n%s\n", filename, string(data))
}

func WriteLog(logfile, message string) {
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	logger := log.New(file, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(message)
}

func MakeHTTPRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %s", err)
	}
	defer resp.Body.Close()

	log.Printf("Received response from %s: %s", url, resp.Status)
}

func main() {
	ReadFile("./data/example.txt")

	WriteLog("./data/logfile.log", "Routine task executed")

	ticker := time.NewTicker(2 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				WriteLog("./data/logfile.log", "Scheduled task executed")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	MakeHTTPRequest("https://httpbin.org/get")

	time.Sleep(1 * time.Minute)
	close(quit)
}
