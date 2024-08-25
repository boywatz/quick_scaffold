package main

import (
	"log"
	"net/http"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	r := router.New()
	r.GET("/", healthCheck)

	corsMiddleware := func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
			ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
			ctx.Response.Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

			// Handle preflight requests
			if string(ctx.Method()) == "OPTIONS" {
				ctx.SetStatusCode(http.StatusOK)
				return
			}

			next(ctx)
		}
	}

	handler := corsMiddleware(r.Handler)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(fasthttp.ListenAndServe(":8080", handler))
}

func healthCheck(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBodyString(`{"status": "OK", "message": "Service is healthy"}`)
}
