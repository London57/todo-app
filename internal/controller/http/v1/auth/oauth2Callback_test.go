package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/London57/todo-app/config"
	"github.com/London57/todo-app/internal/controller/http/common/controller"
	"github.com/London57/todo-app/internal/domain"
	mock_signup "github.com/London57/todo-app/internal/domain/signup/mocks"
	mock_logger "github.com/London57/todo-app/pkg/logger/mocks"
	validate "github.com/London57/todo-app/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAuthController_oauth2Callback(t *testing.T) {
	testcases := []struct {
		testname           string
		providerParam      string
		queryParams        string
		mockBehaviour      func(*mock_signup.MockSignUpUseCase)
		expectedStatusCode int
	}{
		{
			"OK",
			"google",
			"?code=valid-code&state=123",
			func(c *mock_signup.MockSignUpUseCase) {
				c.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(domain.User{}, nil)
				c.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(uuid.New(), nil)
				c.EXPECT().CreateAccessToken(gomock.Any(), gomock.Any(), gomock.Any()).Return("access_token", nil)
				c.EXPECT().CreateRefreshToken(gomock.Any(), gomock.Any(), gomock.Any()).Return("refresh_token", nil)
			},
			http.StatusCreated,
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

			config := &config.Config{
				OAuth2: config.OAuth2{
					Google: config.Google{
						GoogleClientId:     "test-client-id",
						GoogleClientSecret: "test-client-secret",
					},
					OAuthStateString: "123",
				},
			}

			authC := NewAuthController(bC, singupUC, config)

			tc.mockBehaviour(singupUC)
			r := gin.New()

			binding.Validator = &validate.CustomValidator{
				V: validate.NewValidator(),
			}

			r.GET("/oauth2/:provider/callback", authC.OAuth2Callback)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/oauth2/"+tc.providerParam+"/callback"+tc.queryParams, nil)
			r.ServeHTTP(w, req)
			require.Equal(t, tc.expectedStatusCode, w.Code)
		})
	}
}
