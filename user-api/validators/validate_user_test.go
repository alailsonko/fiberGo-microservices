package validators

import (
	"fmt"
	"testing"
)

type error interface {
	Error() string
}

func TestValidateUser(t *testing.T) {

	// should return error if name is not provided
	ut := new(User)

	ut.Username = "alailson"

	err := ut.ValidateUser()

	expected := "username: cannot be blank."

	if err != nil && err.Error() == expected {
		fmt.Println(err.Error(), expected)
		t.Errorf("should return error if name is not provided: %v", err)
	} else if err == nil {
		fmt.Println(err, expected)
		t.Log("passing: should return error if name is not provided.")
	}

	// should return error if cpf is not provided
	ut.Cpf = "12345678911"

	err = ut.ValidateUser()

	expected = "cpf: cannot be blank."

	if err != nil && err.Error() == expected {
		fmt.Println(err.Error(), expected)
		t.Errorf("should return error if cpf is not provided: %v", err)
	} else if err == nil {
		fmt.Println(err, expected)
		t.Log("passing: should return error if name is not provided.")
	}

}

// name cpf email phone_number

// should return error if cpf min11-max11 is not provided
// should return error if email is not provided
// should return error if email is not invalid
// should return error if phone_number is not provided
// should return error if phone_number min11-max11 is not provided
// should return nil if all data is valid
