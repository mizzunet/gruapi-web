package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	gd "gruapi"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("static/*")
	// GETs
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/search", gruapi_search)
	r.GET("/view", gruapi_view)

	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.StaticFile("/style.css", "./static/style.css")

	r.Run()
	gin.SetMode(gin.ReleaseMode)
}

func gruapi_search(c *gin.Context) {
	Query := c.Query("q")

	if Query == "" {
		c.String(200, "Search for something")
	} else {
		Filter, _ := strconv.Atoi(c.Query("filter"))
		Count, _ := strconv.Atoi(c.Query("count"))
		BooksJSON := gd.Search(Query, Filter, Count)
		c.JSONP(200, BooksJSON)
	}
}

func gruapi_view(c *gin.Context) {
	Link := c.Query("link")

	if Link == "" {
		c.String(200, "Search for something")
	} else {
		BooksJSON := gd.View(Link)
		c.JSONP(200, BooksJSON)
	}
}
