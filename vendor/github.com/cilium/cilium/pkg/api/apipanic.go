// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package api

import (
	"errors"
	"net/http"
	"os"
	"runtime/debug"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/cilium/cilium/pkg/logging"
)

// APIPanicHandler recovers from API panics and logs encountered panics
type APIPanicHandler struct {
	Next http.Handler
}

// ServeHTTP implements the http.Handler interface.
// It recovers from panics of all next handlers and logs them
func (h *APIPanicHandler) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fields := logrus.Fields{
				"url":    req.URL.String(),
				"method": req.Method,
				"client": req.RemoteAddr,
			}

			if err, ok := r.(error); ok && errors.Is(err, syscall.EPIPE) {
				log.WithError(err).WithFields(fields).Debug("Failed to write API response: client connection closed")
				return
			}

			log.WithFields(fields).WithField("panic_message", r).Warn("Cilium API handler panicked")
			if logging.DefaultLogger.IsLevelEnabled(logrus.DebugLevel) {
				os.Stdout.Write(debug.Stack())
			}
			wr.WriteHeader(http.StatusInternalServerError)
			if _, err := wr.Write([]byte("Internal error occurred, check Cilium logs for details.")); err != nil {
				log.WithError(err).Debug("Failed to write API response")
			}
		}
	}()
	h.Next.ServeHTTP(wr, req)
}
