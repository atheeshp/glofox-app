package classes

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test ValidateCreateClass Middleware
func TestValidateCreateClass(t *testing.T) {
	tests := []struct {
		name       string
		payload    string
		expectCode int
		expectMsg  string
	}{
		{
			name: "Valid Request",
			payload: `{
                "name": "Pilates",
                "start_date": "2024-06-01",
                "end_date": "2024-06-10",
                "capacity": 20
            }`,
			expectCode: http.StatusOK, // Middleware should pass to next handler
			expectMsg:  "",
		},
		{
			name: "invalid Request",
			payload: `{
                "name1": "Pilates",
            }`,
			expectCode: http.StatusBadRequest, // Middleware should pass to next handler
			expectMsg:  "Invalid request format",
		},
		{
			name: "Invalid Start Date",
			payload: `{
                "name": "Yoga",
                "start_date": "10-02-2025",
                "end_date": "2024-06-10",
                "capacity": 10
            }`,
			expectCode: http.StatusBadRequest,
			expectMsg:  "start date: 10-02-2025 is not in the format of: 2006-01-02",
		},
		{
			name: "Invalid End Date",
			payload: `{
                "name": "Yoga",
                "start_date": "2024-06-10",
                "end_date": "10-02-2025",
                "capacity": 10
            }`,
			expectCode: http.StatusBadRequest,
			expectMsg:  "end date: 10-02-2025 is not in the format of: 2006-01-02",
		},
		{
			name: "Start Date After End Date",
			payload: `{
                "name": "Strength Training",
                "start_date": "2024-06-10",
                "end_date": "2024-06-01",
                "capacity": 10
            }`,
			expectCode: http.StatusBadRequest,
			expectMsg:  "start date shouldn't be after the end date",
		},
		{
			name: "Invalid Capacity",
			payload: `{
                "name": "CrossFit",
                "start_date": "2024-06-01",
                "end_date": "2024-06-10",
                "capacity": 0
            }`,
			expectCode: http.StatusBadRequest,
			expectMsg:  "class capacity should be at least 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// fake buffer data
			c.Request = httptest.NewRequest("POST", "/api/classes", bytes.NewBuffer([]byte(tt.payload)))
			c.Request.Header.Set("Content-Type", "application/json")

			ValidateCreateClass(c)

			assert.Equal(t, tt.expectCode, w.Code)

			if tt.expectCode == http.StatusBadRequest {
				assert.JSONEq(t, fmt.Sprintf(`{"error": "%s"}`, tt.expectMsg), w.Body.String())
			}
		})
	}
}
