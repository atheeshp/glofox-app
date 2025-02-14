package bookings

import (
	"fmt"
	"net/http"
	"time"

	"github.com/atheeshp/glofox-app/utils"
	"github.com/gin-gonic/gin"
)

// CreateBooking is a handler method for creating bookings
func CreateBooking(c *gin.Context) {
	parsedBody, ok := c.Get("parsedBody")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error reading parsed body"})
		return
	}

	req, ok := parsedBody.(reqBooking)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error reading data"})
		return
	}

	date, _ := time.Parse(utils.DateFormat, req.Date)

	booking := NewBooking(req.Member, date)
	id := bs.AddBooking(booking)

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("You're booking: %d added", id),
	})
}
