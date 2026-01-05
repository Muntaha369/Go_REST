package rest

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	gtypes "github.com/Muntaha369/Go_REST/internals/types"
	"github.com/Muntaha369/Go_REST/internals/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user gtypes.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if errors.Is(err, io.EOF){
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return 
		}

		response.WriteJson(w, http.StatusCreated, map[string] string {"success":"ok"})

		w.Write([]byte("Welcome to REST api"))
	}
}