package main

import (
	"fmt"

	"gofr.dev/pkg/gofr"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	app := gofr.New()

	app.GET("/", func(ctx *gofr.Context) (interface{}, error) {
		return "index", nil
	})

	app.GET("/user", func(ctx *gofr.Context) (interface{}, error) {
		// get all query param,
		// return ctx.Params(), nil

		// get query param
		return ctx.Param("name"), nil
	})

	// get path param
	app.GET("/user/{id}", func(ctx *gofr.Context) (interface{}, error) {
		return fmt.Sprintf("Hello %s", ctx.PathParam("id")), nil
	})

	// get body
	app.POST("/user/new", func(ctx *gofr.Context) (interface{}, error) {
		var u user
		if err := ctx.BindStrict(&u); err != nil {
			return nil, err
		}
		return u.Name, nil
	})

	// get Header
	app.POST("/header", func(ctx *gofr.Context) (interface{}, error) {
		return ctx.Header("Content-Type"), nil
	})

	app.Start()
}
