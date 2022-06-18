package test

type User struct {
	ID    int    `json:"-"`
	Email string `json:"email"`
}
