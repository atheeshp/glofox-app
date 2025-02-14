package classes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Reset the in-memory store before each test
func resetClassStore() {
	cs = ClassStore{} // Clear in-memory storage
}

// Test CreateClass handler without using mocks
func TestCreateClass(t *testing.T) {
	tests := []struct {
		name       string
		parsedBody any // Allows testing incorrect values
		expectCode int
		expectMsg  string
	}{
		{
			name: "Successful Class Creation",
			parsedBody: reqClass{
				Name:      "Yoga",
				StartDate: "2024-06-01",
				EndDate:   "2024-06-10",
				Capacity:  20,
			},
			expectCode: http.StatusCreated,
			expectMsg:  "You're class: 1 created",
		},
		{
			name:       "Missing Parsed Body",
			parsedBody: nil, // Simulates missing parsedBody in context
			expectCode: http.StatusBadRequest,
			expectMsg:  "error reading parsed body",
		},
		{
			name:       "Invalid Parsed Body Type",
			parsedBody: "InvalidData", // Wrong data type
			expectCode: http.StatusBadRequest,
			expectMsg:  "error reading data",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetClassStore() // Reset before each test

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Store mock data in context
			if tt.parsedBody != nil {
				c.Set("parsedBody", tt.parsedBody)
			}

			// Call actual function
			CreateClass(c)

			// Validate response code
			assert.Equal(t, tt.expectCode, w.Code)

			// Validate response message
			if tt.expectCode == http.StatusCreated {
				assert.Contains(t, w.Body.String(), tt.expectMsg)
			} else {
				assert.JSONEq(t, fmt.Sprintf(`{"error": "%s"}`, tt.expectMsg), w.Body.String())
			}
		})
	}
}
