package test

type Transactions struct {
	ID               string `json:"id" db:"id"`
	User_id          string `json:"user_id" db:"user_id"`
	User_email       string `json:"email" db:"email"`
	Sum              string `json:"sum" db:"sum"`
	Currensy         string `json:"currensy" db:"currensy"`
	Date_create      string `json:"date_create" db:"date_create"`
	Date_last_update string `json:"date_last_update" db:"date_last_update"`
	Status           string `json:"status" db:"status"`
}
