package docs

// HTTPError
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type HTTPOk struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"status successfully"`
}

type Order struct {
	UserID          string  `json:"user_id" example:"6"`
	ItemDescription string  `json:"item_description" example:"valid description"`
	ItemQuantity    float64 `json:"item_quantity" example:"6"`
	ItemPrice       float64 `json:"item_price" example:"353.5"`
}

type User struct {
	Username    string `json:"username" example:"valid username"`
	Email       string `json:"email" example:"valid@mail.com"`
	Cpf         string `json:"cpf" example:"12345678911"`
	PhoneNumber string `json:"phone_number" example:"8699999999"`
}

type OrdersByUser struct {
	orders []Order
	user   User
}
