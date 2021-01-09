package validators

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"order-api.com/models"
)

type error interface {
	Error() string
}

// Order - types of the user
type Order struct {
	models.Order
}

// ValidateOrder - can validate the data of user
func (ot Order) ValidateOrder() error {
	return validation.ValidateStruct(&ot,
		validation.Field(&ot.UserID, validation.Required),
	)
}

func main() {
	fmt.Println("hello")
}
