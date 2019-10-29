package apiserver

import (
	"context"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"strconv"
)


func (s *server) RequestIDMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := strconv.Itoa(rand.Int())
		ctx := r.Context()
		ctx = context.WithValue(ctx, "rIDKey", reqID)
		w.Header().Set("Request-ID", reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *server) AccessLogMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Info(r.URL.Path,
			zap.String("method:", r.Method),
			zap.String("remote_addr:", r.RemoteAddr),
			zap.String("url:", r.URL.Path),
		)
		next.ServeHTTP(w, r)
	})
}


func (s *server)CORSMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST,PUT,DELETE,GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Lol")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", s.clientUrl)
		if r.Method == http.MethodOptions{
			s.respond(w , r , http.StatusOK, nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}