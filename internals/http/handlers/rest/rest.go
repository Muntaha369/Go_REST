package rest

import (
	"encoding/json"
	"errors"
	"fmt"
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

		err := json.NewDecoder(r.Body).Decode(&user)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err = validator.New().Struct(user); err != nil {

			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidatorError(validateErrs))
			return
		}

		lastId, err := storage.CreateUser(
			user.Name,
			user.Email,
			user.Password,
		)

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		id := r.PathValue("id")
		slog.Info("getting User", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)

		if err !=nil{
			response.WriteJson(w, http.StatusInternalServerError, err)
			return 
		}

		user, err := storage.GetUserById(intId)

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

		user, err := storage.GetUserList()

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return 
		}

		response.WriteJson(w, http.StatusOK, user)
	}
}
