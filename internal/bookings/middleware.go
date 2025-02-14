package bookings

import (
	"fmt"
	"net/http"
	"time"

	"github.com/atheeshp/glofox-app/utils"
	"github.com/gin-gonic/gin"
)

// ValidateCreateClass is a middleware which validates the class request
func ValidateBooking(c *gin.Context) {
	var req reqBooking
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		c.Abort()
		return
	}

	if err := utils.ValidateClassName(req.Member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	_, err := time.Parse(utils.DateFormat, req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("start date: %s is not in the format of: %s", req.Date, utils.DateFormat)})
		c.Abort()
		return
	}

	c.Set("parsedBody", req)

	c.Next()
}
