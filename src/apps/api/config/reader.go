package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"go.uber.org/multierr"
)

type reader struct{}

func (r reader) config() (*Config, error) {
	var errs error
	server, err := r.server()
	if err != nil {
		errs = multierr.Append(errs, err)
	}

	postgres, err := r.postgres()
	if err != nil {
		errs = multierr.Append(errs, err)
	}

	auth, err := r.auth()
	if err != nil {
		errs = multierr.Append(errs, err)
	}

	cors, err := r.cors()
	if err != nil {
		errs = multierr.Append(errs, err)
	}

	if errs != nil {
		return nil, errs
	}

	return &Config{
		HttpServer:    server,
		Postgres:      postgres,
		Authorization: auth,
		CORSPolicy:    cors,
	}, nil
}

func (r reader) server() (*ServerConfiguration, error) {
	var serverErrs error
	serverHost, ok := os.LookupEnv("SERVER_HOST")
	if !ok {
		serverErrs = fmt.Errorf("could not read the server host in the env")
	}
	serverPortStr, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		serverErrs = multierr.Append(serverErrs, fmt.Errorf("could not read the server port in the env"))
	}
	serverPort, err := strconv.Atoi(serverPortStr)
	if err != nil {
		serverErrs = multierr.Append(serverErrs, fmt.Errorf("could not parse the server port: %s", err))
	}

	if serverErrs != nil {
		return nil, serverErrs
	}

	return &ServerConfiguration{
		HostAddress: serverHost,
		Port:        serverPort,
	}, nil
}

func (r reader) postgres() (*PostgresConfiguration, error) {
	var postgresErrs error
	user, ok := os.LookupEnv("DATABASE_USER")
	if !ok {
		postgresErrs = fmt.Errorf("could not find postgres user in the env")
	}
	password, ok := os.LookupEnv("DATABASE_PASSWORD")
	if !ok {
		postgresErrs = multierr.Append(postgresErrs, fmt.Errorf("could not find postgres password in the env"))
	}
	dbName, ok := os.LookupEnv("DATABASE_NAME")
	if !ok {
		postgresErrs = multierr.Append(postgresErrs, fmt.Errorf("could not find postgres db name in the env"))
	}
	host, ok := os.LookupEnv("DATABASE_HOST")
	if !ok {
		postgresErrs = multierr.Append(postgresErrs, fmt.Errorf("could not find postgres host in the env"))
	}
	sslMode, ok := os.LookupEnv("DATABASE_SSL_MODE")
	if !ok {
		postgresErrs = multierr.Append(postgresErrs, fmt.Errorf("could not find postgres ssl mode in the env"))
	}
	migrationsPath, ok := os.LookupEnv("DATABASE_MIGRATIONS_PATH")
	if !ok {
		postgresErrs = multierr.Append(postgresErrs, fmt.Errorf("could not find postgres migrations path in the env"))
	}
	postgresPortStr, ok := os.LookupEnv("DATABASE_PORT")
	if !ok {
		postgresErrs = multierr.Append(postgresErrs, fmt.Errorf("could not find postgres port in the env"))
	}
	postgresPort, err := strconv.Atoi(postgresPortStr)
	if err != nil {
		postgresErrs = multierr.Append(postgresErrs, fmt.Errorf("could not parse postgres port: %s", err))
	}

	if postgresErrs != nil {
		return nil, postgresErrs
	}

	return &PostgresConfiguration{
		User:           user,
		Password:       password,
		DBName:         dbName,
		Host:           host,
		Port:           postgresPort,
		SSLMode:        sslMode,
		MigrationsPath: migrationsPath,
	}, nil
}

func (r reader) auth() (*AuthorizationConfig, error) {
	var authErrs error
	authModelPath, ok := os.LookupEnv("AUTH_MODEL_PATH")
	if !ok {
		authErrs = fmt.Errorf("could not find auth model path in env")
	}
	authPolicyPath, ok := os.LookupEnv("AUTH_POLICY_PATH")
	if !ok {
		authErrs = multierr.Append(authErrs, fmt.Errorf("could not find auth policy path in env"))
	}

	if authErrs != nil {
		return nil, authErrs
	}

	authPaths := AuthPaths{
		AuthModelPath:  authModelPath,
		AuthPolicyPath: authPolicyPath,
	}

	return &AuthorizationConfig{
		AuthPaths: authPaths,
	}, nil
}

func (r reader) cors() (*CORSPolicy, error) {
	rawHosts, ok := os.LookupEnv("SERVER_ALLOWED_HOSTS")
	if !ok {
		return nil, fmt.Errorf("could not find server allowed hosts in the env")
	}

	allowedHosts := strings.Split(rawHosts, ",")

	if len(allowedHosts) == 0 {
		allowedHosts = append(allowedHosts, rawHosts)
	}

	return &CORSPolicy{
		ServerAllowedHosts: allowedHosts,
	}, nil
}
