// Package api is the Runphase HTTP process.
// It does not open a database or a Temporal client.
package api

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/mhomaid/runphase/internal/config"
)

const (
	// Control-plane requests are small JSON documents, not token streams.
	// Header timeout limits a client that connects and withholds headers.
	readHeaderTimeout = 5 * time.Second
	readTimeout       = 10 * time.Second
	writeTimeout      = 10 * time.Second
	idleTimeout       = 60 * time.Second
	shutdownTimeout   = 10 * time.Second
)

// NewServer builds an HTTP server for the process listen address.
func NewServer(cfg config.Config) *http.Server {
	return &http.Server{
		Addr:              cfg.HTTP.Addr,
		Handler:           NewHandler(),
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
	}
}

// NewHandler serves process health only.
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", writeOK)
	mux.HandleFunc("GET /readyz", writeOK)
	return mux
}

// Run listens until ctx is cancelled, then drains in-flight requests.
func Run(ctx context.Context, srv *http.Server) error {
	slog.Info("api starting", "addr", srv.Addr)

	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.ListenAndServe()
	}()

	select {
	case err := <-errCh:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	case <-ctx.Done():
		slog.Info("shutdown requested")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		shutErr := srv.Shutdown(shutdownCtx)
		listenErr := <-errCh
		if shutErr != nil {
			return shutErr
		}
		if listenErr != nil && !errors.Is(listenErr, http.ErrServerClosed) {
			return listenErr
		}
		slog.Info("api stopped")
		return nil
	}
}

func writeOK(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"}); err != nil {
		slog.Error("write json response", "err", err)
	}
}
