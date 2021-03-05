package login

type oauth_bearer_request_response struct {

	Error string `json:"error"`
	ErrorDescription string `json:"error_description"`
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
	TokenType string `json:"token_type"`
	Scope string `json:"scope"`
	RefreshToken string `json:"refresh_token"`

}

type oauth_identity_response struct {

	ID string `json:"id"`
	Username string `json:"user_login"`
	Email string `json:"user_email"`
	CreationDate string `json:"user_registered"`
	DisplayName string `json:"display_name"`

}