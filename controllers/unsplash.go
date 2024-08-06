package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type ImageUrls struct {
	Regular string `json:"regular"`
	Small   string `json:"small"`
}

type UnsplashResult struct {
	ID             string    `json:"id"`
	AltDescription string    `json:"alt_description"`
	Urls           ImageUrls `json:"urls"`
}

type UnsplashResponse struct {
	Total      int              `json:"total"`
	TotalPages int              `json:"total_pages"`
	Results    []UnsplashResult `binding:"dive" json:"results"`
}

func SearchPhotos(c *gin.Context) {
	body, err := unsplashRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var response UnsplashResponse
	if err := json.Unmarshal(body, &response); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func buildUri(c *gin.Context) (uri string) {
	var base string = "https://api.unsplash.com/search/photos"
	v := url.Values{}
	v.Add("query", c.Query("search"))
	v.Add("page", c.Query("page"))
	query := v.Encode()
	uri = fmt.Sprintf("%s?%s", base, query)
	fmt.Println(uri)

	return
}

func unsplashRequest(c *gin.Context) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", buildUri(c), nil)
	if err != nil {
		return nil, err

	}
	req.Header.Add("Authorization", "Client-ID LQk7MJuPVUOrDAywuSgfUbAEf72vg8Wxig3XFtlUfJs")
	res, err := client.Do(req)
	if err != nil {
		return nil, err

	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	return body, err
}
