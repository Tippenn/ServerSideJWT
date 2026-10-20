package request

type CheckJWTRequest struct {
	AccessToken string `json:"accessToken" validate:"required"`
}
