package routes

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"user-api.com/setup_app"
	"user-api.com/usecases"
)

type user struct {
	usecases.UserType
}

type test struct {
	description string
	// Test input
	route string
	// method
	method string
	// data
	userData string
	// expected output
	expectedError bool
	expectedCode  int
	expectedBody  string
}

// CREATEUser - handler for create user
func TestRoutes(t *testing.T) {
	// UserData := `{
	// 	"username": "valid username",
	// 	"cpf": "12345678911",
	// 	"email": "valid@mail.com",
	// 	"phone_number": "86999999999"
	// }`
	// req, err := http.NewRequest("POST", "//user", bytes.NewBuffer([]byte(UserData)))
	// fmt.Println(err)
	// fmt.Println("testtt", req)
	// res, _ := app.Test(req)
	// fmt.Println("resppp", res)

	tests := []test{
		{
			description: "create use route test - if name is no provided",
			route:       "/user",
			method:      "POST",
			userData: `{
				"username": "alailson",
				"cpf": "12345678911",
				"email": "valid@mail.com",
				"phone_number": "86999999999"
		}`,
			expectedError: false,
			expectedCode:  400,
			expectedBody:  `{"error": username: cannot be blank.}`,
		},
	}

	for _, test := range tests {
		req, e := http.NewRequest(test.method, "/user", bytes.NewBuffer([]byte(test.userData)))
		fmt.Println("error", e)
		fmt.Println("req", req)
		app := setup_app.SetupApp()
		res, err := app.Test(req)
		// fmt.Println(res)
		respRec := httptest.NewRecorder()
		fmt.Println(respRec)
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}
		fmt.Println("res", res)

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		body, err := ioutil.ReadAll(res.Body)

		assert.Nilf(t, err, test.description)

		assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}
