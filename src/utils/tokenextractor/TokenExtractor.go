package tokenextractor

import (
	"eletronic_point/src/core/domain/session"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/utils/validator"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var (
	authHeaderParts []string
	tokenParts      []string
)

func GetRolesFromAuthHeader(authHeader string) ([]string, error) {
	token, err := getTokenFromAuthHeader(authHeader)
	if err != nil {
		return []string{AnonymousRole}, nil
	}

	claims, err := getMappedClaimsFromToken(token)
	if err != nil {
		return nil, err
	}

	concatenatedRolesBytes, err := base64.StdEncoding.DecodeString(claims[rolesKey].(string))
	if err != nil {
		return nil, err
	}

	concatenatedRoles := strings.TrimSpace(
		strings.ToLower(string(concatenatedRolesBytes)),
	)

	var roles []string
	roles = strings.Split(concatenatedRoles, commaSymbol)
	return roles, nil
}

func getTokenFromAuthHeader(authHeader string) (string, error) {
	authHeader = strings.TrimSpace(authHeader)
	if !isAuthHeaderValid(authHeader) {
		return emptyString, errors.New(invalidAuthHeaderMsg)
	}

	return authHeaderParts[secondElement], nil
}

func isAuthHeaderValid(authHeader string) bool {
	if validator.TextIsBlank(authHeader) {
		return false
	}
	authHeaderParts = strings.Split(authHeader, spaceInString)
	if len(authHeaderParts) < 2 {
		return false
	}
	return true
}

func getMappedClaimsFromToken(token string) (map[string]interface{}, error) {
	tokenParts = strings.Split(token, dotSymbol)
	payload := tokenParts[secondElement]
	payloadBytes, err := jwt.DecodeSegment(payload)
	if err != nil {
		return nil, err
	}
	claims := make(map[string]interface{})
	err = json.Unmarshal(payloadBytes, &claims)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func GetSessionReferenceFromAuthHeader(authHeader string) (*session.Session, error) {
	token, err := getTokenFromAuthHeader(authHeader)
	if err != nil {
		return &session.Session{}, nil
	}
	claims, err := getMappedClaimsFromToken(token)
	if err != nil {
		return nil, err
	}

	accountID := uuid.MustParse(claims[messages.AccountIDClaim].(string))

	return session.NewReference(accountID), nil
}
