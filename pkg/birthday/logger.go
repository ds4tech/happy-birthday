package birthday

import (
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func Logger(inner http.Handler, name string) http.Handler {
	// log to File
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		//			log.Fatal(err)
		log.WithFields(log.Fields{
			"omg":    true,
			"number": 100,
		}).Fatal("The ice breaks!")
	}
	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})
	log.Print("Logging to a file.")

	// end

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)

		defer file.Close()
	})

	return handler
}
