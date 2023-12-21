package main

import (
    "fmt"
    "gofr.dev/pkg/gofr"
)

func main() {
    app := gofr.New()

    app.GET("/", func(ctx *gofr.Context) (interface{}, error) {
        return "index", nil
    })

    // get all query param,
    app.GET("/all", func(ctx *gofr.Context) (interface{}, error) {
        return ctx.Params(), nil
    })

    // get query param
    app.GET("/user", func(ctx *gofr.Context) (interface{}, error) {
        return fmt.Sprintf("Hello %s", ctx.Param("name")), nil
    })

    // get path param
    app.GET("/user/{id}", func(ctx *gofr.Context) (interface{}, error) {
        return fmt.Sprintf("Hello %s", ctx.PathParam("id")), nil
    })

    app.Start()
}
