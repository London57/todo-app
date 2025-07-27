package auth

import (
	"errors"
	"strings"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/London57/todo-app/config"
	"github.com/London57/todo-app/internal/controller/http/common/controller"
	"github.com/London57/todo-app/internal/domain"
	mock_signup "github.com/London57/todo-app/internal/domain/signup/mocks"
	"github.com/London57/todo-app/internal/transport/signup"
	mock_logger "github.com/London57/todo-app/pkg/logger/mocks"
	validate "github.com/London57/todo-app/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAuthController_signup(t *testing.T) {
	testcases := []struct {
		test_name          string
		inputBody          string
		inputRequest       signup.SignUpRequest
		mockBehavior       func(*mock_signup.MockSignUpUseCase, signup.SignUpRequest)
		expectedStatusCode int
	}{
		{
			test_name: "OK",
			inputBody: `{"name":"John", "username":"john123", "password":"123456", "email":"john@123.ru"}`,
			inputRequest: signup.SignUpRequest{
				Name:     "John",
				Username: "john123",
				Email:    "john@123.ru",
				Password: "123456",
			},
			mockBehavior: func(s *mock_signup.MockSignUpUseCase, signupRequest signup.SignUpRequest) {
				uuid := uuid.New()
				s.EXPECT().GetUserByEmail(gomock.Any(), signupRequest.Email).Return(domain.User{}, nil)
				s.EXPECT().CreateUser(gomock.Any(), signupRequest).Return(uuid, nil)
				s.EXPECT().CreateAccessToken(gomock.Any(), gomock.Any(), gomock.Any()).Return("access_token", nil)
				s.EXPECT().CreateRefreshToken(gomock.Any(), gomock.Any(), gomock.Any()).Return("refresh_token", nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			test_name: "user with this email already exists",
			inputBody: `{"name":"John", "username":"john123", "password":"123456", "email":"john@123.ru"}`,
			inputRequest: signup.SignUpRequest{
				Name:     "John",
				Username: "john123",
				Email:    "john@123.ru",
				Password: "123456",
			},
			mockBehavior: func(s *mock_signup.MockSignUpUseCase, signupRequest signup.SignUpRequest) {
				s.EXPECT().GetUserByEmail(gomock.Any(), signupRequest.Email).Return(domain.User{
					Name:     "John",
					Username: "john123",
					Email:    "john@123.ru",
					Password: "123456",
				}, nil)
			},
			expectedStatusCode: http.StatusConflict,
		},
		{
			test_name: "database error",
			inputBody: `{"name":"John", "username":"john123", "password":"123456", "email":"john@123.ru"}`,
			inputRequest: signup.SignUpRequest{
				Name:     "John",
				Username: "john123",
				Email:    "john@123.ru",
				Password: "123456",
			},
			mockBehavior: func(s *mock_signup.MockSignUpUseCase, signupRequest signup.SignUpRequest) {
				s.EXPECT().GetUserByEmail(gomock.Any(), signupRequest.Email).Return(domain.User{}, errors.New("database error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			test_name:          "invalid request",
			inputBody:          `{"name":"", "username":"john123", "password":"123456", "email":"john@123.ru"}`,
			inputRequest:       signup.SignUpRequest{},
			mockBehavior:       func(s *mock_signup.MockSignUpUseCase, signupRequest signup.SignUpRequest) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			test_name: "generate jwt token problems",
			inputBody: `{"name":"John", "username":"john123", "password":"123456", "email":"john@123.ru"}`,
			inputRequest: signup.SignUpRequest{
				Name:     "John",
				Username: "john123",
				Email:    "john@123.ru",
				Password: "123456",
			},
			mockBehavior: func(s *mock_signup.MockSignUpUseCase, signupRequest signup.SignUpRequest) {
				uuid := uuid.New()
				s.EXPECT().GetUserByEmail(gomock.Any(), signupRequest.Email).Return(domain.User{}, nil)
				s.EXPECT().CreateUser(gomock.Any(), signupRequest).Return(uuid, nil)
				s.EXPECT().CreateAccessToken(gomock.Any(), gomock.Any(), gomock.Any()).Return(gomock.Any().String(), errors.New("generate jwt token problem"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.test_name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			l := mock_logger.NewMockInterface(c)
			v := validate.NewValidator()
			bC := controller.New(l, v)
			singupUC := mock_signup.NewMockSignUpUseCase(c)

			authC := NewAuthController(bC, singupUC, &config.Config{})
			tc.mockBehavior(singupUC, tc.inputRequest)

			r := gin.New()

			binding.Validator = &validate.CustomValidator{
				V: validate.NewValidator(),
			}

			r.POST("/sign-up", authC.SignUp)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", strings.NewReader(tc.inputBody))
			r.ServeHTTP(w, req)

			require.Equal(t, tc.expectedStatusCode, w.Code)
		})
	}
}
