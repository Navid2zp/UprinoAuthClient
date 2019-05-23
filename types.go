package UprinoAuthClient

type TokenCheck struct {
	Token      string
	PublicKey  string
	PrivateKey string
	Valid      bool
	Checked    bool
}

type TokenUserData struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type TokenCheckResponse struct {
	Status   string        `json:"status"`
	Message  string        `son:"message"`
	Error    string        `json:"error"`
	UserData TokenUserData `json:"user_data"`
}
