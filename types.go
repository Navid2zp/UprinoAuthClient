package UprinoAuthClient

type TokenCheck struct {
	Token      string
	PublicKey  string
	PrivateKey string
	Valid      bool
	Checked    bool
}

type TokenCheckResponse struct {
	Status   string `db:"status" json:"status"`
	Message  string `db:"message" json:"message"`
	Error    string `db:"error" json:"error"`
	UserData struct {
		ID        int    `db:"id" json:"id"`
		Username  string `db:"username" json:"username"`
		Email     string `db:"email" json:"email"`
		FirstName string `db:"first_name" json:"first_name"`
		LastName  string `db:"last_name" json:"last_name"`
	}
}
