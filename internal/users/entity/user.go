package entity

type User struct {
	Id       string `db:"id"`
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"password"`
}
