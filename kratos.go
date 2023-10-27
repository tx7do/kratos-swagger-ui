package swaggerUI

import (
	"github.com/tx7do/kratos-swagger-ui/internal/swagger"
	"io"
	"net/http"
	"os"
	"strings"
)

type openJsonFileHandler struct {
	content []byte
}

func (h *openJsonFileHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	_, _ = writer.Write(h.content)
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

	h.content = content
	return nil
}

type httpServerInterface interface {
	HandlePrefix(prefix string, h http.Handler)
}

func RegisterSwaggerUIServer[T httpServerInterface](srv T, title, swaggerJSONPath string, basePath string) {
	swaggerHandler := newHandler(title, swaggerJSONPath, basePath)

	if swaggerHandler.LocalOpenApiFile != "" {
		registerOpenApiFileRouter(srv, swaggerHandler)
	}

	srv.HandlePrefix(swaggerHandler.BasePath, swaggerHandler)
}

func RegisterSwaggerUIServerWithOption[T httpServerInterface](srv T, handlerOpts ...HandlerOption) {
	opts := swagger.NewConfig()

	for _, o := range handlerOpts {
		o(opts)
	}

	swaggerHandler := newHandlerWithConfig(opts)

	if swaggerHandler.LocalOpenApiFile != "" {
		registerOpenApiFileRouter(srv, swaggerHandler)
	}

	srv.HandlePrefix(swaggerHandler.BasePath, swaggerHandler)
}

func registerOpenApiFileRouter[T httpServerInterface](srv T, swaggerHandler *Handler) {
	_openJsonFileHandler := openJsonFileHandler{}
	err := _openJsonFileHandler.LoadFile(swaggerHandler.LocalOpenApiFile)
	if err == nil {
		pattern := strings.TrimRight(swaggerHandler.BasePath, "/") + "/openapi.json"
		srv.HandlePrefix(pattern, &_openJsonFileHandler)
		swaggerHandler.SwaggerJSON = pattern
	}
}
