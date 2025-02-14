package bookings

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test ValidateBooking Middleware
func TestValidateBooking(t *testing.T) {
	tests := []struct {
		name       string
		payload    string
		expectCode int
		expectMsg  string
	}{
		{
			name: "Valid Booking Request",
			payload: `{
                "member": "John Doe",
                "date": "2024-06-10"
            }`,
			expectCode: http.StatusOK,
			expectMsg:  "",
		},
		{
			name: "Invalid JSON Format",
			payload: `{
                "member": "John Doe",
                "date": 2024-06-10
            }`,
			expectCode: http.StatusBadRequest,
			expectMsg:  "Invalid request format",
		},
		{
			name: "Invalid Member Name",
			payload: `{
                "member": "12345",
                "date": "2024-06-10"
            }`,
			expectCode: http.StatusBadRequest,
			expectMsg:  "class name should contain only letters and spaces",
		},
		{
			name: "Invalid Date Format",
			payload: `{
                "member": "Jane Doe",
                "date": "invalid-date"
            }`,
			expectCode: http.StatusBadRequest,
			expectMsg:  "start date: invalid-date is not in the format of: 2006-01-02",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request = httptest.NewRequest("POST", "/api/bookings", bytes.NewBuffer([]byte(tt.payload)))
			c.Request.Header.Set("Content-Type", "application/json")

			ValidateBooking(c) // Call the middleware function

			// Verify response code
			assert.Equal(t, tt.expectCode, w.Code)

			// Verify error message (for failed cases)
			if tt.expectCode == http.StatusBadRequest {
				assert.JSONEq(t, fmt.Sprintf(`{"error": "%s"}`, tt.expectMsg), w.Body.String())
			}
		})
	}
}
