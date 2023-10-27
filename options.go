package swaggerUI

import (
	"github.com/tx7do/kratos-swagger-ui/internal"
)

type HandlerOption func(opt *internal.Config)

func WithTitle(title string) HandlerOption {
	return func(opt *internal.Config) {
		opt.Title = title
	}
}

func WithSwaggerJSON(filePath string) HandlerOption {
	return func(opt *internal.Config) {
		opt.SwaggerJSON = filePath
	}
}

func WithBasePath(path string) HandlerOption {
	return func(opt *internal.Config) {
		opt.BasePath = path
	}
}
