package test

type Transaction struct {
	ID               int    `json:"-"`
	User_id          string `json:"user_id" binding:"required"`
	User_email       string `json:"email" binding:"required"`
	Sum              string `json:"sum" binding:"required"`
	Currensy         string `json:"currensy" binding:"required"`
	Date_create      string `json:"date_create"`
	Date_last_update string `json:"date_last_update"`
	Status           string `json:"status"`
}
