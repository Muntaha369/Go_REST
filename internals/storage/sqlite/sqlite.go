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
	db, err := sql.Open("sqlite", cfg.StoragePath) // it creates lazy connection i.e it does not connect untill the first query is executed but it verifies that the connection is ready to estblish or not if not then it returns an error

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS user(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    email TEXT,
    password TEXT
)`) //typical SQL query

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db, // it returns all the methods which are below 
	}, nil

}

func (s *Sqlite) CreateUser(name string, email string, password string) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO user (name, email, password) VALUES(?, ?, ?)") //prepare statement for later queries

	if err != nil {
		return 0, err
	}
	defer stmt.Close() //need to close the connection after this function is compleatly executed

	result, err := stmt.Exec(name, email, password) //now executing the statement

	if err != nil {
		return 0, err
	}

	lastid, err := result.LastInsertId() // retuns the id of the newly geneated data or entry in the database

	if err != nil {
		return 0, err
	}

	return lastid, nil
}

func (s *Sqlite) GetUserById(id int64) (gtypes.User, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM user WHERE id = ? LIMIT 1") //prepare statement for later queries

	if err != nil {
		return gtypes.User{}, err // return nil or error
	}

	defer stmt.Close()  //need to close the connection after this function is compleatly executed

	var user gtypes.User // creates a struct of User

	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Email, &user.Password) //it scans the row based on id using this QueryRow(id) and sets the value inside the user variable (this operation does this Scan(&user.Id, &user.Name, &user.Email, &user.Password))

	if err != nil {

		if err == sql.ErrNoRows {
			return gtypes.User{}, fmt.Errorf("No rows of id %s exists", fmt.Sprint(id)) 
		}

		return gtypes.User{}, fmt.Errorf("querry error : %w", err)
	}

	return user, nil
}

func (s *Sqlite) GetUserList() ([]gtypes.User, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM user") //prepare statement for later queries

	if err != nil {
		return []gtypes.User{}, err
	}

	defer stmt.Close() //need to close the connection after this function is compleatly executed

	rows, err := stmt.Query() // it returns all rows

	if err != nil {
		return []gtypes.User{}, err
	}

	defer rows.Close() //need to close the connection after this function is compleatly executed

	var users []gtypes.User //it contains array of User

	for rows.Next() {
		var user gtypes.User

		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password) //it copies the rows in the User struct

		if err != nil {
			return []gtypes.User{}, err
		}
		users = append(users, user) //it attaches all the users to above user struct
	}

	return users, nil
}
