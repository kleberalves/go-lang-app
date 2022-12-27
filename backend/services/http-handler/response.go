package httphandler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
)

type ResponseError struct {
	Message string `json:"message"`
}

type RParams struct {
	Context *gin.Context
	Err     error
	Obj     any
}

func Response(p RParams) {

	//Means OK, no errors
	if p.Err == nil {

		//OK with valid Object to return.
		if p.Obj != nil {
			p.Context.JSON(http.StatusOK, p.Obj)
		} else {
			//OK but without Object to return, then 204.
			p.Context.Writer.WriteHeader(http.StatusNoContent)
		}
	} else {

		var pgError *pgconn.PgError

		if errors.As(p.Err, &pgError) {
			//Check Database Postgres errors
			switch pgError.Code {
			case "23505":
				//Duplicate. Constraint conflict.
				p.Context.Writer.WriteHeader(http.StatusConflict)
			case "23503":
				//FK conflict.
				p.Context.JSON(http.StatusNotAcceptable, ResponseError{Message: p.Err.Error()})

			default:
				//Generic database error
				p.Context.JSON(http.StatusExpectationFailed, ResponseError{Message: p.Err.Error()})
			}

		} else {
			strError := p.Err.Error()

			switch strError {
			case "record not found":
				//Despite being a database error, its not Postgres error it also has no code
				p.Context.JSON(http.StatusNotFound, ResponseError{Message: strError})
			default:
				p.Context.JSON(http.StatusInternalServerError, ResponseError{Message: strError})

			}

		}
	}

}
