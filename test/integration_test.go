package test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atheeshp/glofox-app/internal/bookings"
	"github.com/atheeshp/glofox-app/internal/classes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Setup a test Gin router with middleware and handler
func setupTestServer() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	api.POST("/classes", classes.ValidateCreateClass, classes.CreateClass)
	api.POST("/bookings", bookings.ValidateBooking, bookings.CreateBooking)
	return router
}

// TestBookingIntegration
func TestBookingIntegration(t *testing.T) {
	router := setupTestServer()

	tests := []struct {
		name       string
		payload    string
		expectCode int
		expectMsg  string
	}{
		{
			name: "Valid Booking",
			payload: `{
                "member": "John Doe",
                "date": "2024-06-10"
            }`,
			expectCode: http.StatusCreated,
			expectMsg:  "You're booking: 1 added",
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
			req := httptest.NewRequest("POST", "/api/bookings", bytes.NewBuffer([]byte(tt.payload)))
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectCode, w.Code)

			if tt.expectCode == http.StatusCreated {
				assert.Contains(t, w.Body.String(), tt.expectMsg)
			} else {
				assert.JSONEq(t, fmt.Sprintf(`{"error": "%s"}`, tt.expectMsg), w.Body.String())
			}
		})
	}
}

// TestClassIntegration
func TestClassIntegration(t *testing.T) {
	router := setupTestServer()

	tests := []struct {
		name       string
		payload    string
		expectCode int
		expectMsg  string
	}{
		{
			name: "Valid Class Creation",
			payload: `{
                "name": "Pilates",
                "start_date": "2024-06-01",
                "end_date": "2024-06-10",
                "capacity": 20
            }`,
			expectCode: http.StatusCreated,
			expectMsg:  "You're class: 1 created",
		},
		{
			name: "Invalid JSON Format",
			payload: `{
                "name": "Yoga",
                "start_date": "2024-06-01",
                "end_date": 2024-06-10,
                "capacity": 10
            }`,
			expectCode: http.StatusBadRequest,
			expectMsg:  "Invalid request format",
		},
		{
			name: "Invalid Start Date Format",
			payload: `{
                "name": "Cardio",
                "start_date": "22-01-2001",
                "end_date": "2024-06-10",
                "capacity": 10
            }`,
			expectCode: http.StatusBadRequest,
			expectMsg:  "start date: 22-01-2001 is not in the format of: 2006-01-02",
		},
		{
			name: "Invalid End Date Format",
			payload: `{
                "name": "CrossFit",
                "start_date": "2024-06-01",
                "end_date": "22-01-2001",
                "capacity": 15
            }`,
			expectCode: http.StatusBadRequest,
			expectMsg:  "end date: 22-01-2001 is not in the format of: 2006-01-02",
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
                "name": "HIIT",
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
			req := httptest.NewRequest("POST", "/api/classes", bytes.NewBuffer([]byte(tt.payload)))
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectCode, w.Code)

			if tt.expectCode == http.StatusCreated {
				assert.Contains(t, w.Body.String(), tt.expectMsg)
			} else {
				assert.JSONEq(t, fmt.Sprintf(`{"error": "%s"}`, tt.expectMsg), w.Body.String())
			}
		})
	}
}
