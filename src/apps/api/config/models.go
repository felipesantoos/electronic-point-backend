package config

import (
	"errors"

	"eletronic_point/src/utils/validator"

	"go.uber.org/multierr"
)

type ServerConfiguration struct {
	Port        int
	HostAddress string
}

func (s ServerConfiguration) Validate() error {
	var errs error

	if s.Port < 0 {
		errs = multierr.Append(errs, errors.New("the parameter 'port' is invalid"))
	}

	if !validator.HostAddressIsValid(s.HostAddress) {
		errs = multierr.Append(errs, errors.New("the parameter 'host_address' is invalid or is typed incorrectly"))
	}

	return errs
}

type PostgresConfiguration struct {
	User           string
	Password       string
	DBName         string
	Host           string
	Port           int
	SSLMode        string
	MigrationsPath string
}

type AuthorizationConfig struct {
	AuthPaths AuthPaths
}

type AuthPaths struct {
	AuthModelPath  string
	AuthPolicyPath string
}

type CORSPolicy struct {
	ServerAllowedHosts []string
}

type Config struct {
	HttpServer    *ServerConfiguration
	Postgres      *PostgresConfiguration
	Authorization *AuthorizationConfig
	CORSPolicy    *CORSPolicy
}

func (c Config) Validate() error {
	errs := multierr.Combine(c.HttpServer.Validate())
	return errs
}
