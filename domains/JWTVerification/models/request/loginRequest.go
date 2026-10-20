package request

type LoginRequest struct {
	UserId      string `json:"userId" validate:"required"`
	AccessToken string `json:"accessToken" validate:"required"`
}
