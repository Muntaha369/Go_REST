package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string
	Error  string
}

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json") //this line sets the content type application/json in header by default it is plain/text this let the client or browser know that the you are recieving json
	w.WriteHeader(status) //this set status in browser like 404 201 etc.

	return json.NewEncoder(w).Encode(data) //this line actually set the data to json
}

func GeneralError(err error) Response { // This error Response is recieved in WriteJson func
	return Response{
		Status: "Error",
		Error:  err.Error(),
	}
}

func ValidatorError(errs validator.ValidationErrors) Response { 
	var errMsgs []string //declared to hold errs from the parameter recieved 

	for _, err := range errs { //the error recieved from parameters is in array
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid", err.Field()))
		}
	}

	return Response{
		Status: "Error",
		Error:  strings.Join(errMsgs, ", "), //this joins all err message with coma and returns in string
	}
}
