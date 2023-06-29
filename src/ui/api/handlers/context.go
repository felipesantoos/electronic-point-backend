package handlers

import (
	"backend_template/src/core/domain/authorization"
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/role"
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EnhancedContext interface {
	echo.Context

	AccountID() *uuid.UUID
	ProfileID() *uuid.UUID
	RoleName() string
	IsAdmin() bool

	GetUUIDPathParam(key string) (*uuid.UUID, errors.Error)
}

type EnhancedHandler = func(EnhancedContext) error

type enhancedContext struct {
	echo.Context

	accountID *uuid.UUID
	profileID *uuid.UUID
	roleName  string
}

func NewEnhancedContext(ctx echo.Context, claims *authorization.AuthClaims) (EnhancedContext, error) {
	var isAuthenticated = false
	if v, ok := ctx.Get("authenticated").(bool); ok && v {
		isAuthenticated = true
	}
	if isAuthenticated && claims != nil {
		accountID, _ := uuid.Parse(claims.AccountID)
		profileID, _ := uuid.Parse(claims.ProfileID)
		return &enhancedContext{ctx, &accountID, &profileID, claims.Role}, nil
	}
	return &enhancedContext{ctx, nil, nil, ""}, nil
}

func (c *enhancedContext) AccountID() *uuid.UUID {
	return c.accountID
}

func (c *enhancedContext) ProfileID() *uuid.UUID {
	return c.profileID
}

func (c *enhancedContext) RoleName() string {
	return c.RoleName()
}

func (c *enhancedContext) IsAdmin() bool {
	return c.roleName == role.ADMIN_ROLE_CODE
}

func (c *enhancedContext) GetUUIDPathParam(key string) (*uuid.UUID, errors.Error) {
	strUUID := c.Param(key)
	if strUUID == "" {
		return nil, errors.NewFromString(fmt.Sprintf("you must provide a valid %s", key))
	} else if uuid, err := uuid.Parse(strUUID); err != nil {
		return nil, errors.NewFromString("the provided id is invalid")
	} else {
		return &uuid, nil
	}
}
