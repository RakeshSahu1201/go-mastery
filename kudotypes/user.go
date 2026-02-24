package kudotypes

type User struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
