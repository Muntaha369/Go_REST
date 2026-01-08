package storage

import gtypes "github.com/Muntaha369/Go_REST/internals/types"

type Storage interface {
	CreateUser(name string, email string, password string) (int64, error)
	GetUserById(id int64) (gtypes.User, error)
}
