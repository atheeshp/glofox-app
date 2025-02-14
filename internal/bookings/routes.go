package bookings

import "github.com/gin-gonic/gin"

func BookingsRouter(router *gin.RouterGroup) {
	router.POST("/", ValidateBooking, CreateBooking)
}
