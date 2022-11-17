package defined

type LoginRequest struct {
	Telephone string `json:"telephone"`
	Code      string `json:"code"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
