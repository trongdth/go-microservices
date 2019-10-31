package serializers

// UserLoginReq : struct
type UserLoginReq struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

// UserRegisterReq : struct
type UserRegisterReq struct {
	FullName        string `json:"FullName"`
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	ConfirmPassword string `json:"ConfirmPassword"`
}
