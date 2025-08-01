package signup

import (
	"testing"
	"strconv"

	validate "github.com/London57/todo-app/pkg/validator"
	"github.com/stretchr/testify/require"
)

func TestSignUpRequest_Validation(t *testing.T) {
	validator := validate.NewValidator()
	cases := []struct {
		name string
		request SignUpRequest
		hasError bool
	}{
		{
			"Valid",
			SignUpRequest{Email: "gleb.yurov.1998@mail.ru", Username: "dafjhfg", Password: "123463iuhiog", Name: "ivan"},
			false,
		},
		{
			"Empty all fields",
			SignUpRequest{Email: "", Username: "", Password: "", Name: ""},
			true,
		},
		{
			"Incorrect email",
			SignUpRequest{Email: "addadadsv@", Username: "dadf", Password: "134242", Name: "oleg"},
			true,
		},
		{
			"Incorrect username",
			SignUpRequest{Email: "sdfsfdfsf@mail.ru", Username: "авыаы-=а", Name: "олег", Password: "134242"},
			true,
		},
		{
			"Incorrect username",
			SignUpRequest{Email: "sdfsfdfsf@mail.ru", Username: "dad''dada", Name: "олег", Password: "134242"},
			true,
		},
		{
			"Incorrect username",
			SignUpRequest{Email: "sdfsfdfsf@mail.ru", Username: "d", Name: "gdgd", Password: "134242"},
			true,
		},
		{
			"Empty name",
			SignUpRequest{Email: "dffsfd@sdfs.ru", Username: "adadad", Name: "", Password: "fhgfghh"},
			true,
		},
		{
			"Empty email",
			SignUpRequest{Email: "", Username: "adadad", Name: "oleg", Password: "fhgfghh"},
			true,
		},
		{
			"Empty username",
			SignUpRequest{Email: "dffsfd@sdfs.ru", Username: "", Name: "oleg", Password: "fhgfghh"},
			true,
		},
		{
			"Empty password",
			SignUpRequest{Email: "dffsfd@sdfs.ru", Username: "adadad", Name: "agdnfh", Password: ""},
			true,
		},
		{
			"Incorrect email",
			SignUpRequest{Email: "adad@.r", Username: "12345", Name: "sfsfsfsf", Password: "123456"},
			true,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Struct(tt.request)
			if tt.hasError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	cases := []struct {
		str string
		isValid bool
	}{
		{
			"12311",
			false,
		},
		{
			"3131@mail.ru",
			true,
		},
		{
			"sdfsgs@dadru",
			false,
		},
		{
			"s@g.h",
			true,
		},
	}
	for i, tt := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := IsEmail(tt.str)
			require.Equal(t, res, tt.isValid)
		})
	}
}