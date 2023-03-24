package entity

type User struct {
	Id       string `db:"id" bson:"_id"`
	Email    string `db:"email" bson:"email"`
	Name     string `db:"name" bson:"name"`
	Password string `db:"password" bson:"password"`

	// This field is only for mongo, for sql DB use the UserRole struct that have his own table
	Roles []UserRole `db:"-" bson:"roles"`
}
