package routes

import (
	"net/http"

	"github.com/bahodurnazarov/calculator/cors"
	"github.com/gin-gonic/gin"
)

func Listen() {
	router := gin.Default()
	router.Use(cors.CORSMiddleware())
	router.LoadHTMLGlob("assets/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
		})
	})
	router.GET("/second", func(c *gin.Context) {
		c.HTML(http.StatusOK, "second.html", gin.H{
		})
	})

	router.Run(":8181")
}
