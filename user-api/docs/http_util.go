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

type User struct {
	Username    string `json:"username" example:"valid username"`
	Email       string `json:"email" example:"valid@mail.com"`
	Cpf         string `json:"cpf" example:"12345678911"`
	PhoneNumber string `json:"phone_number" example:"8699999999"`
}
