package request

type SignUpRequest struct {
	Username        string `json:"username" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmpassword" validate:"required"`
}
