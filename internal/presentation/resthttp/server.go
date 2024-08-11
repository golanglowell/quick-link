package resthttp

import (
	"net/http"

	"github.com/golanglowell/quick-link/internal/application"
	"github.com/golanglowell/quick-link/pkg/config"
	"github.com/golanglowell/quick-link/pkg/logger"
)

func NewServer(
	logger *logger.Logger,
	config *config.Config,
	shortenLinkUC *application.ShortenURL,
	getLinkUC *application.GetLinkUseCase,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		logger,
		shortenLinkUC,
		getLinkUC,
	)

	var handler http.Handler = mux
	handler = loggingMiddleware(logger, handler)

	return handler
}

func addRoutes(
	mux *http.ServeMux,
	logger *logger.Logger,
	shortenLinkUC *application.ShortenURL,
	getLinkUC *application.GetLinkUseCase,
) {
	mux.Handle("/shorten", handleShortenLink(logger, shortenLinkUC))
	mux.Handle("/", handleGetLink(logger, getLinkUC))
	mux.HandleFunc("/healthz", handleHealthz(logger))
}
