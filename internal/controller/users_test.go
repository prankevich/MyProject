package controller_test

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_contracts "github.com/prankevich/MyProject/internal/contracts/mocks"
	"github.com/prankevich/MyProject/internal/controller"
)

func TestController_DeleteUserByID(t *testing.T) {
	type mockBehaviour func(s *mock_contracts.MockServiceI)

	testCases := []struct {
		name                 string
		paramID              string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:    "Valid ID - success",
			paramID: "42",
			mockBehaviour: func(s *mock_contracts.MockServiceI) {
				s.EXPECT().DeleteUsersByID(42).Return(nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"message":"Users deleted"}`,
		},
		{
			name:    "Invalid ID format",
			paramID: "abc",
			mockBehaviour: func(s *mock_contracts.MockServiceI) {

			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"strconv.Atoi: parsing \"abc\": invalid syntax"}`,
		},
		{
			name:    "Negative ID",
			paramID: "-5",
			mockBehaviour: func(s *mock_contracts.MockServiceI) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"user id must be positive"}`,
		},

		{
			name:    "Service error",
			paramID: "7",
			mockBehaviour: func(s *mock_contracts.MockServiceI) {
				s.EXPECT().DeleteUsersByID(7).Return(errors.New("database failure"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"error":"database failure"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Init mocks
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mock_contracts.NewMockServiceI(ctrl)
			tc.mockBehaviour(service)

			handler := controller.New(service)

			// Setup router
			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.DELETE("/users/:id", handler.DeleteUserByID)

			// Create request
			req := httptest.NewRequest(http.MethodDelete, "/users/"+tc.paramID, nil)
			w := httptest.NewRecorder()

			// Perform
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, tc.expectedStatusCode, w.Code)
			assert.Equal(t, tc.expectedResponseBody, w.Body.String())
		})
	}
}
