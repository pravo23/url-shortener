package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pravo23/url-shortener/shortener"
	"github.com/pravo23/url-shortener/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(ctx *gin.Context) {
	var request UrlCreationRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(request.LongUrl, request.UserId)
	store.SaveUrlMapping(shortUrl, request.LongUrl, request.UserId)

	host := "http://localhost:9808/"
	ctx.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func ShortUrlRedirect(ctx *gin.Context) {

	shortUrl := ctx.Param("shortUrl")
	originalUrl := store.RetrieveOriginalUrl(shortUrl)
	ctx.Redirect(302, originalUrl)

}
