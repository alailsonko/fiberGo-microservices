package validators

import (
	"fmt"
	"testing"
)

func TestValidateOrder(t *testing.T) {
	// should return error case user_id is not provided
	ot := new(Order)

	ot.UserID = "valid_id"

	err := ot.ValidateOrder()

	expected := "user_id: cannot be blank."

	if err != nil && err.Error() == expected {
		fmt.Println(err.Error(), expected)
		t.Errorf("should return error if user_id is not provided: %v", err)
	} else if err == nil {
		fmt.Println(err, expected)
		t.Log("passing: should return error if user_id is not provided.")
	}
	// should return error case item_description is not provided
	ot.ItemDescription = `v
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss

	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	i`

	err = ot.ValidateOrder()

	expected = "item_description: cannot be blank."

	if err != nil && err.Error() == expected {
		fmt.Println(err.Error(), expected)
		t.Errorf("should return error if item_description is not provided: %v", err)
	} else if err == nil {
		fmt.Println(err, expected)
		t.Log("passing: should return error if item_description is not provided.")
	}
	// should return error case item_description min5-max 200 is not provided
	ot.ItemDescription = `v
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss
	fdsssssssssssssssssssssssss

	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	fsfsdfsd
	i`

	err = ot.ValidateOrder()

	expected = "item_description: error case item_description min5-max 200 is not provided."

	if err != nil && err.Error() == expected {
		fmt.Println(err.Error(), expected)
		t.Errorf("error case item_description min5-max 200 is not provided: %v", err)
	} else if err == nil {
		fmt.Println(err, expected)
		t.Log("passing: error case item_description min5-max 200 is not provided.")
	}
	// should return error case item_quantity is not provided

}

// should return error case item quantity is not float
// should return error case item_price is not provided
// should return error case item_price is not float
// should return error case order data is valid
