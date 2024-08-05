package main

import (
    logstash_logger "github.com/KaranJagtiani/go-logstash"
	"fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    logger := logstash_logger.Init("logstash", 5400, "tcp", 5)

    payload := map[string]interface{}{
        "message": "TEST_MSG",
        "error":   false,
    }

    logger.Log(payload) // Generic log
    logger.Info(payload) // Adds "severity": "INFO"
    logger.Debug(payload) // Adds "severity": "DEBUG"
    logger.Warn(payload) // Adds "severity": "WARN"
    logger.Error(payload) // Adds "severity": "ERROR"


	router := mux.NewRouter()

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		payload := map[string]interface{}{
			"message": "/ requested",
			"error":   false,
		}

		logger.Info(payload)

        fmt.Fprint(w, "Hello, World!")
    })


	router.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {

		payload := map[string]interface{}{
			"message": "/error requested",
			"error":   true,
		}

		logger.Error(payload)

        fmt.Fprint(w, "Error!")
    })

    http.ListenAndServe(":8000", router)
}