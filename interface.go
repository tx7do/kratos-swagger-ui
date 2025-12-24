package swaggerUI

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/tx7do/kratos-swagger-ui/internal/swagger"
)

type httpServerInterface interface {
	HandlePrefix(prefix string, h http.Handler)
	Handle(path string, h http.Handler)
	HandleFunc(path string, h http.HandlerFunc)
}

func RegisterSwaggerUIServer[T httpServerInterface](srv T, title, swaggerJSONPath string, basePath string) {
	swaggerHandler := newHandler(title, swaggerJSONPath, basePath)
	srv.HandlePrefix(swaggerHandler.BasePath, swaggerHandler)
}

func RegisterSwaggerUIServerWithOption[T httpServerInterface](srv T, handlerOpts ...HandlerOption) {
	opts := swagger.NewConfig()

	for _, o := range handlerOpts {
		o(opts)
	}

	if opts.LocalOpenApiFile != "" {
		registerOpenApiLocalFileRouter(srv, opts)
	} else if len(opts.OpenApiData) != 0 {
		registerOpenApiMemoryDataRouter(srv, opts)
	}

	swaggerHandler := newHandlerWithConfig(opts)

	srv.HandlePrefix(swaggerHandler.BasePath, swaggerHandler)
}

// var _openJsonFileHandler = &openApiFileHandler{}

func registerOpenApiLocalFileRouter[T httpServerInterface](srv T, cfg *swagger.Config) {
	var _openJsonFileHandler = &openApiFileHandler{}
	err := _openJsonFileHandler.LoadFile(cfg.LocalOpenApiFile)
	if err == nil {
		pattern := strings.TrimRight(cfg.BasePath, "/") + "/openapi" + path.Ext(cfg.LocalOpenApiFile)
		cfg.SwaggerJsonUrl = pattern
		srv.Handle(pattern, _openJsonFileHandler)
	} else {
		fmt.Println("load openapi file failed: ", err)
	}
}

func registerOpenApiMemoryDataRouter[T httpServerInterface](srv T, cfg *swagger.Config) {
	var _openJsonFileHandler = &openApiFileHandler{}
	_openJsonFileHandler.Content = cfg.OpenApiData
	pattern := strings.TrimRight(cfg.BasePath, "/") + "/openapi." + cfg.OpenApiDataType
	cfg.SwaggerJsonUrl = pattern
	srv.Handle(pattern, _openJsonFileHandler)
	cfg.OpenApiData = nil
}
