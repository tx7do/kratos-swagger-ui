package swaggerUI

import (
	"io"
	"net/http"
	"os"
)

type openApiFileHandler struct {
	Content []byte
}

func (h *openApiFileHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	_, _ = writer.Write(h.Content)
}

func (h *openApiFileHandler) loadOpenApiFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	return content, err
}

func (h *openApiFileHandler) LoadFile(filePath string) error {
	content, err := h.loadOpenApiFile(filePath)
	if err != nil {
		return err
	}

	h.Content = content
	return nil
}
