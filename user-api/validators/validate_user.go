package validators

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"user-api.com/usecases"
)

// name cpf email phone_number

type error interface {
	Error() string
}

// User - types of the user
type User struct {
	usecases.UserType
}

// ValidateUser - can validate the data of user
func (ut User) ValidateUser() error {
	return validation.ValidateStruct(&ut,
		validation.Field(&ut.Username, validation.Required),
		validation.Field(&ut.Cpf, validation.Required, validation.Length(11, 11)),
		validation.Field(&ut.Email, validation.Required, is.Email),
		validation.Field(&ut.PhoneNumber, validation.Required, validation.Length(11, 11)),
	)
}

func main() {
	fmt.Println("helo")
}
