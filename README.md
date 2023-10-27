# kratos-swagger-ui

## How To Use

at the first, you need install this lib:

```shell
go get -u github.com/tx7do/kratos-swagger-ui
```

direct use:

```go
package main

import (
	"net/http"

	swaggerUI "github.com/tx7do/kratos-swagger-ui"
)

func main() {
	swaggerHandler := swaggerUI.New(
		"Petstore",
		"https://petstore3.swagger.io/api/v3/openapi.json",
		"/docs/",
	)
	
	http.Handle("/docs/", swaggerHandler)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello World!"))
	})

	println("docs at http://localhost:8080/docs/")

	_ = http.ListenAndServe("localhost:8080", http.DefaultServeMux)
}
```

use it in kratos:

```go
import (
    "net/http"
    
    swaggerUI "github.com/tx7do/kratos-swagger-ui"
    rest "github.com/go-kratos/kratos/v2/transport/http"
)

func NewRESTServer() *rest.Server {
    swaggerHandler := swaggerUI.New(
        "Petstore",
        "https://petstore3.swagger.io/api/v3/openapi.json",
        "/docs/",
    )
    srv.HandlePrefix("/docs/", swaggerHandler)
}
```

or

```go
import (
    "net/http"
    
    swaggerUI "github.com/tx7do/kratos-swagger-ui"
    rest "github.com/go-kratos/kratos/v2/transport/http"
)

func NewRESTServer() *rest.Server {
    swaggerUI.RegisterSwaggerUIServer(
        srv,
        "Petstore",
        "https://petstore3.swagger.io/api/v3/openapi.json",
        "/docs/",
    )
}
```

## References

- [Serve SwaggerUI within your Golang application](https://ribice.medium.com/serve-swaggerui-within-your-golang-application-5486748a5ed4)
- [go-kratos swagger-api](https://github.com/go-kratos/swagger-api)
- [swagger-ui - github](https://github.com/swagger-api/swagger-ui)
- [Swagger Open API Specification 2.0 and 3.0 in Go](https://kecci.medium.com/swagger-open-api-specification-2-0-and-3-0-in-go-c1f05b51a595)
- [Embedded Swagger UI for Go](https://github.com/swaggest/swgui)
- [Tutorial: Developing a RESTful API with Go, JSON Schema validation and OpenAPI docs](https://dev.to/vearutop/tutorial-developing-a-restful-api-with-go-json-schema-validation-and-openapi-docs-2490)
