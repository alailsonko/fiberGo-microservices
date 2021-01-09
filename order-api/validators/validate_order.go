package validators

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"order-api.com/models"
)

type error interface {
	Error() string
}

// Order - types of the order
type Order struct {
	models.Order
}

// ValidateOrder - can validate the data of order
func (ot Order) ValidateOrder() error {
	return validation.ValidateStruct(&ot,
		validation.Field(&ot.UserID, validation.Required),
		validation.Field(&ot.ItemDescription, validation.Length(5, 100), validation.Required),
	)
}

func main() {
	fmt.Println("hello")
}
