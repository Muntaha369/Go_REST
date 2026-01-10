package gtypes

//it's a field that would be required in order to create User in database

type User struct {
	Id       int64
	Name     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}