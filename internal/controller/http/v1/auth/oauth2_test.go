package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/London57/todo-app/config"
	"github.com/London57/todo-app/internal/controller/http/common/controller"
	mock_signup "github.com/London57/todo-app/internal/domain/signup/mocks"
	mock_logger "github.com/London57/todo-app/pkg/logger/mocks"
	validate "github.com/London57/todo-app/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAuthController_oauth2(t *testing.T) {
	testcases := []struct {
		testname           string
		providerParam      string
		redirectUrl        string
		expectedStatusCode int
	}{
		{
			"OK",
			"google",
			"https://accounts.google.com",
			http.StatusFound,
		},
		{
			"bad provider param",
			"123",
			"",
			http.StatusBadRequest,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.testname, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			l := mock_logger.NewMockInterface(c)
			v := validate.NewValidator()
			bC := controller.New(l, v)
			singupUC := mock_signup.NewMockSignUpUseCase(c)

			config := &config.Config{}
			authC := NewAuthController(bC, singupUC, config)

			r := gin.New()

			binding.Validator = &validate.CustomValidator{
				V: validate.NewValidator(),
			}

			r.GET("/oauth2/:provider", authC.OAuth2)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/oauth2/"+tc.providerParam, nil)
			r.ServeHTTP(w, req)

			require.Contains(t, w.Header().Get("Location"), tc.redirectUrl)
			require.Equal(t, tc.expectedStatusCode, w.Code)
		})
	}
}
