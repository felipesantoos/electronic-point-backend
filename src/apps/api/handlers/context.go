package handlers

import (
	"eletronic_point/src/apps/api/utils"
	"eletronic_point/src/core/domain/authorization"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/role"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RichContext interface {
	echo.Context

	AccountID() *uuid.UUID
	Name() string
	ProfileID() *uuid.UUID
	RoleName() string
	IsAdmin() bool

	GetUUIDPathParam(key string) (*uuid.UUID, errors.Error)
}

type RichHandler = func(RichContext) error

type richContext struct {
	echo.Context

	accountID *uuid.UUID
	name      string
	profileID *uuid.UUID
	roleName  string
}

func NewRichContext(ctx echo.Context, claims *authorization.AuthClaims) (RichContext, error) {
	if claims != nil {
		role, _ := utils.DecodeRoleData(claims.Role)
		accountID, _ := uuid.Parse(claims.AccountID)
		profileID, _ := uuid.Parse(claims.ProfileID)
		return &richContext{ctx, &accountID, claims.Name, &profileID, strings.ToLower(role)}, nil
	}
	return &richContext{ctx, nil, "", nil, role.ANONYMOUS_ROLE_CODE}, nil
}

func (c *richContext) AccountID() *uuid.UUID {
	return c.accountID
}

func (c *richContext) Name() string {
	return c.name
}

func (c *richContext) ProfileID() *uuid.UUID {
	return c.profileID
}

func (c *richContext) RoleName() string {
	return c.roleName
}

func (c *richContext) IsAdmin() bool {
	return c.roleName == role.ADMIN_ROLE_CODE
}

func (c *richContext) GetUUIDPathParam(key string) (*uuid.UUID, errors.Error) {
	strUUID := c.Param(key)
	if strUUID == "" {
		return nil, errors.NewFromString(fmt.Sprintf("you must provide a valid %s", key))
	} else if uuid, err := uuid.Parse(strUUID); err != nil {
		return nil, errors.NewFromString("the provided id is invalid")
	} else {
		return &uuid, nil
	}
}

func (c *richContext) GetStringPathParam(key string) string {
	return c.GetPathParam(key)
}

func (c *richContext) GetIntPathParam(key string) (int, *echo.HTTPError) {
	value := c.GetPathParam(key)
	if intValue, err := strconv.Atoi(value); err != nil {
		return 0, &echo.HTTPError{
			Message: fmt.Sprintf("the provided value for %s must be an integer", key),
		}
	} else {
		return intValue, nil
	}
}

func (c *richContext) GetPathParam(key string) string {
	value := c.Param(key)
	return value
}
