package rest

import (
	"encoding/json"
	"errors"
	// "fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Muntaha369/Go_REST/internals/storage"
	gtypes "github.com/Muntaha369/Go_REST/internals/types"
	"github.com/Muntaha369/Go_REST/internals/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {  
	return func(w http.ResponseWriter, r *http.Request) {

		var user gtypes.User

		err := json.NewDecoder(r.Body).Decode(&user) // it reads the Json data from the body and decodes it into the user format or decodes it as per the parameter
		if errors.Is(err, io.EOF) { // if the body is empty throw this error
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err = validator.New().Struct(user); err != nil { // validator.New() creates a validator instance and .Struct(user) is used to validate that user is type User struct mentioned in gtypes.go

			validateErrs := err.(validator.ValidationErrors) // it shows what type of error it is it looks up to User struct in gtype.go and check all fields in User are mentioned or not it return error in an array suppose if more than 2 fields are missing thats why array is used
			response.WriteJson(w, http.StatusBadRequest, response.ValidatorError(validateErrs))
			return
		}

		lastId, err := storage.CreateUser(
			user.Name,
			user.Email,
			user.Password,
		) // it Creates user mentioned in sqlite.go

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		id := r.PathValue("id") //this is used to get the url path ex:(http://localhost:8082/api/getUser/2)
		slog.Info("getting User", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64) //this line converts the id (which is in string) to integer of base 10 (Decimal numberes) and 

		if err !=nil{
			response.WriteJson(w, http.StatusInternalServerError, err)
			return 
		}

		user, err := storage.GetUserById(intId) //this line get users mentioned in sqlite.go

		if err != nil{
			response.WriteJson(w, http.StatusInternalServerError, err)
			return 
		}

		response.WriteJson(w, http.StatusOK, user)
	}
}

func GetByList(storage storage.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		slog.Info("Getting all User")

		user, err := storage.GetUserList() // it gets all user

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return 
		}

		response.WriteJson(w, http.StatusOK, user)
	}
}
