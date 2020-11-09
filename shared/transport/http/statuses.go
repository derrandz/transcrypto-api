package http

import (
	"net/http"
)

type (
	Status struct {
		HttpCode    int
		HttpMessage string
	}
)

var (
	ErrInvalidData = &Status{
		HttpCode:    http.StatusBadRequest,
		HttpMessage: "Not a valid JSON",
	}

	ErrServerError = &Status{
		HttpCode:    http.StatusInternalServerError,
		HttpMessage: "Error on our side, sorry, please retry later",
	}

	RequestOk = &Status{
		HttpCode:    http.StatusAccepted,
		HttpMessage: "Success",
	}
)
