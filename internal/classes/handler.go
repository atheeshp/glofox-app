package classes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/atheeshp/glofox-app/utils"
	"github.com/gin-gonic/gin"
)

func CreateClass(c *gin.Context) {
	var parsedBody, ok = c.Get("parsedBody")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error reading parsed body"})
		c.Abort()
		return
	}

	req, ok := parsedBody.(reqClass)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error reading data"})
		c.Abort()
		return
	}

	startDate, _ := time.Parse(utils.DateFormat, req.StartDate)
	EndDate, _ := time.Parse(utils.DateFormat, req.EndDate)

	class := NewClass(req.Name, startDate, EndDate, req.Capacity)

	id := cs.AddClass(class)
	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("You're class: %d created", id),
	})
}
