package handlers

import (
	"dit_backend/src/api/handlers/dto/response"
	"dit_backend/src/core/domain/authorization"
	"dit_backend/src/core/domain/errors"
	"dit_backend/src/core/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func responseFromError(context echo.Context, err errors.Error) error {
	statusCode := statusCodeByError(err)
	return responseFromErrorAndStatus(context, err, statusCode)
}

func responseFromErrorAndStatus(context echo.Context, err errors.Error, statusCode int) error {
	dtoErr := response.NewErrorFromCore(err, statusCode)
	return context.JSON(statusCode, dtoErr)
}

func statusCodeByError(err errors.Error) int {
	if err.CausedByValidation() {
		return http.StatusUnprocessableEntity
	} else if err.CausedInternally() {
		return http.StatusInternalServerError
	}
	return http.StatusBadRequest
}

func getAuthClaims(authHeader string) (*authorization.AuthClaims, errors.Error) {
	_, token := utils.ExtractToken(authHeader)
	authClaims, err := utils.ExtractTokenClaims(token)
	if err != nil {
		return nil, errors.NewFromString("Invalid authorization. Please login and try again.")
	}
	return authClaims, nil
}

func getAccountIDFromAuthorization(ctx echo.Context) (*uuid.UUID, errors.Error) {
	claims, err := getAuthClaims(ctx.Request().Header.Get("Authorization"))
	if err != nil {
		return nil, err
	}
	if accountID, parseErr := uuid.Parse(claims.AccountID); parseErr != nil {
		return nil, errors.NewFromString("Invalid Account ID. Please login and try again.")
	} else {
		return &accountID, nil
	}
}
