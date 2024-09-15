package controller

import (
	"fmt"
	"net/http"

	"notjustadeveloper.com/sales-agent-server/pkg/errors"
)

func HandleError(w http.ResponseWriter, wrapErr *errors.WrapError) {
	fmt.Println(wrapErr.Error())
	switch wrapErr.ErrorType {
	case errors.NotFoundError:
		w.WriteHeader(http.StatusNotFound)
	case errors.BadRequest:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
