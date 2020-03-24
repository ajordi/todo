package authenticating

type UserDTO struct {
	ID       uint   `json:"id,string,omitempty"`
	UserName string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password string `json:"password"`
}
