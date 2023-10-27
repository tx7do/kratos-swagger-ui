package swaggerUI

import (
	"net/http"

	"github.com/tx7do/kratos-swagger-ui/internal"
)

// Handler handles swagger UI request.
type Handler = internal.Handler

// New creates HTTP handler for Swagger UI.
func New(title, swaggerJSONPath string, basePath string) http.Handler {
	return newHandler(title, swaggerJSONPath, basePath)
}

// NewWithOption creates configurable handler constructor.
func NewWithOption(handlerOpts ...HandlerOption) http.Handler {
	opts := internal.NewConfig()

	for _, o := range handlerOpts {
		o(opts)
	}

	return newHandlerWithConfig(opts)
}

// newHandlerWithConfig creates HTTP handler for Swagger UI.
func newHandlerWithConfig(config *internal.Config) *Handler {
	return internal.NewHandlerWithConfig(config, assetsBase, faviconBase, staticServer)
}

// NewHandler creates HTTP handler for Swagger UI.
func newHandler(title, swaggerJSONPath string, basePath string) *Handler {
	return newHandlerWithConfig(&internal.Config{
		Title:       title,
		SwaggerJSON: swaggerJSONPath,
		BasePath:    basePath,
	})
}

type httpServerInterface interface {
	HandlePrefix(prefix string, h http.Handler)
}

func RegisterSwaggerUIServer[T httpServerInterface](srv T, title, swaggerJSONPath string, basePath string) {
	swaggerHandler := newHandler(title, swaggerJSONPath, basePath)
	srv.HandlePrefix(swaggerHandler.BasePath, swaggerHandler)
}

func RegisterSwaggerUIServerWithOption[T httpServerInterface](srv T, handlerOpts ...HandlerOption) {
	opts := internal.NewConfig()

	for _, o := range handlerOpts {
		o(opts)
	}

	swaggerHandler := newHandlerWithConfig(opts)

	srv.HandlePrefix(swaggerHandler.BasePath, swaggerHandler)
}
