package middleware

import "net/http"

type WrapResponseWriter struct {
	http.ResponseWriter
	status int
}

func NewWrapResponseWriter(w http.ResponseWriter) *WrapResponseWriter {
	return &WrapResponseWriter{
		ResponseWriter: w,
		status:         http.StatusOK,
	}
}

func (w *WrapResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *WrapResponseWriter) Status() int {
	return w.status
}
