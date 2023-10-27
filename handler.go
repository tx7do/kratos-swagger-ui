package swaggerUI

import (
	"github.com/tx7do/kratos-swagger-ui/internal"
	"net/http"
)

// Handler handles swagger UI request.
type Handler = internal.Handler

// NewWithConfig creates configurable handler constructor.
func NewWithConfig(config Config) func(title, swaggerJSONPath string, basePath string) http.Handler {
	return func(title, swaggerJSONPath string, basePath string) http.Handler {
		if config.Title == "" {
			config.Title = title
		}

		if config.SwaggerJSON == "" {
			config.SwaggerJSON = swaggerJSONPath
		}

		if config.BasePath == "" {
			config.BasePath = basePath
		}

		return NewHandlerWithConfig(config)
	}
}

// NewHandlerWithConfig creates HTTP handler for Swagger UI.
func NewHandlerWithConfig(config Config) *Handler {
	return internal.NewHandlerWithConfig(config, assetsBase, faviconBase, staticServer)
}

// New creates HTTP handler for Swagger UI.
func New(title, swaggerJSONPath string, basePath string) http.Handler {
	return NewHandler(title, swaggerJSONPath, basePath)
}

// NewHandler creates HTTP handler for Swagger UI.
func NewHandler(title, swaggerJSONPath string, basePath string) *Handler {
	return NewHandlerWithConfig(Config{
		Title:       title,
		SwaggerJSON: swaggerJSONPath,
		BasePath:    basePath,
	})
}
