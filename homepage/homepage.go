package homepage

import (
	"log"
	"net/http"
	"time"
)

const message = "Hello from the other side"

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func (h *Handlers) LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		// A defer statement defers the execution of a function until the surrounding function returns.
		defer h.logger.Println("Request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.LoggerMiddleware(h.Home))
}

// NewHandlers Creates a class
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
