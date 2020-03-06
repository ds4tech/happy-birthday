package birthday

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func Logger(inner http.Handler, name string, file *os.File) http.Handler {
	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})
	//log.SetLevel(log.InfoLevel)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inner.ServeHTTP(w, r)

		log.WithFields(log.Fields{
			"uri":     r.RequestURI,
			"handler": r.Method,
			"name":    name,
		}).Info("Normal")

	})

	return handler
}
