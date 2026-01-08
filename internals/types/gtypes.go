package gtypes

type User struct {
	Id       int64
	Name     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
