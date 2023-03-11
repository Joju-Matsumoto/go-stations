package middleware

import (
	"log"
	"net/http"
)

type LogFunc func(r *http.Request, recover interface{})

type RecoveryMiddleware struct {
	LogFunc LogFunc
}

func (m *RecoveryMiddleware) Recovery(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				m.LogFunc(r, rec)
			}
		}()

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func NewRecovery() *RecoveryMiddleware {
	return &RecoveryMiddleware{
		LogFunc: func(r *http.Request, recover interface{}) {
			log.Printf("path: %v, recovered: %v\n", r.URL.Path, recover)
		},
	}
}
