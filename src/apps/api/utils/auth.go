package utils

import (
	"eletronic_point/src/core"
	"eletronic_point/src/core/domain/authorization"
	"eletronic_point/src/core/domain/role"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

var logger = core.Logger()

func ExtractAuthorizationAccountRole(authToken string) (string, bool) {
	if authToken == "" {
		return role.ANONYMOUS_ROLE_CODE, true
	} else if claims, ok := authorizationIsValid(authToken); !ok {
		return role.ANONYMOUS_ROLE_CODE, false
	} else {
		return DecodeRoleData(claims.Role)
	}
}

func ExtractToken(authHeader string) (authType string, token string) {
	authorization := strings.Split(strings.TrimSpace(authHeader), " ")
	if len(authorization) < 2 {
		return "", ""
	}
	authType = authorization[0]
	token = authorization[1]
	return authType, token
}

func authorizationIsValid(authToken string) (*authorization.AuthClaims, bool) {
	secret := os.Getenv("SERVER_SECRET")
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		logger.Error().Msg("error parsing the provided token on (signature is invalid?)")
		return nil, false
	}
	if !token.Valid || token.Claims.Valid() != nil {
		logger.Error().Msg("the provided token is invalid or expired")
		return nil, false
	}
	claims, err := ExtractTokenClaims(authToken)
	if err != nil {
		return nil, false
	}
	return claims, true
}

func ExtractTokenClaims(authToken string) (*authorization.AuthClaims, error) {
	if authToken == "" {
		return nil, errors.New("empty token")
	}
	parts := strings.Split(authToken, ".")
	if len(parts) < 3 {
		return nil, errors.New("invalid token format")
	}
	payload := parts[1]
	payloadBytes, err := jwt.DecodeSegment(payload)
	if err != nil {
		logger.Error().Msg("an error occurred when decoding the token payload: " + err.Error())
		return nil, err
	}
	var claims authorization.AuthClaims
	err = json.Unmarshal(payloadBytes, &claims)
	if err != nil {
		logger.Error().Msg("an error occurred when unmarshalling the token payload: " + err.Error())
		return nil, err
	}
	return &claims, nil
}

func DecodeRoleData(encodedRoleName string) (string, bool) {
	unmarshedRoleData := make(map[string]interface{})
	roleData, err := base64.StdEncoding.DecodeString(encodedRoleName)
	if err != nil {
		logger.Error().Msg("an error occurred when decoding the role data: " + err.Error())
		return role.ANONYMOUS_ROLE_CODE, false
	}
	if err := json.Unmarshal(roleData, &unmarshedRoleData); err != nil {
		logger.Error().Msg("an error occurred when unmarshaling the role data: " + err.Error())
		return role.ANONYMOUS_ROLE_CODE, false
	}
	return fmt.Sprint(unmarshedRoleData["code"]), true
}
