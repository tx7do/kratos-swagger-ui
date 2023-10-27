package main

import (
	swaggerUI "github.com/tx7do/kratos-swagger-ui"
	"net/http"
)

func main() {
	//swaggerHandler := swaggerUI.New(
	//	"Petstore",
	//	"https://petstore3.swagger.io/api/v3/openapi.json",
	//	"/docs/",
	//)

	swaggerHandler := swaggerUI.NewWithOption(
		swaggerUI.WithTitle("Petstore"),
		swaggerUI.WithRemoteFile("https://petstore3.swagger.io/api/v3/openapi.json"),
		swaggerUI.WithBasePath("/docs/"),
	)

	//swaggerHandler := swaggerUI.NewWithOption(
	//	swaggerUI.WithTitle("Petstore"),
	//	swaggerUI.WithLocalOpenApiFile(".\\openapi.yaml"),
	//	swaggerUI.WithBasePath("/docs/"),
	//)

	http.Handle("/docs/", swaggerHandler)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello World!"))
	})

	http.HandleFunc("/docs/openapi.yml", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello World!"))
	})

	println("docs at http://localhost:8080/docs/")

	_ = http.ListenAndServe("localhost:8080", http.DefaultServeMux)
}
