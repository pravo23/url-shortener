package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pravo23/url-shortener/handler"
	"github.com/pravo23/url-shortener/store"
)

func main() {

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "The URL Shortener Service!",
		})
	})

	r.POST("/shorten", func(ctx *gin.Context) {
		handler.CreateShortUrl(ctx)
	})

	r.GET("/:shortUrl", func(ctx *gin.Context) {
		handler.ShortUrlRedirect(ctx)
	})

	store.InitializeStore()

	err := r.Run(":9808")

	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
