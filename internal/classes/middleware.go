package classes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/atheeshp/glofox-app/utils"
	"github.com/gin-gonic/gin"
)

// ValidateCreateClass is a middleware which validates the class request
func ValidateCreateClass(c *gin.Context) {
	var req reqClass
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		c.Abort()
		return
	}

	startDate, err := time.Parse(utils.DateFormat, req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("start date: %s is not in the format of: %s", req.StartDate, utils.DateFormat)})
		c.Abort()
		return
	}

	EndDate, err := time.Parse(utils.DateFormat, req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("end date: %s is not in the format of: %s", req.EndDate, utils.DateFormat)})
		c.Abort()
		return
	}

	if startDate.After(EndDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start date shouldn't be after the end date"})
		c.Abort()
		return
	}

	if req.Capacity < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "class capacity should be at least 1"})
		c.Abort()
		return
	}

	c.Set("parsedBody", req)

	c.Next()
}
