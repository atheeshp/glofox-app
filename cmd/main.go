package main

import (
	"net/http"

	"github.com/atheeshp/glofox-app/internal/bookings"
	"github.com/atheeshp/glofox-app/internal/classes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api := r.Group("/api")
	classes.ClassRouter(api.Group("/classes"))
	bookings.BookingsRouter(api.Group("/bookings"))

	http.ListenAndServe(":8080", r)
}
