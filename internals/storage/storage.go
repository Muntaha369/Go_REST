package storage

import gtypes "github.com/Muntaha369/Go_REST/internals/types"

type Storage interface {
	//This are all the methods that can be used by variables type storage 
	CreateUser(name string, email string, password string) (int64, error)
	GetUserById(id int64) (gtypes.User, error)
	GetUserList()([]gtypes.User, error)
}
