package helpers

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func InitializeLogger() {
	config, _ := LoadEnv(".")
	logDir := config.LogDir
	_ = os.Mkdir(logDir, os.ModePerm)

	f, err := os.OpenFile(logDir+config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		LogEvent("ERROR", "creating log file"+err.Error())
		log.Fatalf("error opening file: %v", err)
	}
	log.SetFlags(0)
	log.SetOutput(f)
}

// level can be INFO or ERROR
func LogEvent(level string, message interface{}) {

	data, err := json.Marshal(struct {
		TimeStamp string      `json:"@timestamp"`
		Level     string      `json:"level"`
		Message   interface{} `json:"message"`
	}{
		TimeStamp: time.Now().Format(time.RFC3339),
		Message:   message,
		Level:     level,
	})

	if err != nil {
		log.Fatal("Logevent helper: " + err.Error())
	}
	log.Printf("%s\n", data)

}
