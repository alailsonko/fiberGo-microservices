package validators

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"user-api.com/usecases"
)

// name cpf email phone_number

// User - types of the user
type User struct {
	usecases.UserType
}

// ValidateUser - can validate the data of user
func (ut User) ValidateUser() error {
	return validation.ValidateStruct(&ut,
		validation.Field(&ut.Username, validation.Required),
	)
}

func main() {
	fmt.Println("helo")
}
