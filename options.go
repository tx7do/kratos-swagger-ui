package swaggerUI

import (
	"github.com/tx7do/kratos-swagger-ui/internal/swagger"
)

type HandlerOption func(opt *swagger.Config)

// WithTitle Title of an index file.
func WithTitle(title string) HandlerOption {
	return func(opt *swagger.Config) {
		opt.Title = title
	}
}

// WithBasePath Base URL to docs.
func WithBasePath(path string) HandlerOption {
	return func(opt *swagger.Config) {
		opt.BasePath = path
	}
}

// WithShowTopBar Show navigation top bar, hidden by default.
func WithShowTopBar(show bool) HandlerOption {
	return func(opt *swagger.Config) {
		opt.ShowTopBar = show
	}
}

// WithHideCurl Hide curl code snippet
func WithHideCurl(hide bool) HandlerOption {
	return func(opt *swagger.Config) {
		opt.HideCurl = hide
	}
}

// WithJsonEditor Enable visual JSON editor support (experimental can fail with complex schemas).
func WithJsonEditor(enable bool) HandlerOption {
	return func(opt *swagger.Config) {
		opt.JsonEditor = enable
	}
}

// WithPreAuthorizeApiKey Map of security name to key value
func WithPreAuthorizeApiKey(keys map[string]string) HandlerOption {
	return func(opt *swagger.Config) {
		opt.PreAuthorizeApiKey = keys
	}
}

// WithSettingsUI contains keys and plain javascript values of SwaggerUIBundle configuration.
// Overrides default values.
// See https://swagger.io/docs/open-source-tools/swagger-ui/usage/configuration/ for available options.
func WithSettingsUI(settings map[string]string) HandlerOption {
	return func(opt *swagger.Config) {
		opt.SettingsUI = settings
	}
}

func WithLocalFile(filePath string) HandlerOption {
	return func(opt *swagger.Config) {
		opt.LocalOpenApiFile = filePath
	}
}

func WithMemoryData(content []byte, ext string) HandlerOption {
	return func(opt *swagger.Config) {
		opt.OpenApiData = content
		opt.OpenApiDataType = ext
	}
}

// WithRemoteFileURL URL to openapi.json/swagger.json document specification.
func WithRemoteFileURL(url string) HandlerOption {
	return func(opt *swagger.Config) {
		opt.SwaggerJsonUrl = url
	}
}
