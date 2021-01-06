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

	// should return error if cpf is not length 11

	err = ut.ValidateUser()
	fmt.Println(err)

	expected = "cpf: the length must be exactly 11."

	if err != nil && err.Error() == expected {
		fmt.Println(err.Error(), expected)
		t.Errorf("should return error if cpf is not length 11: %v", err)
	} else if err == nil {
		fmt.Println(err, expected)
		t.Log("passing: should return error if cpf is not length 11")
	}
	// should return error if email is not provided
	ut.Email = "valid@mail.com"

	err = ut.ValidateUser()
	fmt.Println(err)

	expected = "email: cannot be blank."

	if err != nil && err.Error() == expected {
		fmt.Println(err.Error(), expected)
		t.Errorf("should return error if email is not provided: %v", err)
	} else if err == nil {
		fmt.Println(err, expected)
		t.Log("passing: should return error if email is not provided")
	}
	// should return error if email is not invalid
	err = ut.ValidateUser()
	fmt.Println(err)

	expected = "email: must be a valid email address."

	if err != nil && err.Error() == expected {
		fmt.Println(err.Error(), expected)
		t.Errorf("should return error if email is not valid: %v", err)
	} else if err == nil {
		fmt.Println(err, expected)
		t.Log("passing: should return error if email is not valid")
	}
	// should return error if phone_number is not provided
	ut.PhoneNumber = "86999999999"
	err = ut.ValidateUser()
	fmt.Println(err)

	expected = "phone_number: cannot be blank."

	if err != nil && err.Error() == expected {
		fmt.Println(err.Error(), expected)
		t.Errorf("should return error if phone_number is not provided: %v", err)
	} else if err == nil {
		fmt.Println(err, expected)
		t.Log("passing: should return error if phone_number is not provided")
	}

	// should return error if phone_number min11-max11 is not provided
	expected = "phone_number: the length must be exactly 11."

	if err != nil && err.Error() == expected {
		fmt.Println(err.Error(), expected)
		t.Errorf("should return error if phone_number the length must be exactly 11.: %v", err)
	} else if err == nil {
		fmt.Println(err, expected)
		t.Log("passing: should return error if phone_number the length must be exactly 11.")
	}

	// should return nil if all data is valid
	if err == nil {
		t.Log("passing: should return nil if all data is valid.")
	}
}

// name cpf email phone_number
