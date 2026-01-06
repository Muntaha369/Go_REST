package gtypes

type User struct {
	Id       int
	Name     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
