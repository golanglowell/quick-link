package resthttp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golanglowell/quick-link/internal/application"
	"github.com/golanglowell/quick-link/pkg/logger"
)

func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func handleShortenLink(logger *logger.Logger, uc *application.ShortenURL) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			logger.Error("Method not allowed",
				"method", r.Method,
				"path", r.URL.Path)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Invalid request body",
				"error", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		link, err := uc.Execute(req.URL)
		if err != nil {
			logger.Error("Failed to shorten link",
				"error", err)
			http.Error(w, "Failed to shorten link", http.StatusInternalServerError)
			return
		}

		resp := struct {
			ShortURL string `json:"short_url"`
		}{
			ShortURL: fmt.Sprintf("http://%s/%s", r.Host, link.ShortCode),
		}
		logger.Info("Link shortened successfully",
			"long_url", req.URL,
			"short_code", link.ShortCode)
		encode(w, r, http.StatusCreated, resp)
	})
}

func handleGetLink(logger *logger.Logger, uc *application.GetLinkUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			logger.Error("Method not allowed",
				"method", r.Method,
				"path", r.URL.Path)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		shortCode := r.URL.Path[1:]
		link, err := uc.Execute(shortCode)
		if err != nil {
			logger.Error("Link not found",
				"short_code", shortCode)
			http.Error(w, "Link not found", http.StatusNotFound)
			return
		}

		logger.Info("Redirecting",
			"short_code", shortCode,
			"long_url", link.LongURL)
		http.Redirect(w, r, link.LongURL, http.StatusMovedPermanently)
	})
}

func handleHealthz(logger *logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Health check", "status", "OK")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
