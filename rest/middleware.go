package rest

import (
	"log"
	"net/http"
	"strings"
	"time"

	b "github.com/betonr/go-utils/base"
)

// DefaultResponse - struct default to simple response
type DefaultResponse struct {
	Status  int
	Message string
}

// Cors - struct of list to configurate CORS at application
type Cors struct {
	Methods []string
	Origins []string
}

// EnableCors - enable cors at application with base at configurate passed by params
func EnableCors(handler http.Handler, c Cors) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(c.Methods, ","))
		w.Header().Set("Access-Control-Allow-Origin", strings.Join(c.Origins, ","))

		// validation of method (request)
		if !b.ContainStr(c.Methods, r.Method) {
			err := DefaultResponse{http.StatusMethodNotAllowed, "Method not allowed!"}
			RespondWithJson(w, http.StatusMethodNotAllowed, err)
			return
		}

		// validation of origin (request)
		// TODO:
		originRequest := ""
		if len(c.Origins) > 0 && c.Origins[0] != "*" && !b.ContainStr(c.Origins, originRequest) {
			err := DefaultResponse{http.StatusForbidden, "Origin not allowed!"}
			RespondWithJson(w, http.StatusForbidden, err)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

// EnableLogs - add logs at all request
func EnableLogs(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s - %v\n", r.RemoteAddr, r.Method, r.URL, time.Since(start))
		handler.ServeHTTP(w, r)
	})
}
