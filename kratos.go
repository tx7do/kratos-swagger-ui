package swaggerUI

import (
	"fmt"
	"github.com/tx7do/kratos-swagger-ui/internal/swagger"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

type openJsonFileHandler struct {
	Content []byte
}

func (h *openJsonFileHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	_, _ = writer.Write(h.Content)
}

func (h *openJsonFileHandler) loadOpenApiFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	return content, err
}

func (h *openJsonFileHandler) LoadFile(filePath string) error {
	content, err := h.loadOpenApiFile(filePath)
	if err != nil {
		return err
	}

	h.Content = content
	return nil
}

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
		registerOpenApiFileRouter(srv, opts)
	} else if len(opts.OpenApiData) != 0 {
		registerOpenApiDataRouter(srv, opts)
	}

	swaggerHandler := newHandlerWithConfig(opts)

	srv.HandlePrefix(swaggerHandler.BasePath, swaggerHandler)
}

var _openJsonFileHandler = &openJsonFileHandler{}

func registerOpenApiFileRouter[T httpServerInterface](srv T, cfg *swagger.Config) {
	err := _openJsonFileHandler.LoadFile(cfg.LocalOpenApiFile)
	if err == nil {
		pattern := strings.TrimRight(cfg.BasePath, "/") + "/openapi" + path.Ext(cfg.LocalOpenApiFile)
		cfg.SwaggerJSON = pattern
		srv.Handle(pattern, _openJsonFileHandler)
	} else {
		fmt.Println("load openapi file failed: ", err)
	}
}

func registerOpenApiDataRouter[T httpServerInterface](srv T, cfg *swagger.Config) {
	_openJsonFileHandler.Content = cfg.OpenApiData
	pattern := strings.TrimRight(cfg.BasePath, "/") + "/openapi." + cfg.OpenApiDataType
	cfg.SwaggerJSON = pattern
	srv.Handle(pattern, _openJsonFileHandler)
}
