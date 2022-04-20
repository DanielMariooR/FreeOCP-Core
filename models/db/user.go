package db

type User struct {
	ID       string `db:"id"`
	Name     string `db:"fullname"`
	Email    string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
	Admin    bool   `db:"isAdmin"`
}
