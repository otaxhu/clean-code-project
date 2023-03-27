package entity

type UserRole struct {
	Id     string `db:"id"`
	UserId string `db:"user_id"`
	RoleId int    `db:"role_id"`
}
