package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/Muntaha369/Go_REST/internals/config"
	gtypes "github.com/Muntaha369/Go_REST/internals/types"
	_ "modernc.org/sqlite"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite", cfg.StoragePath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS user(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    email TEXT,
    password TEXT
)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil

}

func (s *Sqlite) CreateUser(name string, email string, password string) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO user (name, email, password) VALUES(?, ?, ?)")

	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, password)

	if err != nil {
		return 0, err
	}

	lastid, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lastid, nil
}

func (s *Sqlite) GetUserById(id int64) (gtypes.User, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM user WHERE id = ? LIMIT 1")

	if err != nil {
		return gtypes.User{}, err
	}

	defer stmt.Close()

	var user gtypes.User

	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	if err != nil {

		if err == sql.ErrNoRows{
			return  gtypes.User{}, fmt.Errorf("No rows of id %s exists", fmt.Sprint(id))
		}

		return gtypes.User{}, fmt.Errorf("querry error : %w", err)
	}

	return user, nil
}